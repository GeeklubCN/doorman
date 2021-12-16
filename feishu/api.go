package feishu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var A Api

func Init(c Config) {
	A = &api{
		client: &http.Client{},
		config: c,
	}
}

type Config struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

type Api interface {
	getToken(code string) ([]byte, error)
	getUserInfo(authorization string) ([]byte, error)
}

type api struct {
	client *http.Client
	config Config
}

func (a *api) getToken(code string) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("%s/suite/passport/oauth/token", a.config.BaseUrl), url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {a.config.ClientId},
		"client_secret": {a.config.ClientSecret},
		"code":          {code},
		"redirect_uri":  {a.config.RedirectUri},
	})
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)

}

func (a *api) getUserInfo(authorization string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/suite/passport/oauth/userinfo", a.config.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", authorization)
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

type oauthToken struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

func (t *oauthToken) getAuthorization() string {
	return fmt.Sprintf("%s %s", t.TokenType, t.AccessToken)
}

type UserInfo struct {
	Sub          string `json:"sub"`
	Picture      string `json:"picture"`
	Name         string `json:"name"`
	EnName       string `json:"en_name"`
	TenantKey    string `json:"tenant_key"`
	AvatarUrl    string `json:"avatar_url"`
	AvatarThumb  string `json:"avatar_thumb"`
	AvatarMiddle string `json:"avatar_middle"`
	AvatarBig    string `json:"avatar_big"`
	OpenId       string `json:"open_id"`
	UnionId      string `json:"union_id"`
}
