# go-zero 使用

[go-zero框架概述](https://go-zero.dev/docs/concepts/overview)
[goctl工具使用](https://go-zero.dev/docs/tasks/installation/goctl)

## 1. goctl 安装

1. 升级goctl 和 protoc，protoc-gen-go，protoc-gen-go-grpc
```sh
   GO111MODULE=on 
   go install github.com/zeromicro/go-zero/tools/goctl@latest
   goctl env check --install --verbose --force
```

2. 安装protoc-gen-go (建议使用goctl安装即可)
```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## 2. goctl 创建一个go-zero demo仓库

```
# 生成 snake case 文件和目录示例
goctl api new blog --style go_zero

goctl rpc new blog --style go_zero
```

## 3. goctl 高级用法
```
1. 替换代码模板
   goctl template init
   编辑 ~/.goctl/${goctl版本号}/api/handler.tpl

2. 生成api代码
   goctl api go -api blog.api -dir ../ --style go_zero
   
3. 格式化api代码
   goctl api format --dir blog.api
   
4. 生成api.ts代码
   goctl api ts -api blog.api -dir ../ts
```

## 4. goctl goctl-swagger生成文档

https://github.com/zeromicro/goctl-swagger

```
1. 编译goctl-swagger插件
   GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest

2. 使用goctl-swagger插件
   goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
```

# 项目启动

## 1.本地启动项目

启动etcd（可选,使用直连rpc模式时不需要）

```sh
etcd
```

启动rpc服务

```sh
go run service/rpc/blog/blog.go -f service/rpc/blog/etc/blog-rpc.yaml
```

启动api服务

```sh
go run service/api/blog/blog.go -f service/api/blog/etc/blog-api.yaml
```

## 2.使用nacos配置启动

```sh
go run service/rpc/blog/blog.go
```

```sh
go run service/api/blog/blog.go
```

## 3.使用docker启动

构造镜像

```sh
goctl docker --go service/rpc/blog/blog.go --exe blog
```
