package route

import (
	"net/http"
)

type TokenCookie interface {
	Token(r *http.Request) string
	Cookie(token string) http.Cookie
}
type tokenCookie struct {
	key    string
	domain string
}

func NewTokenCookie(key, domain string) TokenCookie {
	return &tokenCookie{
		key:    key,
		domain: domain,
	}
}

func (t *tokenCookie) Token(r *http.Request) string {
	for _, c := range r.Cookies() {
		if c.Name == t.key {
			return c.Value
		}
	}
	return ""
}

func (t *tokenCookie) Cookie(token string) http.Cookie {
	return http.Cookie{
		Domain: t.domain,
		Name:   t.key,
		Value:  token,
	}
}
