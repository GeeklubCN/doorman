package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SSO(key, location string) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryToken := c.Query(key)
		if queryToken == "" {
			cookieToken, err := c.Cookie(key)
			if err != nil {
				logrus.Warn("get cookie fail!", err)
				url := c.Request.Proto + c.Request.Host + c.Request.RequestURI
				c.Redirect(http.StatusFound, fmt.Sprintf("%s?redirectUrl=%s", location, url))
				return
			}
			c.Set(key, cookieToken)
		}
		c.Set(key, queryToken)

		c.Next()
	}
}
