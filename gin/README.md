# 项目介绍


## 1.项目结构

```shell
├── cmd 项目启动命令
├── config 配置文件
├── core 核心运行组件
├── docs swagger文档
├── infra 项目基础设施
├── initialize 项目初始化
├── service 服务接口处理逻辑,根据 router,controller,service,repository,model 分层
│ └── blog 博客服务
│    ├── router
│    └── controller
│    ├── service
│    └── repository
│    ├── model
├── svcctx service_context服务上下文,持有所有项目运行时的资源
```

## 2.启动项目

准备工作：
1. 确认已经安装好golang环境

2. 确保运行了 mysql、redis、rabbitmq 等服务。在[data](../deploy/docker-compose/data)目录下输入命令`docker-compose up -d`快速启动这些服务。

- 初始化数据库,创建表和数据
```shell
go run main.go migrate --action=migrate \
--file=./blog-veweiyi.sql \
--host=127.0.0.1 \
--port=3307 \
--username=root \
--password=mysql7914 \
--name='blog-veweiyi'  
```

3. 修改配置文件
在server目录下新增config.yaml文件，修改config.yaml配置文件中的配置信息，文件内容参考[config.default.yaml](server/config.default.yaml)

### 使用nacos配置启动
使用nacos的配置文件启动项目
```shell
go run main.go api -c=nacos --n-namespace=test
```

### 项目本地启动

使用本地配置文件 config.yaml 启动项目
```shell
go run main.go api -c=file -f=./config.yaml
```

## 3.服务部署


## 4.其他

### 格式化代码
格式化代码
```shell
go fmt ./...
```

格式化导入包
```shell
go install golang.org/x/tools/cmd/goimports
goimports -w .
```

### 生成swagger文档
安装swag
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

生成swagger注解
```shell
swag init
```

格式化swagger注解

```shell
swag fmt
```
