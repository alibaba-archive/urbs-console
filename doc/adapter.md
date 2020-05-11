## 集成
Urbs 支持扩展集成到，使用者现有的用户体系中，便于与现有系统整合。

主要分两大块，一是在访问灰度后台的用户身份验证、权限控制。二是灰度用户如何同步到 Urbs 中及后续的新增、删除用户的增量数据同步。

## 用户体系

Urbs 没有单独的用户登录模块，主要考虑的是接入方一般都有自己的用户登陆体系，和已有体系整合的方便性远大于独立的一套。所以 Urbs 定义了一套 API 规范，使用方在接入时实现接口，即可对访问的用户进行身份验证。

验证流程：

1. 访问 Urbs 后台域名，比如 urbs.teambition.com；
2. Urbs 后台携带身份（cookie）通过 http 的方式请求接入方实现的接口；
3. 接口验证身份（cookie）成功则返回用户信息，失败返回错误表示验证失败；

### 步骤

一、配置接口：

```
thrid:
    key: "a29b3a29b3a29b30x1a29b38"
    user_auth: 
        url: "已实现的接口地址，比如 http://local:8080//v1/users:verify"
        cookie_key: "要读取的 cookie 名称，会携带给接口"

```

二、Urbs 使用 POST 方式请求 API
携带的 `Body` 参数：

```json
{
  "cookie": "cookie",
  "signed": "cookie 的签名，Urbs 自动读取 cookie_key.sig 名字的 cookie 值，携带到，没有则为空"
}
```
需要返回 `json` 格式的数据：
```json
{
    "result": "用户标示"
}
```

