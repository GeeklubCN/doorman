package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/wangyuheng/doorman/core"
)

var Jwt = &jwtToken{
	signingMethod: jwt.SigningMethodHS256,
	secret:        []byte("doorman-secret"),
	issuer:        "doorman",
	expire:        24 * time.Hour,
	notBefore:     time.Now(),
}

type jwtToken struct {
	signingMethod jwt.SigningMethod
	secret        []byte
	issuer        string
	expire        time.Duration
	notBefore     time.Time
}

func (j *jwtToken) Create(id core.Identification) (string, error) {
	token := jwt.NewWithClaims(j.signingMethod, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.expiresAt()),
			Issuer:    j.issuer,
		},
		Id: id,
	})
	return token.SignedString(j.secret)
}

func (j *jwtToken) Verify(token string) (core.Identification, bool) {
	c := claims{}
	if _, err := jwt.ParseWithClaims(token, &c, j.verifyKey); err != nil {
		return "", false
	}
	return c.Id, true
}

type claims struct {
	jwt.RegisteredClaims
	Id core.Identification `json:"id"`
}

func (j *jwtToken) expiresAt() time.Time {
	return time.Now().Add(j.expire)
}

func (j *jwtToken) verifyKey(t *jwt.Token) (interface{}, error) {
	return j.secret, nil
}
