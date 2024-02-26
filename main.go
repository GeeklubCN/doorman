package main

import (
	"flag"
	"log"

	"github.com/geeklubcn/doorman/middleware"
	"github.com/geeklubcn/doorman/proxy"
	"github.com/geeklubcn/doorman/sso"

	"github.com/geeklubcn/doorman/conf"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Cmd struct {
	configPath string
}

func parseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.configPath, "c", "./conf/config.yaml", "Path to the configuration filename")
	flag.Parse()
	return cmd
}

func main() {
	cmd := parseCmd()

	config, err := conf.NewParser().Parse(cmd.configPath)
	if err != nil {
		logrus.Fatal("read config fail: ", err)
	}
	if err = config.Init(); err != nil {
		logrus.Fatal("init config fail: ", err)
	}

	f := sso.Register(sso.FEISHU, config)

	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	r := gin.New()
	r.Use(middleware.SSO(config.Cookie.Name, config.Domain+"/doorman"))

	r.Any("/*Any", func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/doorman" {
			sso.Handler(f)(ctx)
			return
		}
		proxy.NewGinHandler(config.RealAddr)(ctx)
	})

	err = r.Run(":80")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
