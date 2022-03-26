package dingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/geeklubcn/doorman/conf"
)

type Api interface {
	getToken(code string) ([]byte, error)
	getUserInfo(authorization string) ([]byte, error)
}

func NewApi(config conf.Dingtalk) Api {
	return &api{
		client: &http.Client{},
		config: config,
	}
}

type api struct {
	client *http.Client
	config conf.Dingtalk
}

/**
https://open.dingtalk.com/document/isvapp-server/obtain-user-token
*/
func (a *api) getToken(code string) ([]byte, error) {
	body, _ := json.Marshal(map[string]string{
		"clientId":     a.config.ClientId,
		"clientSecret": a.config.ClientSecret,
		"code":         code,
		"grantType":    "authorization_code",
	})
	resp, err := http.Post(fmt.Sprintf("%s/v1.0/oauth2/userAccessToken", a.config.ApiUrl),
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)

}

/**
https://open.dingtalk.com/document/isvapp-server/dingtalk-retrieve-user-information
*/
func (a *api) getUserInfo(authorization string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1.0/contact/users/me", a.config.ApiUrl), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-acs-dingtalk-access-token", authorization)
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r userInfoResp
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type oauthToken struct {
	ExpireIn     int    `json:"expireIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	CorpId       string `json:"corpId"`
}

func (t *oauthToken) getAuthorization() string {
	return t.AccessToken
}

type userInfoResp struct {
	Nick      string `json:"nick"`
	UnionId   string `json:"unionId"`
	AvatarUrl string `json:"avatarUrl"`
	OpenId    string `json:"openId"`
	Mobile    string `json:"mobile"`
	StateCode string `json:"stateCode"`
}
