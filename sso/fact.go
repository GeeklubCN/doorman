package sso

import (
	"log"

	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/core/auth"
	"github.com/geeklubcn/doorman/core/route"
	"github.com/geeklubcn/doorman/core/route/state"
	"github.com/geeklubcn/doorman/core/token"
	"github.com/geeklubcn/doorman/feishu"
)

const (
	FEISHU = iota
)

func Register(mode int8) *Fact {
	switch mode {
	case FEISHU:
		return &Fact{
			feishu.Identifier,
			token.Jwt,
			token.Jwt,
			route.NewTokenCookie(conf.Config.Cookie.Key, conf.Config.Cookie.Domain),
			feishu.NewRouter(state.SimpleState{}),
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
