# go-zero 使用

[go-zero框架概述](https://go-zero.dev/docs/concepts/overview)
[goctl工具使用](https://go-zero.dev/docs/tasks/installation/goctl)

## 安装

### goctl 安装

1. 查看 go 版本

```sh
  go version
```

2. 安装/升级goctl

```sh
    go install github.com/zeromicro/go-zero/tools/goctl@latest
```

3. 验证

```sh
  goctl --version
```

### protoc 安装

1. 使用goctl 一键安装 protoc，protoc-gen-go，protoc-gen-go-grpc

```sh
  goctl env check --install --verbose --force
```

2. 验证

```sh
  goctl env check --verbose
```

protoc其他安装方式

```sh
  #   查看所有版本
   go list -m -versions google.golang.org/grpc/cmd/protoc-gen-go-grpc
   go list -m -versions google.golang.org/grpc/cmd/protoc
   
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#   安装
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   go install google.golang.org/grpc/cmd/protoc@latest
   
#   查看已安装版本
   protoc-gen-go-grpc --version
```

## 使用

### 1. goctl 创建一个go-zero demo仓库

执行指令生成一个 api 服务(使用下划线命名风格)

```sh
  goctl api new blog --style go_zero
```

执行指令生成一个 rpc 服务(使用下划线命名风格)

```sh
  goctl rpc new blog --style go_zero
```

### 2. goctl 高级用法

#### api 指令

1. 生成api代码

```sh
  goctl api go -api blog.api -dir ../ --style go_zero
```

2. 格式化api代码

```sh
  goctl api format --dir blog.api
```

3. 生成api.ts代码

```sh
  goctl api ts -api blog.api -dir ../ts
```

#### docker 指令

构造镜像

```sh
  goctl docker --go service/rpc/blog/blog.go --exe blog
```

#### template 指令

初始化代码模板

```sh
  goctl template init
```

生成的模板存储于 ~/.goctl/${goctl版本号}
编辑 ~/.goctl/${goctl版本号}/api/handler.tpl 可替换生成代码

#### 插件

生成swagger文档

https://github.com/zeromicro/goctl-swagger

1. goctl 插件安装

```sh
  GOPROXY=https://goproxy.cn/,direct 
  go install github.com/zeromicro/goctl-swagger@latest
```

2. 使用goctl-swagger插件

```sh
  goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
```
