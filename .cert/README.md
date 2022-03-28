# 本地 https 证书

下载 mkcert

```shell
brew install mkcert
```

安装

```shell
mkcert -install
```

生成证书

```shell
mkcert sso.geeklub.cn
```

请求本地https地址 

```shell
kubectl get pods --all-namespaces --kubeconfig config.auth --insecure-skip-tls-verify
```

**config.auth**内容

```yaml
apiVersion: v1
clusters:
- cluster:
    server: https://sso.geeklub.cn:8443
  name: default
contexts:
- context:
    cluster: default
    namespace: default
    user: default
  name: default
current-context: "default"
kind: Config
preferences: {}
users:
- name: default
  user:
    password: 123456
    username: admin
```