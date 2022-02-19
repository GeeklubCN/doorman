package feishu

import (
	"encoding/json"

	"github.com/geeklubcn/doorman/core/auth"
	"github.com/sirupsen/logrus"

	"github.com/geeklubcn/doorman/core"
)

type identifier struct {
	api Api
}

func NewIdentifier(api Api) auth.Identifier {
	return &identifier{api}
}

func (f *identifier) Identify(code string) (core.Identification, bool) {
	body, err := f.api.getToken(code)
	if err != nil {
		logrus.Errorf("get token fail! code:%s.", code, err)
		return "", false
	}
	var t oauthToken
	err = json.Unmarshal(body, &t)
	if err != nil {
		logrus.Errorf("token illegal! body:%s.", body, err)
		return "", false
	}
	body, err = f.api.getUserInfo(t.getAuthorization())
	if err != nil {
		logrus.Errorf("get user info fail! auth:%s.", t.getAuthorization(), err)
		return "", false
	}
	return core.Identification(body), true
}
