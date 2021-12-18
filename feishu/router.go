package feishu

import (
	"fmt"
	"net/url"

	"github.com/geeklubcn/doorman/conf"
	"github.com/geeklubcn/doorman/core/route"
	"github.com/geeklubcn/doorman/core/route/state"
)

type router struct {
	state state.State
}

func NewRouter(s state.State) route.Router {
	return &router{
		state: s,
	}
}

func (r *router) LoginUrl() string {
	redirectUrl := url.PathEscape(conf.Feishu.RedirectUri)
	return fmt.Sprintf("%s/suite/passport/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&state=%s",
		conf.Feishu.BaseUrl, conf.Feishu.ClientId, redirectUrl, r.state.Encode(conf.Feishu.RedirectUri))
}
