package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SSO(key, location string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/doorman" {
			c.Next()
			return
		}
		queryToken := c.Query("doorman_token")
		if queryToken == "" {
			cookieToken, err := c.Cookie(key)
			if err != nil {
				logrus.Warn("get cookie fail!", err)
				url := c.Request.RequestURI
				redirectUrl := base64.URLEncoding.EncodeToString([]byte(url))
				c.Redirect(http.StatusFound, fmt.Sprintf("%s?redirectUrl=%s", location, redirectUrl))
				c.Abort()
				return
			}
			c.Set(key, cookieToken)
		} else {
			c.Set(key, queryToken)
		}
		c.Next()
	}
}
