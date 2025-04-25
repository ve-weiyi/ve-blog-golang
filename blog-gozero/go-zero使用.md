# Go-Zero 框架使用指南

## 官方文档

- [Go-Zero 框架概述](https://go-zero.dev/docs/concepts/overview)
- [Goctl 工具使用](https://go-zero.dev/docs/tasks/installation/goctl)

## 环境安装

### 1. Goctl 安装

```bash
# 查看 Go 版本
go version
```

```bash
# 安装/升级 Goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

```bash
# 验证安装
goctl --version
```

### 2. Protoc 安装

推荐使用 Goctl 一键安装：

```bash
# 安装 protoc 及相关工具
goctl env check --install --verbose --force
```

```bash
# 验证安装
goctl env check --verbose
```

手动安装方式：

```bash
# 查看可用版本
go list -m -versions google.golang.org/grpc/cmd/protoc-gen-go-grpc
go list -m -versions google.golang.org/grpc/cmd/protoc
```

```bash
# 安装 protoc 工具
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/grpc/cmd/protoc@latest
```

```bash
# 验证安装
protoc-gen-go-grpc --version
```

## 项目创建

### 1. 创建新项目

```bash
# 创建 API 服务（使用下划线命名风格）
goctl api new blog --style go_zero
```

```bash
# 创建 RPC 服务（使用下划线命名风格）
goctl rpc new blog --style go_zero
```

## Goctl 工具使用

### 1. 模板管理

```bash
# 初始化代码模板
goctl template init
```

- 模板位置：`~/.goctl/${goctl版本号}`
- 可编辑 `~/.goctl/${goctl版本号}/api/handler.tpl` 自定义生成代码

### 2. API 相关命令

```bash
# 生成 API 代码
goctl api go -api blog.api -dir ../ --style go_zero
```

```bash
# 格式化 API 代码
goctl api format --dir blog.api
```

```bash
# 生成 TypeScript 代码
goctl api ts -api blog.api -dir ../ts
```

### 3. Docker 相关命令

```bash
# 构建 Docker 镜像
goctl docker --go service/rpc/blog/blog.go --exe blog
```

### 4. 插件使用

#### Swagger 文档生成

1. 安装插件：

```bash
# 设置 GOPROXY
export GOPROXY=https://goproxy.cn/,direct
```

```bash
# 安装 goctl-swagger 插件
go install github.com/zeromicro/goctl-swagger@latest
```

2. 使用插件：

```bash
# 生成 Swagger 文档
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
```

## 最佳实践

1. 建议使用 Goctl 一键安装工具，避免手动安装可能出现的版本兼容问题
2. 遵循 Go-Zero 的命名规范，保持代码风格统一
