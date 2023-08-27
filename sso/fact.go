package sso

import (
	"github.com/wangyuheng/doorman/sso/dingtalk"

	"github.com/wangyuheng/doorman/config"
	"github.com/wangyuheng/doorman/core/auth"
	"github.com/wangyuheng/doorman/core/route"
	"github.com/wangyuheng/doorman/core/route/state"
	"github.com/wangyuheng/doorman/core/token"
	"github.com/wangyuheng/doorman/sso/feishu"
)

const (
	CODE = iota
	STATE
)

type PARAM int

func Register(cfg config.Config) *Fact {
	switch cfg.Mode {
	case config.ModeFeishu:
		return &Fact{
			map[PARAM]string{
				CODE:  "code",
				STATE: "state",
			},
			feishu.NewIdentifier(feishu.NewApi(*cfg.Feishu)),
			token.Jwt,
			token.Jwt,
			route.NewTokenCookie(cfg.Cookie.Name, cfg.Cookie.Domain),
			feishu.NewRouter(state.SimpleState{}, *cfg.Feishu),
		}
	case config.ModeDingtalk:
		return &Fact{
			map[PARAM]string{
				CODE:  "authCode",
				STATE: "state",
			},
			dingtalk.NewIdentifier(dingtalk.NewApi(*cfg.Dingtalk)),
			token.Jwt,
			token.Jwt,
			route.NewTokenCookie(cfg.Cookie.Name, cfg.Cookie.Domain),
			dingtalk.NewRouter(state.SimpleState{}, *cfg.Dingtalk),
		}
	}
	return nil
}

type Fact struct {
	Param       map[PARAM]string
	Identifier  auth.Identifier
	Factory     token.Factory
	Verifier    token.Verifier
	TokenCookie route.TokenCookie
	Router      route.Router
}

func (f *Fact) GetCode() string {
	return f.Param[CODE]
}

func (f *Fact) GetState() string {
	return f.Param[STATE]
}
