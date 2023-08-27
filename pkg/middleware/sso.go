package middleware

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wangyuheng/doorman/core/utils"
	"net/http"
)

func SSO(cookieName, domain, path, redirectUri string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.EqualsPath(c.Request.URL.Path, path) {
			c.Next()
			return
		}
		queryToken := c.Query(cookieName)
		if queryToken == "" {
			cookieToken, err := c.Cookie(cookieName)
			if err != nil {
				logrus.WithError(err).Warn("get cookie fail!")
				redirectUrl := base64.URLEncoding.EncodeToString([]byte(redirectUri))
				c.Redirect(http.StatusFound, fmt.Sprintf("%s/%s?redirectUrl=%s", domain, path, redirectUrl))
				c.Abort()
				return
			}
			c.Set(cookieName, cookieToken)
		} else {
			c.Set(cookieName, queryToken)
		}
		c.Next()
	}
}
