package main

import (
	"flag"
	"log"

	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/sso"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	configPath := flag.String("c", "./conf/config.yaml", "Path to the configuration filename")
	flag.Parse()

	config, err := conf.NewParser().Parse(*configPath)
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
	r.Any("/", sso.Handler(f))

	err = r.Run(":80")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
