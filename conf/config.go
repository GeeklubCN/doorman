package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Config *config
	Feishu *feishu
)

const (
	Domain             = "DOMAIN"
	CookieDomain       = "COOKIE_DOMAIN"
	CookieKey          = "COOKIE_KEY"
	FeishuBaseurl      = "FEISHU_BASEURL"
	FeishuClientid     = "FEISHU_CLIENTID"
	FeishuClientsecret = "FEISHU_CLIENTSECRET"
	FeishuRedirecturi  = "FEISHU_REDIRECTURI"
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
	// load env
	v.AutomaticEnv()
	if v.IsSet(Domain) {
		Config.Domain = v.GetString(Domain)
	}
	if v.IsSet(CookieDomain) {
		Config.Cookie.Domain = v.GetString(CookieDomain)
	}
	if v.IsSet(CookieKey) {
		Config.Cookie.Key = v.GetString(CookieKey)
	}
	if v.IsSet(FeishuBaseurl) {
		Config.Feishu.BaseUrl = v.GetString(FeishuBaseurl)
	}
	if v.IsSet(FeishuClientid) {
		Config.Feishu.ClientId = v.GetString(FeishuClientid)
	}
	if v.IsSet(FeishuClientsecret) {
		Config.Feishu.ClientSecret = v.GetString(FeishuClientsecret)
	}
	if v.IsSet(FeishuRedirecturi) {
		Config.Feishu.RedirectUri = v.GetString(FeishuRedirecturi)
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
