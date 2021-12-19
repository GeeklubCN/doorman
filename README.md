# doorman

- SSO

## 单点登录｜SSO

### 飞书｜Feishu

#### 1. 按照[官方文档](https://open.feishu.cn/document/home/introduction-to-custom-app-development/self-built-application-development-process)创建应用，获取 clientId及clientSecret

![doorman-feishu-app](https://raw.githubusercontent.com/geeklubcn/doorman/master/.doc/doorman-feishu-app.png)

#### 2. 管理后台->安全设计->添加sso域名

![doorman-feishu-security](https://raw.githubusercontent.com/geeklubcn/doorman/master/.doc/doorman-feishu-security.png)

#### 3. 配置yaml配置文件

```yaml
feishu:
  baseUrl: "https://passport.feishu.cn"
  clientId: "cli_xxx"
  clientSecret: "6nTXxxx"
  redirectUri: "http://sso.geeklub.cn"
```

#### 参考
- https://open.feishu.cn/document/common-capabilities/sso/web-application-sso/web-app-overview

### 钉钉｜Dingtalk

**【TBD】**

### 企业微信｜weixin

**【TBD】**

### 参考

- Jwt: https://jwt.io