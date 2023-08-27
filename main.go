package main

import (
	"fmt"
	"github.com/wangyuheng/doorman/core/utils"
	"github.com/wangyuheng/doorman/pkg/middleware"
	"github.com/wangyuheng/doorman/proxy"
	"github.com/wangyuheng/doorman/sso"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wangyuheng/doorman/config"
)

func main() {
	config.Init()

	f := sso.Register(*config.GetConfig())

	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	r := gin.New()
	r.Use(middleware.SSO(config.GetConfig().Cookie.Name, config.GetConfig().Domain, config.GetConfig().CallbackPath, config.GetConfig().Feishu.RedirectUri))

	r.Any("/*Any", func(ctx *gin.Context) {
		if utils.EqualsPath(ctx.Request.URL.Path, config.GetConfig().CallbackPath) {
			sso.Handler(f)(ctx)
			return
		}

		proxy.NewGinHandler(config.GetConfig().RealAddr)(ctx)
	})

	err := r.Run(fmt.Sprintf(":%d", config.GetConfig().PORT))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
