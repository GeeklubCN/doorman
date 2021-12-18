package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Config *config
	Feishu *feishu
)

func Init() {
	Config = &config{}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf")
	err := v.ReadInConfig()
	if err != nil {
		logrus.Panicf("Fatal error read config file: %s \n", err)
	}
	if err = v.Unmarshal(Config); err != nil {
		logrus.Panicf("Fatal error unmarshal config file: %s \n", err)
	}

	Feishu = Config.Feishu
	logrus.Debugf("read config yaml and parsed. config:%s. feishu:%s", Config, Feishu)
}

type config struct {
	Domain string
	Cookie struct {
		Key    string
		Domain string
	}
	Feishu *feishu
}

type feishu struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	RedirectUri  string
}
