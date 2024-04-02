# go-zero 使用

## 1.创建go-zero仓库

```
# 生成 snake case 文件和目录示例
goctl api new mini_power --style go_zero

goctl rpc new mini_power --style go_zero
```

## 2. goctl代码生成

```
1. 升级goctl
   GO111MODULE=on go install github.com/zeromicro/go-zero/tools/goctl@latest

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

## 4.启动项目

```sh
etcd
```

```sh
go run rpc/blog.go -f rpc/etc/blog.yaml
```

```sh
go run api/blog.go -f api/etc/blog-api.yaml
```
