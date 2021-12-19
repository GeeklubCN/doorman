package main

import (
	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/sso"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	conf.Init()
	f := sso.Register(sso.FEISHU)

	r := gin.New()
	r.Any("/", sso.Handler(f))

	err := r.Run(":80")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
