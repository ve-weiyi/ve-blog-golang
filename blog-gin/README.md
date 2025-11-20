# Gin 博客项目

## 项目简介

这是一个基于 Gin 框架开发的博客系统，采用分层架构设计，包含完整的博客功能模块。项目提供了两个主要命令：`api` 用于启动博客服务，
`migrate` 用于数据库初始化。

## 项目结构

```
.
├── cmd/              # 项目启动命令
│   ├── api.go        # API服务启动命令
│   ├── migrate.go    # 数据库迁移命令
│   └── root.go       # 根命令
├── common/           # 项目通用文件
├── config/           # 配置文件
├── core/             # 核心运行组件
├── docs/             # Swagger 文档
├── initialize/       # 项目初始化
├── service/          # 服务接口处理逻辑
│   ├── admin/        # 后台管理服务
│   └── blog/         # 博客服务
│       ├── router/   # 路由层
│       ├── controller/ # 控制器层
│       ├── service/  # 服务层
│       └── dto/      # 数据传输对象
└── svctx/            # 服务上下文，持有项目运行时资源
```

## 环境要求

- Go 1.23+
- MySQL 8.0+
- Redis 6.0+
- RabbitMQ 3.8+
- Nacos (可选，使用本地配置不需要)

## 命令使用说明

### 1. API 服务命令 (`api`)

`api` 命令用于启动博客的 HTTP API 服务，支持本地配置文件和 Nacos 配置中心两种配置方式。

#### 基本用法

```bash
# 使用本地配置文件启动（推荐开发环境）
go run main.go api -c file -f ./config.yaml

# 使用 Nacos 配置中心启动（推荐生产环境）
go run main.go api -c nacos --n-namespace prod

# 直接运行（使用默认配置）
go run main.go api
```

#### 参数说明

**配置相关参数：**

- `-c, --config`: 配置读取方式 (file|nacos)，默认: file
- `-f, --filepath`: 本地配置文件路径，默认: config.yaml

**Nacos 配置参数：**

- `--n-host`: Nacos 服务器地址，默认: veweiyi.cn
- `--n-port`: Nacos 服务器端口，默认: 8848
- `--n-namespace`: Nacos 命名空间，默认: dev
- `--n-data-id`: Nacos 配置 DataId，默认: ve-blog-golang
- `--n-group`: Nacos 配置分组，默认: blog
- `--n-user`: Nacos 用户名，默认: nacos
- `--n-password`: Nacos 密码，默认: nacos

#### 使用示例

```bash
# 开发环境 - 使用本地配置
go run main.go api -c file -f ./config.yaml

# 测试环境 - 使用 Nacos 配置
go run main.go api -c nacos --n-namespace test --n-host nacos.test.com

# 生产环境 - 使用 Nacos 配置
go run main.go api -c nacos --n-namespace prod --n-host nacos.prod.com --n-user admin --n-password admin123
```

### 2. 数据库迁移命令 (`migrate`)

`migrate` 命令用于初始化数据库，包括创建数据库、执行表结构 SQL 和导入初始数据。

#### 基本用法

```bash
# 使用默认配置进行数据库迁移
go run main.go migrate

# 指定 SQL 文件进行迁移
go run main.go migrate -i blog-veweiyi-init.sql -d blog-veweiyi-data.sql

# 指定数据库连接参数
go run main.go migrate --host localhost --port 3306 --username root --password 123456 --name blog
```

#### 参数说明

**SQL 文件参数：**

- `-i, --file`: 数据库结构 SQL 文件路径，默认: blog-veweiyi-init.sql
- `-d, --data`: 数据库初始数据 SQL 文件路径，默认: blog-veweiyi-data.sql

**数据库连接参数：**

- `--host`: 数据库主机地址，默认: localhost
- `--port`: 数据库端口，默认: 3306
- `--username`: 数据库用户名，默认: root
- `--password`: 数据库密码，默认: 123456
- `--name`: 数据库名称，默认: blog-veweiyi
- `--config`: 数据库连接配置，默认: charset=utf8mb4&parseTime=True&loc=Local

#### 使用示例

```bash
# 使用默认配置初始化数据库
go run main.go migrate

# 指定数据库连接信息
go run main.go migrate --host 192.168.1.100 --port 3306 --username admin --password admin123 --name myblog

# 只执行表结构初始化，不导入数据
go run main.go migrate -i blog-veweiyi-init.sql -d ""

# 使用自定义 SQL 文件
go run main.go migrate -i custom-schema.sql -d custom-data.sql
```

#### 注意事项

⚠️ **重要警告**：`migrate` 命令会先删除指定的数据库，然后重新创建，请确保：

1. 在执行前备份重要数据
2. 确认数据库名称正确
3. 数据库用户具有 DROP、CREATE 权限

## 快速开始

### 1. 环境准备

1. 安装 Go 环境

2. 启动依赖服务：

```bash
# 在 deploy/docker-compose/data 目录下执行
docker-compose up -d
```

### 2. 配置修改

复制配置文件：

```bash
cp config.default.yaml config.yaml
```

根据实际环境修改配置文件中的数据库、Redis、RabbitMQ 等连接信息。

### 3. 数据库初始化

```bash
# 初始化数据库（会删除并重新创建数据库）
go run main.go migrate
```

```bash
# 或者指定数据库连接参数
go run main.go migrate --host localhost --username root --password yourpassword --name blog
```

### 4. 启动服务

```bash
# 启动 API 服务
go run main.go api
```

```bash
# 或者指定配置文件
go run main.go api -f ./config.yaml
```

### 5. 访问服务

- API 文档地址：http://localhost:8080/api/v1/swagger/index.html
- 健康检查：http://localhost:8080/api/v1/version

## 部署建议

### 开发环境

- 使用本地配置文件：`go run main.go api -c file -f ./config.yaml`
- 数据库迁移：`go run main.go migrate`

### 生产环境

- 使用 Nacos 配置中心：`go run main.go api -c nacos --n-namespace prod`
- 数据库迁移：建议手动执行 SQL 文件，避免使用 migrate 命令

## 注意事项

1. 确保所有依赖服务（MySQL、Redis、RabbitMQ）正常运行
2. 开发环境建议使用本地配置文件
3. 生产环境建议使用 Nacos 配置中心
4. 数据库迁移前请确保已备份数据
5. migrate 命令会重置数据库，生产环境请谨慎使用
