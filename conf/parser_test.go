package conf

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	configPath := "/tmp/config.yaml"
	configContent := []byte(`
domain: "http://sso.geeklub.cn"
cookie:
  name: "doorman_token"
  domain: "sso.geeklub.cn"
feishu:
  base_url: "https://passport.feishu.cn"
  client_id: "cli_xxx"
  client_secret: "6nTXxxxx"
  redirect_uri: "http://sso.geeklub.cn"
`)
	if err := ioutil.WriteFile(configPath, configContent, 0644); err != nil {
		t.FailNow()
	}
	config, err := NewParser().Parse(configPath)
	assert.NoError(t, err)
	assert.Equal(t, "http://sso.geeklub.cn", config.Feishu.RedirectUri)
	assert.Equal(t, "http://sso.geeklub.cn", config.Domain)
	assert.Equal(t, "doorman_token", config.Cookie.Name)
}
