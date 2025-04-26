# Gin 博客项目

## 项目简介

这是一个基于 Gin 框架开发的博客系统，采用分层架构设计，包含完整的博客功能模块。

## 项目结构

```
.
├── cmd/              # 项目启动命令
├── common/           # 项目通用文件
├── config/           # 配置文件
├── core/             # 核心运行组件
├── docs/             # Swagger 文档
├── initialize/       # 项目初始化
├── service/          # 服务接口处理逻辑
│   └── blog/         # 博客服务
│       ├── router/   # 路由层
│       ├── controller/ # 控制器层
│       ├── service/  # 服务层
│       ├── repository/ # 数据访问层
│       └── model/    # 数据模型层
└── svctx/            # 服务上下文，持有项目运行时资源
```

## 环境要求

- Go 1.23+
- MySQL 8.0+
- Redis 6.0+
- RabbitMQ 3.8+
- Nacos (可选，使用本地配置不需要)

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

复制配置文件：
```bash
cp config.default.yaml config.yaml
```

根据实际环境修改以下配置文件：

- `config.yaml`

### 3. 启动服务

#### 开发模式（使用本地配置）

3. 启动服务：

```bash
go run main.go api -c=file -f=./config.yaml
```

#### 生产模式（使用 Nacos 配置）

```bash
go run main.go api -c=nacos --n-namespace=prod
```

## 注意事项

1. 确保所有依赖服务（MySQL、Redis、RabbitMQ）正常运行
2. 开发环境建议使用本地配置文件
3. 生产环境建议使用 Nacos 配置中心
4. 数据库迁移前请确保已备份数据
