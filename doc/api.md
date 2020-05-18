---
title: urbs-console v0.1.0
language_tabs:
  - shell: Shell
  - http: HTTP
language_clients:
  - shell: ""
  - http: ""
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.0 -->

<h1 id="urbs-console">urbs-console v0.1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Urbs 灰度平台管控后台应用

Base URLs:

* <a href="http://urbs-console:8080">http://urbs-console:8080</a>

<h1 id="urbs-console-version">Version</h1>

获取 urbs-console 服务版本信息

## 获取版本信息

> Code samples

```shell
# You can also use wget
curl -X GET http://urbs-console:8080/version \
  -H 'Accept: application/json'

```

```http
GET http://urbs-console:8080/version HTTP/1.1
Host: urbs-console:8080
Accept: application/json

```

`GET /version`

> Example responses

> 200 Response

```json
{
  "name": "urbs-console",
  "version": "v1.2.0",
  "gitSHA1": "cd7e82a",
  "buildTime": "2020-03-25T06:24:25Z"
}
```

<h3 id="获取版本信息-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|version 返回结果|[Version](#schemaversion)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="urbs-console-user">User</h1>

User 用户相关接口

## 获取 user 的所有 settings 信息

> Code samples

```shell
# You can also use wget
curl -X GET http://urbs-console:8080/v1/users/settings:unionAll?product=string \
  -H 'Accept: application/json' \
  -H 'Authorization: string'

```

```http
GET http://urbs-console:8080/v1/users/settings:unionAll?product=string HTTP/1.1
Host: urbs-console:8080
Accept: application/json
Authorization: string

```

`GET /v1/users/settings:unionAll`

<h3 id="获取-user-的所有-settings-信息-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string|true|请求 JWT token, 格式如: `Bearer xxx`|
|product|query|string|true|产品名称|
|client|query|string|false|客户端标识，例如 web、ios、android、windows、macos|
|channel|query|string|false|客户端渠道，例如 stable、beta、dev|
|pageSize|query|integer(int32)|false|分页大小，默认为 10，(1-1000]|
|pageToken|query|string|false|分页请求标记，来自于响应结果的 nextPageToken|

> Example responses

> 200 Response

```json
{
  "nextPageToken": "",
  "result": [
    {
      "hid": "AwAAAAAAAAB25V_QnbhCuRwF",
      "product": "teambition",
      "module": "task",
      "name": "task-share",
      "desc": "string",
      "value": "disable",
      "lastValue": "",
      "release": 1,
      "assignedAt": "2020-03-25T06:24:25Z"
    }
  ]
}
```

<h3 id="获取-user-的所有-settings-信息-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|用户或群组被指派的配置项列表返回结果|Inline|

<h3 id="获取-user-的所有-settings-信息-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» nextPageToken|[NextPageToken](#schemanextpagetoken)|false|none|用于分页查询时用于获取下一页数据的 token，当为空值时表示没有下一页了|
|» result|[[MySetting](#schemamysetting)]|false|none|none|
|»» hid|string|false|none|配置项的 hid|
|»» product|string|false|none|配置项所属的产品名称|
|»» module|string|false|none|配置项所属的功能模块名称|
|»» name|string|false|none|配置项名称|
|»» desc|string|false|none|配置项描述，|
|»» value|string|false|none|配置项值|
|»» lastValue|string|false|none|配置项值|
|»» release|integer(int64)|false|none|被设置批次|
|»» assignedAt|string(date-time)|false|none|被设置时间|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_NextPageToken">NextPageToken</h2>
<!-- backwards compatibility -->
<a id="schemanextpagetoken"></a>
<a id="schema_NextPageToken"></a>
<a id="tocSnextpagetoken"></a>
<a id="tocsnextpagetoken"></a>

```json
""

```

nextPageToken

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|nextPageToken|string|false|none|用于分页查询时用于获取下一页数据的 token，当为空值时表示没有下一页了|

<h2 id="tocS_TotalSize">TotalSize</h2>
<!-- backwards compatibility -->
<a id="schematotalsize"></a>
<a id="schema_TotalSize"></a>
<a id="tocStotalsize"></a>
<a id="tocstotalsize"></a>

```json
1

```

totalSize

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|totalSize|integer(int64)|false|none|当前分页查询的总数据量|

<h2 id="tocS_Version">Version</h2>
<!-- backwards compatibility -->
<a id="schemaversion"></a>
<a id="schema_Version"></a>
<a id="tocSversion"></a>
<a id="tocsversion"></a>

```json
{
  "name": "urbs-console",
  "version": "v1.2.0",
  "gitSHA1": "cd7e82a",
  "buildTime": "2020-03-25T06:24:25Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|false|none|服务名称|
|version|string|false|none|当前版本|
|gitSHA1|string|false|none|git commit hash|
|buildTime|string(date-time)|false|none|打包构建时间|

<h2 id="tocS_MySetting">MySetting</h2>
<!-- backwards compatibility -->
<a id="schemamysetting"></a>
<a id="schema_MySetting"></a>
<a id="tocSmysetting"></a>
<a id="tocsmysetting"></a>

```json
{
  "hid": "AwAAAAAAAAB25V_QnbhCuRwF",
  "product": "teambition",
  "module": "task",
  "name": "task-share",
  "desc": "string",
  "value": "disable",
  "lastValue": "",
  "release": 1,
  "assignedAt": "2020-03-25T06:24:25Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|hid|string|false|none|配置项的 hid|
|product|string|false|none|配置项所属的产品名称|
|module|string|false|none|配置项所属的功能模块名称|
|name|string|false|none|配置项名称|
|desc|string|false|none|配置项描述，|
|value|string|false|none|配置项值|
|lastValue|string|false|none|配置项值|
|release|integer(int64)|false|none|被设置批次|
|assignedAt|string(date-time)|false|none|被设置时间|

