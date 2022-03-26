package sso

import (
	"log"

	"github.com/geeklubcn/doorman/sso/dingtalk"

	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/core/auth"
	"github.com/geeklubcn/doorman/core/route"
	"github.com/geeklubcn/doorman/core/route/state"
	"github.com/geeklubcn/doorman/core/token"
	"github.com/geeklubcn/doorman/sso/feishu"
)

const (
	FEISHU   = iota
	DINGTALK // https://open.dingtalk.com/document/isvapp-server/obtain-identity-credentials
)

const (
	CODE = iota
	STATE
)

type PARAM int

func Register(mode int8, config conf.Config) *Fact {
	switch mode {
	case FEISHU:
		return &Fact{
			map[PARAM]string{
				CODE:  "code",
				STATE: "state",
			},
			feishu.NewIdentifier(feishu.NewApi(config.Feishu)),
			token.Jwt,
			token.Jwt,
			route.NewTokenCookie(config.Cookie.Name, config.Cookie.Domain),
			feishu.NewRouter(state.SimpleState{}, config.Feishu),
		}
	case DINGTALK:
		return &Fact{
			map[PARAM]string{
				CODE:  "authCode",
				STATE: "state",
			},
			dingtalk.NewIdentifier(dingtalk.NewApi(config.Dingtalk)),
			token.Jwt,
			token.Jwt,
			route.NewTokenCookie(config.Cookie.Name, config.Cookie.Domain),
			dingtalk.NewRouter(state.SimpleState{}, config.Dingtalk),
		}
	}
	log.Fatalf("Unsupport mode: %d", mode)
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
