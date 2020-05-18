[基本概念](https://github.com/teambition/urbs-setting/blob/master/doc/concepts.md)

## 配置项灰度

### 后台添加配置项

访问后台界面，添加模块（比如 task）、添加配置项（robot-entrance）。

Module 、Setting 命名必须是英文字母，多个单词以连字符（-）拼接。

Setting 名至少由 2 个单词组成，一个配置项支持多个值，通过`,` 分割。

发布配置项到用户或群组上时，选择要分配的值。值的存储格式是字符串，支持任意字符，比如开关语义的 "true" 或 "false"，复杂语义的 json `{"key":value}`

### 客户端接入

访问 `GET /v1/users/settings:unionAll` API 获取当前用户拥有的配置项，返回是数组，具体见 [API 文档](https://github.com/teambition/urbs-console/blob/develop/doc/api.md)。

API 优先读取 cookie 中身份信息作为查询条件，适用于 web 端。

如果 Cookie 中身份信息不存在，读取 header 中 Authorization 的身份值作为查询条件，适用于移动端。

如果用户下配置较多，API 支持分页读取，用返回的 pageToken 来获取下一页。

## 环境灰度

如果功能改动较大，配置项不足以满足灰度要求，可以通过 label 方式实现多套环境同时存在，由用户身上 label 来决定访问的环境。

一、在管理后台添加 label。

二、在发布系统部署对应 label 的服务环境。

三、给用户打上 label ，网关根据用户身上的 label 来路由到不同的环境。

## 灰度用户

添加用户灰度，支持 4 种方式：

1. 批量指定具体用户
2. 批量指定群组
3. 按百分比灰度用户
4. 按百分比灰度新用户

