package config

import (
	"github.com/spf13/viper"
	"strings"
)

var cfg *Config

type SSOMode string

const (
	ModeFeishu   SSOMode = "FEISHU"
	ModeDingtalk SSOMode = "DINGTALK" // https://open.dingtalk.com/document/isvapp-server/obtain-identity-credentials
)

const (
	Port               = "PORT"
	Domain             = "DOMAIN"
	RealAddr           = "REAL_ADDR"
	Mode               = "MODE"
	CallbackPath       = "CALLBACK_PATH"
	CookieName         = "COOKIE_NAME"
	CookieDomain       = "COOKIE_DOMAIN"
	FeishuBaseURL      = "FEISHU_BASEURL"
	FeishuClientID     = "FEISHU_CLIENTID"
	FeishuClientSecret = "FEISHU_CLIENTSECRET"
	FeishuRedirectUri  = "FEISHU_RedirectUri"
)

type Config struct {
	PORT         int64     `mapstructure:"port"`
	Domain       string    `mapstructure:"domain"`
	CallbackPath string    `mapstructure:"callback_path"`
	RealAddr     []string  `mapstructure:"real_addr"`
	Mode         SSOMode   `mapstructure:"mode"`
	Cookie       Cookie    `mapstructure:"cookie"`
	Feishu       *Feishu   `mapstructure:"feishu"`
	Dingtalk     *Dingtalk `mapstructure:"dingtalk"`
}

type Cookie struct {
	Name   string `mapstructure:"name"`
	Domain string `mapstructure:"Domain"`
}

type Feishu struct {
	BaseUrl      string `mapstructure:"base_url"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUri  string `mapstructure:"redirect_uri"`
}

type Dingtalk struct {
	ApiUrl       string `mapstructure:"api_url"`
	LoginUrl     string `mapstructure:"login_url"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUri  string `mapstructure:"redirect_uri"`
}

func GetConfig() *Config {
	return cfg
}

func Init() {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault(Port, 8080)

	for _, k := range []string{
		Port, Domain, RealAddr, Mode, CallbackPath,
		CookieName, CookieDomain,
		FeishuBaseURL, FeishuClientID, FeishuClientSecret, FeishuRedirectUri,
	} {
		if err := v.BindEnv(k); err != nil {
			panic(err)
		}

	}

	cfg = &Config{
		PORT:         v.GetInt64(Port),
		Domain:       v.GetString(Domain),
		CallbackPath: v.GetString(CallbackPath),
		RealAddr:     v.GetStringSlice(RealAddr),
		Mode:         SSOMode(strings.ToUpper(v.GetString(Mode))),
		Cookie: Cookie{
			Name:   v.GetString(CookieName),
			Domain: v.GetString(CookieDomain),
		},
	}

	switch cfg.Mode {
	case ModeFeishu:
		cfg.Feishu = &Feishu{
			BaseUrl:      v.GetString(FeishuBaseURL),
			ClientId:     v.GetString(FeishuClientID),
			ClientSecret: v.GetString(FeishuClientSecret),
			RedirectUri:  v.GetString(FeishuRedirectUri),
		}
	}

}
