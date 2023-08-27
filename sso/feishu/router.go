package feishu

import (
	"fmt"
	"net/url"

	"github.com/wangyuheng/doorman/config"
	"github.com/wangyuheng/doorman/core/route"
	"github.com/wangyuheng/doorman/core/route/state"
)

type router struct {
	state  state.State
	config config.Feishu
}

func NewRouter(s state.State, config config.Feishu) route.Router {
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
		r.config.BaseUrl, r.config.ClientId, redirectUrl, r.state.Encode(s))
}
