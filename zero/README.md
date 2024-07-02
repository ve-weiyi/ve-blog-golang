# go-zero 使用

## 1.创建go-zero仓库

```
# 生成 snake case 文件和目录示例
goctl api new mini_power --style go_zero

goctl rpc new mini_power --style go_zero
```

## 2. goctl代码生成

```
安装grpc
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

1. 升级goctl
   GO111MODULE=on 
   go install github.com/zeromicro/go-zero/tools/goctl@latest

2. 替换代码模板
   goctl template init
   编辑 ~/.goctl/${goctl版本号}/api/handler.tpl

3. 生成api代码
   goctl api go -api blog.api -dir ../ --style go_zero

4. 生成api.ts代码
   goctl api ts -api blog.api -dir ../ts
   
5. 格式化api代码
   goctl api format --dir blog.api
```

## 3.goctl-swagger文档生成

https://github.com/zeromicro/goctl-swagger

```
1. 编译goctl-swagger插件
   GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest

2. 使用goctl-swagger插件
   goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
```

## 4.本地启动项目

```sh
etcd
```

```sh
go run service/blog/rpc/blog.go -f service/blog/rpc/etc/blog-rpc.yaml
```

```sh
go run service/blog/api/blog.go -f service/blog/api/etc/blog-api.yaml
```

```sh
goctl docker --go service/blog/rpc/blog.go --exe blog
```

## 5.使用nacos配置启动

```sh
go run service/blog/rpc/blog.go
```

```sh
go run service/blog/api/blog.go
```
 
构造镜像
```sh
docker build -t my-image .
```
