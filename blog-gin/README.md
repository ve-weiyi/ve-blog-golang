# 项目介绍

## 1.项目结构

```shell
├── cmd 项目启动命令
├── common 项目通用文件
├── config 配置文件
├── core 核心运行组件
├── docs swagger文档
├── initialize 项目初始化
├── service 服务接口处理逻辑,根据 router,controller,service,repository,model 分层
│ └── blog 博客服务
│    ├── router
│    └── controller
│    ├── service
│    └── repository
│    ├── model
├── svctx service_context服务上下文,持有所有项目运行时的资源
```

## 2.初始化数据库

准备工作：

1. 确认已经安装好golang环境

2. 确保运行了 mysql、redis、rabbitmq 等服务。在[data](../deploy/docker-compose/data)目录下输入命令`docker-compose up -d`
   快速启动这些服务。

- 初始化数据库,创建表和数据

```shell
go run main.go migrate --action=migrate \
--file=./blog-veweiyi-init.sql \
--host=127.0.0.1 \
--port=3307 \
--username=root \
--password=mysql7914 \
--name='blog-veweiyi'  
```

## 3.启动项目

### 使用nacos配置启动(线上环境)

使用nacos的配置文件启动项目

```shell
go run main.go api -c=nacos --n-namespace=test
```

### 项目本地启动(本地环境)

**修改配置文件**

在server目录下新增config.yaml文件，修改config.yaml配置文件中的配置信息，文件内容参考[config.default.yaml](server/config.default.yaml)

使用本地配置文件 config.yaml 启动项目

```shell
go run main.go api -c=file -f=./config.yaml
```
