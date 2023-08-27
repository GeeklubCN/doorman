package main

import (
	"github.com/wangyuheng/doorman/config"
	"os"
)

func init() {
	os.Setenv(config.Domain, "https://sso.yuheng.wang")
	os.Setenv(config.RealAddr, "http://t.tt")
	os.Setenv(config.Mode, "FEISHU")
	os.Setenv(config.CallbackPath, "doorman")
}
