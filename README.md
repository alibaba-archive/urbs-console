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
make run
```

浏览器访问：

>  http://localhost:8080


## Documentation (WIP)

[API 文档](https://github.com/teambition/urbs-console/blob/master/doc/api.md)

[集成到现有系统文档](https://github.com/teambition/urbs-console/blob/master/doc/adapter.md)
