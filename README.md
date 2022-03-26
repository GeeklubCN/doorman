# doorman

- SSO

## 使用方式

### Docker

修改 `docker-custom-conf.yaml` 配置

```shell
docker run -v $PWD/docker-custom-conf.yaml:/go/src/github.com/geeklubcn/doorman/conf/config.yaml:ro wangyuheng/doorman:v1
```

### Docker Compose

修改 `docker-compose.yaml` 的环境变量

```shell
docker-compose up
```

## 开发

### Docker

```shell
docker build -t wangyuheng/doorman:v1 .
```

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

#### 1. 按照[官方文档](https://open.dingtalk.com/document/org/create-orgapp)创建H5网页应用，获取 clientId及clientSecret

![doorman-dingtalk-app](https://raw.githubusercontent.com/geeklubcn/doorman/master/.doc/doorman-dingtalk-app.png)

#### 2. 管理后台->应用功能->登录与分享->回调域名

![doorman-dingtalk-security](https://raw.githubusercontent.com/geeklubcn/doorman/master/.doc/doorman-dingtalk-security.png)

#### 3. 管理后台->基础信息->权限管理

添加 **通讯录个人信息读权限** 权限

![doorman-dingtalk-permission](https://raw.githubusercontent.com/geeklubcn/doorman/master/.doc/doorman-dingtalk-permission.png)

#### 4. 配置yaml配置文件

```yaml
dingtalk:
  api_url: "https://api.dingtalk.com"
  login_url: "https://login.dingtalk.com"
  client_id: "dingxxx"
  client_secret: "rxxx"
  redirect_uri: "http://sso.geeklub.cn/doorman"
```

#### 参考

- https://open.dingtalk.com/document/isvapp-server/obtain-identity-credentials

### 企业微信｜weixin

**【TBD】**

## SDK

### Golang

下载依赖包

```shell
go get github.com/geeklubcn/doorman/middleware
```

使用middleware

```shell
r.Use(middleware.SSO("doorman_token", "http://sso.geeklub.cn"))
```

## 参考

- Jwt: https://jwt.io