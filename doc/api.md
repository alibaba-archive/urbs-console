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
  "name": "urbs-setting",
  "version": "v0.1.0",
  "gitSHA1": "cd7e82a",
  "buildTime": "2020-03-25T06:24:25Z"
}
```

<h3 id="获取版本信息-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|请求成功|[Version](#schemaversion)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="urbs-console-user">User</h1>

User 用户相关接口

## 接口返回 user 的 settings，按照 setting 设置时间反序，支持分页，包含了 user 从属的 group 的 settings

> Code samples

```shell
# You can also use wget
curl -X GET http://urbs-console:8080/v1/users/{uid}/settings:unionAll \
  -H 'Accept: application/json' \
  -H 'Authorization: string'

```

```http
GET http://urbs-console:8080/v1/users/{uid}/settings:unionAll HTTP/1.1
Host: urbs-console:8080
Accept: application/json
Authorization: string

```

`GET /v1/users/{uid}/settings:unionAll`

<h3 id="接口返回-user-的-settings，按照-setting-设置时间反序，支持分页，包含了-user-从属的-group-的-settings-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string|false|请求 JWT token, 格式如: `Bearer xxx`|
|uid|path|string|true|用户 UID，当 uid 对应用户不存在时，该接口会返回空列表|
|product|query|string|false|产品名称，当 product 对应产品不存在时，该接口会返回空列表|
|client|query|string|false|产品名称，当 client 对应产品不存在时，该接口会返回空列表|
|pageSize|query|integer(int32)|false|分页大小，默认为 10，(1-1000]|
|pageToken|query|string|false|分页请求标记，来自于响应结果的 nextPageToken|

> Example responses

> 200 Response

```json
{
  "totalSize": 99,
  "nextPageToken": "hid.fog-AAAAAABWPxcs-4UJ3Aw9",
  "result": [
    {
      "hid": "AwAAAAAAAAB25V_QnbhCuRwF",
      "module": "teambition",
      "name": "beta",
      "value": "string",
      "last_value": "string",
      "created_at": "2020-03-25T06:24:25Z",
      "updated_at": "2020-03-25T06:24:25Z"
    }
  ]
}
```

<h3 id="接口返回-user-的-settings，按照-setting-设置时间反序，支持分页，包含了-user-从属的-group-的-settings-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|请求成功|[MySettingsRes](#schemamysettingsres)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|请求失败|[ErrorResponse](#schemaerrorresponse)|

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
"hid.fog-AAAAAABWPxcs-4UJ3Aw9"

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
99

```

totalSize

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|totalSize|integer(int64)|false|none|当前分页查询的总数据量|

<h2 id="tocS_ErrorResponse">ErrorResponse</h2>
<!-- backwards compatibility -->
<a id="schemaerrorresponse"></a>
<a id="schema_ErrorResponse"></a>
<a id="tocSerrorresponse"></a>
<a id="tocserrorresponse"></a>

```json
{
  "error": "NotFound",
  "message": "user 50c32afae8cf1439d35a87e6 not found"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|error|string|false|none|错误代号|
|message|string|false|none|错误详情|

<h2 id="tocS_BoolRes">BoolRes</h2>
<!-- backwards compatibility -->
<a id="schemaboolres"></a>
<a id="schema_BoolRes"></a>
<a id="tocSboolres"></a>
<a id="tocsboolres"></a>

```json
{
  "result": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|result|boolean|false|none|是否成功|

<h2 id="tocS_Version">Version</h2>
<!-- backwards compatibility -->
<a id="schemaversion"></a>
<a id="schema_Version"></a>
<a id="tocSversion"></a>
<a id="tocsversion"></a>

```json
{
  "name": "urbs-setting",
  "version": "v0.1.0",
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

<h2 id="tocS_MySettingsRes">MySettingsRes</h2>
<!-- backwards compatibility -->
<a id="schemamysettingsres"></a>
<a id="schema_MySettingsRes"></a>
<a id="tocSmysettingsres"></a>
<a id="tocsmysettingsres"></a>

```json
{
  "totalSize": 99,
  "nextPageToken": "hid.fog-AAAAAABWPxcs-4UJ3Aw9",
  "result": [
    {
      "hid": "AwAAAAAAAAB25V_QnbhCuRwF",
      "module": "teambition",
      "name": "beta",
      "value": "string",
      "last_value": "string",
      "created_at": "2020-03-25T06:24:25Z",
      "updated_at": "2020-03-25T06:24:25Z"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|totalSize|[TotalSize](#schematotalsize)|false|none|当前分页查询的总数据量|
|nextPageToken|[NextPageToken](#schemanextpagetoken)|false|none|用于分页查询时用于获取下一页数据的 token，当为空值时表示没有下一页了|
|result|[[MySetting](#schemamysetting)]|false|none|none|

<h2 id="tocS_MySetting">MySetting</h2>
<!-- backwards compatibility -->
<a id="schemamysetting"></a>
<a id="schema_MySetting"></a>
<a id="tocSmysetting"></a>
<a id="tocsmysetting"></a>

```json
{
  "hid": "AwAAAAAAAAB25V_QnbhCuRwF",
  "module": "teambition",
  "name": "beta",
  "value": "string",
  "last_value": "string",
  "created_at": "2020-03-25T06:24:25Z",
  "updated_at": "2020-03-25T06:24:25Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|hid|string|false|none|灰度标签的 HID|
|module|string|false|none|产品模块名称|
|name|string|false|none|模块配置名称|
|value|string|false|none|模块配置的值|
|last_value|string|false|none|上次模块配置的值|
|created_at|string(date-time)|false|none|创建时间|
|updated_at|string(date-time)|false|none|更新时间|

