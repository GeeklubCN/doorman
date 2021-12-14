package token

import (
	"time"

	"github.com/geeklubcn/doorman/core"
	"github.com/golang-jwt/jwt/v4"
)

var JwtFactory = &jwtFactory{
	signingMethod: jwt.SigningMethodHS256,
	secret:        []byte("doorman-secret"),
	issuer:        "doorman",
	expire:        24 * time.Hour,
	notBefore:     time.Now(),
}

type jwtFactory struct {
	signingMethod jwt.SigningMethod
	secret        []byte
	issuer        string
	expire        time.Duration
	notBefore     time.Time
}

func (j *jwtFactory) verify(token string) (core.Identification, bool) {
	c := claims{}
	if _, err := jwt.ParseWithClaims(token, &c, JwtFactory.verifyKey); err != nil {
		return "", false
	}
	return c.Id, true
}

type claims struct {
	jwt.RegisteredClaims
	Id core.Identification `json:"id"`
}

func (j *jwtFactory) create(id core.Identification) (string, error) {
	token := jwt.NewWithClaims(j.signingMethod, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.expiresAt()),
			Issuer:    j.issuer,
		},
		Id: id,
	})
	return token.SignedString(j.secret)
}

func (j *jwtFactory) expiresAt() time.Time {
	return time.Now().Add(j.expire)
}

func (j *jwtFactory) verifyKey(t *jwt.Token) (interface{}, error) {
	return j.secret, nil
}
