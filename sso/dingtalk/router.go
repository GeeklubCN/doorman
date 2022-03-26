package dingtalk

import (
	"fmt"
	"net/url"

	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/core/route"
	"github.com/geeklubcn/doorman/core/route/state"
)

type router struct {
	state  state.State
	config conf.Dingtalk
}

func NewRouter(s state.State, config conf.Dingtalk) route.Router {
	return &router{
		state:  s,
		config: config,
	}
}

func (r *router) SourceUrl(s string) string {
	uri, _ := r.state.Decode(s)
	return uri
}

func (r *router) LoginUrl(s string) string {
	redirectUrl := url.PathEscape(r.config.RedirectUri)
	return fmt.Sprintf("%s/oauth2/auth?redirect_uri=%s&response_type=code&client_id=%s&scope=openid&state=%s&prompt=consent",
		r.config.LoginUrl,
		redirectUrl,
		r.config.ClientId,
		r.state.Encode(fmt.Sprintf("%s?%s", r.config.RedirectUri, s)),
	)
}
