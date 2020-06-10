# KarlMalone

一个可以实现高并发的IM api项目

## 内容列表

- [使用技术](#使用技术)
- [安装](#安装)
- [项目结构](#项目结构)
- [相关仓库](#相关仓库)
- [维护者](#维护者)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可)

## 使用技术

数据库: mysql + [Gorm](https://github.com/go-gorm/gorm)

日志框架: [Zap](https://github.com/uber-go/zap)

映射配置: [Viper](https://github.com/spf13/viper)

## 安装

1. 首先安装MySQL
2. 创建数据库karlmalone，执行sql/create_table.sql，完成初始化表的创建
3. 修改config文件夹下配置文件config.yml，使之和你本地配置一致
4. 执行以下命令运行该api项目

    ```sh
    $ go build
    $ ./KarlMalone
    ```

## 项目结构
项目结构遵循 https://github.com/golang-standards/project-layout
```
    .
    ├── api
    ├── assets
    │   ├── attach
    │   ├── css
    │   ├── fonts
    │   ├── images
    │   ├── js
    │   └── plugins
    ├── build
    │   ├── ci
    │   └── package
    ├── cmd
    ├── config
    ├── docs
    ├── internal
    │   ├── args
    │   ├── controller
    │   ├── db
    │   ├── model
    │   ├── service
    │   └── view
    ├── log
    │   └── event
    ├── pkg
    │   ├── logger
    │   └── util
    ├── resources
    ├── sql
    └── third_party
```

api:          服务对外提供的grpc接口(待写)

config:       服务配置

internal:     服务私有代码

pkg:          服务共有代码

sql:          项目sql文件

## 相关仓库

- [gim](https://github.com/alberliu/gim) — golang写的IM服务器，支持多业务接入。
- [fastIM](https://github.com/GuoZhaoran/fastIM) — an example of im system。

## 维护者

[@UncleMaple](https://github.com/UncleMaple)

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/UncleMaple/KarlMalone/issues/new) 或者提交一个 Pull Request。


KarlMalone 遵循 [Contributor Covenant](http://contributor-covenant.org/version/1/3/0/) 行为规范。

## 使用许可

[MIT]() © Drew Lee
