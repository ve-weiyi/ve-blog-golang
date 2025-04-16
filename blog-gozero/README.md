# GoZero 博客项目

## 项目简介

这是一个基于 GoZero 框架开发的博客系统，包含管理后台和博客前台两个主要模块。项目采用微服务架构，使用 RPC 进行服务间通信。

## 项目结构

```
.
├── .goctl/          # goctl 工具模板配置
├── internal/        # 公共代码模块
└── service/         # 服务代码
    ├── api/         # API 服务
    │   ├── admin/   # 管理后台服务
    │   └── blog/    # 博客前台服务
    ├── model/       # 数据库操作层
    └── rpc/         # RPC 服务
        └── blog/    # 博客 RPC 服务
```

## 环境要求

- Go 1.23+
- MySQL 8.0+
- Redis 6.0+
- RabbitMQ 3.8+
- Nacos (可选，使用本地配置不需要)
- ETCD (可选，直连 RPC 模式不需要)

## 快速开始

### 1. 环境准备

1. 安装 Go 环境

2. 启动依赖服务：

```bash
# 在 deploy/docker-compose/data 目录下执行
docker-compose up -d
```

3. 初始化数据库：
   - 执行 `blog-veweiyi-init.sql` 初始化表结构
   - 执行 `blog-veweiyi-data.sql` 导入初始数据

### 2. 配置修改

根据实际环境修改以下配置文件：

- `service/api/blog/etc/blog-api.yaml`
- `service/api/admin/etc/admin-api.yaml`
- `service/rpc/blog/etc/blog-rpc.yaml`

### 3. 启动服务

#### 开发模式（使用本地配置）

# 启动 RPC 服务

```bash
go run service/rpc/blog/blog.go -f service/rpc/blog/etc/blog-rpc.yaml
```

# 启动 API 服务

```bash
go run service/api/blog/blog.go -f service/api/blog/etc/blog-api.yaml
```

```bash
go run service/api/admin/admin.go -f service/api/admin/etc/admin-api.yaml
```

#### 生产模式（使用 Nacos 配置）

# 启动 RPC 服务

```bash
go run service/rpc/blog/blog.go
```

# 启动 API 服务

```bash
go run service/api/blog/blog.go
```

```bash
go run service/api/admin/admin.go
```

## 部署说明

使用 Docker Compose 进行部署：

```bash
docker-compose up -d -f docker-compose.yaml
```

## 注意事项

1. 确保所有依赖服务（MySQL、Redis、RabbitMQ）正常运行
2. 生产环境建议使用 Nacos 配置中心
3. 开发环境可以使用本地配置文件进行调试
