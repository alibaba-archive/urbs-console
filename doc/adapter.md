# 身份验证

Urbs 支持三方实现用户的身份验证，便于集成到现有用户体系中；主要分二块：

1. 获取用户配置项的身份验证；
2. 访问灰度管控后台的身份验证、权限控制。

其 1 和 2 的身份验证是同一套用户体系，管控后台多了操作权限控制。

## 配置项

Urbs 没有单独的用户登录模块，主要考虑的是接入方一般都有自己的用户登陆体系，和已有体系整合的方便性远大于独立的一套。所以 Urbs 定义了一套 API 规范，使用方在接入时实现接口，即可对访问的用户进行身份验证。

验证流程：

1. 访问 Urbs API，比如 urbs.teambition.com/v1/users/settings:unionAll；
2. Urbs API携带身份（cookie）通过 http 的方式请求接入方实现的接口；
3. 接口验证身份（cookie）成功则返回用户信息，失败返回错误表示验证失败。

### 使用步骤

配置实现的接口：

```yaml
thrid:
    key: "a29b3a29b3a29b30x1a29b38"
    user_auth: 
        url: "已实现的接口地址，比如 http://localhost:8081/v1/users:verify"
        cookie_key: "要读取的 cookie 名称，会携带给接口"

```

访问 API 时，Urbs 使用 POST 方式请求上面配置的接口，body 是 json 格式，字段格式如下：

```json
{
  "cookie": "cookie",
  "signed": "cookie 的签名，Urbs 自动读取 cookie_key.sig 名字的 cookie 值，携带到，没有则为空"
}
```
需要返回 json 格式的数据：
```json
{
    "result": "用户标识"
}
```

Urbs 通过返回的用户标识去查询其下的配置项。

## 管控后台

管控后台身份验证相比 API 身份证多了一步，拿到用户标识后，再去查用户是否有访问后台界面的权限，用户的权限由超级管理员添加。

超级管理员可以在配置添加：

```yaml
superAdmins: ["用户标识"]
```

用户的后台浏览权限，可在后台 `管理` 中添加。用户的具体某项操作权限可以通过添加负责人添加。比如：

超级管理员创建个产品，指定 A 为负责人，A 即可更新产品、创建其下的模块1。 

A 可以指定 B 为模块1 负责人，然后 B 可以更新模块、创建配置项， B 也可以指定其下配置项的负责人。

权限不可以继承 ，产品的负责人，不可以操作其下的模块，如需要的话，在创建模块时，指定自己也为模块负责人。
# 功能扩展
除身份验证外，Urbs 还支持其他自定义功能，以增强 Urbs 灰度系统。

## 配置项通知

对于单页面的 Web 应用、移动端、PC 端来说，发布灰度或回滚灰度后，并不能及时通知前端，需要主动定时拉取新配置项才可以。

另一种方式是主动推送变化，Urbs 提供 hook 机制，使用者可以订阅配置项的变化，接收到事件后，利用已有的消息系统，推送给前端。

Urbs 提供订阅的事件有：

- 发布配置项：`setting.publish`；
- 撤回配置项发布记录：`setting.recall`；
- 配置项中移除群组或成员： `setting.remove`。

推送的内容和通过 API 拉取的配置项保持一致，Urbs 使用 POST 方式请求 API。body 是 json 格式，字段格式如下：

```json
{
        "product": "产品名称", 
        "module": "功能模块名称",
        "name": "配置项名称",
        "value": "配置项的值，除发布配置项外，其他事件为空字符串",
        "assignedAt": "分配配置项的时间"
}
```

配置 hook 地址：

```yaml
thrid:
   hook:
      url: "hook 地址，如 http://localhost:8081/v1/urbs/setting/hooks"
      events: ["setting.publish","setting.recall","setting.remove"]  
```

## 自定义灰度

Urbs 是面向开源来设计开发的，不依赖特定的用户身份体系、特定的业务逻辑。在与自身产品内部系统集成时，可以开发适配应用，主要用来：

- 在获取用户配置项时、访问管控后台的身份验证；
- 存量用户数据导入 Urbs 系统，增量用户可以通过监听内部系统的 Webhook、MQ 添加到 Urbs 系统；
- 除了 Urbs 系统本身支持的用户、群组、百分比、新用户百分比的灰度方式，如需要扩展更多的灰度方式，可以通过调用 Urbs-setting API 配合自身的用户体系实现。

比如以用户自身属性的进行多维度的灰度，如下伪代码：

```go
user := watch.Users()
// 添加用户到 Urbs 系统
request.Post("/api/v1/users:batch", user)

if IsChild(User) { // 针对特定的人群进行灰度
   Assign(User)
}
if IsChina(User) {  // 针对地区进行灰度
   Assign(User)
}
if IsLevel3(User) {  // 针对用户会员级别进行灰度
   Assign(User)
}
if IsIos(User) {  // 针对 IOS 用户进行灰度
   Assign(User)
}
// 及其他维度

func Assign(user User){
   // 发布配置项到用户上
   request.Post("/api/v1/products/product/modules/module/settings/setting:assign", user)
   // 或发布标签到用户上
   request.Post("/api/v1/products/product/label/label:assign", user)
}
```
