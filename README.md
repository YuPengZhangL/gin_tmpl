# gg_web_tmpl

## 简介

这是一个基于Gin和Gorm技术栈的Restful服务模板，包含一个目录结构与部分常用组件的封装，方便快速开发web项目使用，同时也可以供还在学习Golang的同学作为参考。本项目包含了一个用户登陆、注册、查询用户信息功能的 Restful API 接口，用于演示相关组件的使用，同时包含了完整的单元测试编写参考。

## 相关文档

- [接口文档](/doc/接口文档.md)

## 技术栈

- Web框架: [gin](https://gin-gonic.com/)
- ORM框架: [gorm](https://gorm.io/)
- 日志框架: [logrus](https://github.com/sirupsen/logrus)
- 数据库: [mysql](https://www.mysql.com/)
- Redis: [go-redis](https://github.com/go-redis/redis)

## 包含的内容

- 通过GORM对数据库的增查
- Redis的基本使用
- JWT鉴权及中间件
- RequestID封装
- 跨域中间件
- logrus日志框架相关封装
- gin和gorm接入logrus
- 配置文件的读取
- 完整的单元测试

## 快速启动

```shell
git clone https://github.com/kakkk/gg_web_tmpl.git
cd gg_web_tmpl
vi conf/config.yaml #修改相关配置
go mod tidy
go build
./gg_web_tmpl
```

## 单元测试

```shell
go test -v -gcflags=all=-l -coverprofile=cover.out ./...
go tool cover -html=cover.out #查看覆盖率
```

## 依赖

```text
github.com/DATA-DOG/go-sqlmock      #单测 mock 数据库
github.com/agiledragon/gomonkey/v2  #单测 mock 函数
github.com/dgrijalva/jwt-go         #jwt库
github.com/gin-gonic/gin            #gin框架
github.com/go-redis/redis/v8        #Redis客户端
github.com/go-redis/redismock/v8    #单测 mock Redis
github.com/gofrs/uuid               #uuid库
github.com/rifflock/lfshook         #日志写入文件Hook
github.com/sirupsen/logrus          #logrus日志框架
github.com/smartystreets/goconvey   #快速编写测试用例框架
github.com/spf13/cast               #类型转换库
golang.org/x/crypto                 #用于密码加密
gopkg.in/yaml.v3                    #yaml库
gorm.io/driver/mysql                #mysql驱动
gorm.io/gorm                        #gorm框架
```

## 目录&文件结构

```text
.
├── README.md
├── cache                                 #缓存组件
│         ├── redis_client.go             #redis客户端
│         └── user.go
├── common
│         ├── config                      #配置相关
│         │         ├── config.go
│         │         └── config_test.go
│         ├── consts                      #全局变量
│         │         └── consts.go
│         ├── log#日志相关
│         │         ├── gorm_logger.go          #gorm logger实现
│         │         └── logger.go               #logrus相关组件封装
│         ├── middleware                        #gin中间件
│         │         ├── auth_mw.go              #jwt鉴权中间件
│         │         ├── cros_mw.go              #跨域中间件
│         │         ├── gin_logger_mw.go        #gin日志记录中间件
│         │         └── request_id_mw.go        #reques_id生成中间件
│         ├── resp                              #通用Response封装
│         │         └── resp.go
│         └── utils                             #工具类
│             ├── is_in_gotest.go               #判断是否在单测环境
│             ├── jwt.go                        #jwt组件封装
│             └── password.go                   #密码加密相关封装
├── compose                                     #docker-compose文件
│         └── mysql-redis.yaml                  #MySQL Redis部署
├── conf                                        #配置文件
│         └── config.yaml
├── go.mod
├── go.sum
├── handler                     #handler层(controller)
│         └── user.go
├── main.go
├── model                       #model层
│         ├── mysql_client.go   #mysql客户端
│         └── user.go
├── router                      #gin路由
│         └── router.go
├── service                     #service层
│         └── user.go
└── vo                          #value object
    └── user.go

```