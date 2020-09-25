# urbs-console
> Urbs 灰度平台管控后台应用

[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/teambition/urbs-console/master/LICENSE)

## Features
+ 灰度管理界面
+ 支持按用户、群组、比例灰度
+ 支持 ACL，管理后台用户权限
+ 支持操作日志
+ 支持三方实现用户身份验证，便于集成已有系统

## Depends
- urbs-setting >= 1.3.0

## Get Started

编译前端文件:

```shell
make budilweb
```

启动应用：

```shell
make dev
```

浏览器访问：

>  http://localhost:8081


## Documentation

[API 文档](https://github.com/teambition/urbs-console/blob/develop/doc/api.md)

[管控后台帮助文档](https://github.com/teambition/urbs-console/blob/develop/doc/help.md)

[前端接入文档](https://github.com/teambition/urbs-console/blob/develop/doc/client_usage.md)

[集成已有系统和扩展灰度能力文档](https://github.com/teambition/urbs-console/blob/develop/doc/adapter.md)