package feishu

import (
	"fmt"
	"net/url"

	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/core/route"
	"github.com/geeklubcn/doorman/core/route/state"
)

type router struct {
	state  state.State
	config conf.Feishu
}

func NewRouter(s state.State, config conf.Feishu) route.Router {
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
	return fmt.Sprintf("%s/suite/passport/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&state=%s",
		r.config.BaseUrl, r.config.ClientId, redirectUrl, r.state.Encode(fmt.Sprintf("%s?%s", r.config.RedirectUri, s)))
}
