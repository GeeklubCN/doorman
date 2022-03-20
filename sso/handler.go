package sso

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(f *Fact) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 	just forward if authenticated
		rUrl := c.Query("redirectUrl")
		t := f.TokenCookie.Token(c.Request)
		id, pass := f.Verifier.Verify(t)
		if pass {
			token, err := f.Factory.Create(id)
			if err == nil {
				cookie := f.TokenCookie.Cookie(token)
				c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, false, true)
			}
			c.Redirect(http.StatusFound, rUrl+"?doorman_token="+token)
			c.Abort()
			return
		}
		// callback
		code := c.Query("code")
		state := c.Query("state")
		if code != "" {
			id, pass := f.Identifier.Identify(code)
			if pass {
				token, err := f.Factory.Create(id)
				if err == nil {
					cookie := f.TokenCookie.Cookie(token)
					c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
				}
				c.Redirect(http.StatusFound, f.Router.SourceUrl(state))
				c.Abort()
				return
			}
		}
		// redirectToLogin
		c.Redirect(http.StatusFound, f.Router.LoginUrl(c.Request.URL.RawQuery))
		c.Abort()
		return
	}
}
