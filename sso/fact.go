package sso

import (
	"log"

	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/core/auth"
	"github.com/geeklubcn/doorman/core/route"
	"github.com/geeklubcn/doorman/core/route/state"
	"github.com/geeklubcn/doorman/core/token"
	"github.com/geeklubcn/doorman/sso/feishu"
)

const (
	FEISHU = iota
)

func Register(mode int8, config conf.Config) *Fact {
	switch mode {
	case FEISHU:
		return &Fact{
			feishu.NewIdentifier(feishu.NewApi(config.Feishu)),
			token.Jwt,
			token.Jwt,
			route.NewTokenCookie(config.Cookie.Name, config.Cookie.Domain),
			feishu.NewRouter(state.SimpleState{}, config.Feishu),
		}
	}
	log.Fatalf("Unsupport mode: %d", mode)
	return nil
}

type Fact struct {
	Identifier  auth.Identifier
	Factory     token.Factory
	Verifier    token.Verifier
	TokenCookie route.TokenCookie
	Router      route.Router
}
