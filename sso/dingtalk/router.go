package dingtalk

import (
	"fmt"
	"net/url"

	"github.com/wangyuheng/doorman/config"
	"github.com/wangyuheng/doorman/core/route"
	"github.com/wangyuheng/doorman/core/route/state"
)

type router struct {
	state  state.State
	config config.Dingtalk
}

func NewRouter(s state.State, config config.Dingtalk) route.Router {
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
