package feishu

import (
	"encoding/json"

	"github.com/geeklubcn/doorman/core"
)

var Identifier = &identifier{}

type identifier struct{}

func (f *identifier) Identify(code string) core.Identification {
	body, _ := A.getToken(code)
	var t oauthToken
	_ = json.Unmarshal(body, &t)
	body, _ = A.getUserInfo(t.getAuthorization())
	return core.Identification(body)
}
