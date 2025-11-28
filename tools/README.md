# Tools - 代码生成工具

ve-blog-golang 项目的代码生成工具集，用于自动化生成各类代码文件，提高开发效率。

## 功能概述

本工具提供三大核心功能：

### 1. API 代码生成 (api)

从 `.api` 文件生成 Gin 框架的后端代码，包括：

- **types** - 生成数据类型定义文件
- **logic** - 生成业务逻辑层代码
- **handler** - 生成控制器层代码
- **router** - 生成路由配置代码

### 2. Model 代码生成 (model)

从数据库或 SQL 文件生成 GORM 模型代码：

- **dsn** - 从数据库连接生成模型（支持线上/本地数据库）
- **ddl** - 从 SQL DDL 文件生成模型

### 3. Web 前端代码生成 (web)

从 `.api` 文件生成前端 TypeScript 代码：

- **typescript** - 生成 API 接口调用代码和类型定义

## 使用方式

### API 代码生成示例

```bash
# 生成 types 文件
go run main.go api types \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-o='../blog-gin/service/api' \
-n='%s.go'
```

### Model 代码生成示例

```bash
# 从数据库生成
go run main.go model dsn \
-t=./template/go-zero/model.tpl \
-n='%v_model.go' \
-o='./runtime/model' \
-s='root:password@(host:3306)/database?charset=utf8mb4&parseTime=True&loc=Local'

# 从 SQL 文件生成
go run main.go model ddl \
-t=./template/go-zero/model.tpl \
-n='%v_model.go' \
-o='./runtime/model' \
-s='./testdata/test.sql'
```

### TypeScript 代码生成示例

```bash
go run main.go web typescript \
-n='%v.ts' \
-t='./template/web' \
-o='./runtime/blog/api' \
-m='api' \
-i='IApiResponse' \
-f='../blog-gozero/service/api/blog/proto/blog.api'
```

## 快捷脚本

项目提供了便捷的生成脚本：

- `generate_gin.sh` - 一键生成 Gin 框架代码
- `generate_model.sh` - 一键生成数据库模型
- `generate_typescript.sh` - 一键生成前端 TypeScript 代码

## 目录结构

```
tools/
├── cmd/              # 命令行工具实现
│   ├── api/         # API 代码生成
│   ├── model/       # Model 代码生成
│   └── web/         # Web 前端代码生成
├── parserx/         # 解析器
│   ├── apiparser/   # API 文件解析器
│   └── swagparser/  # Swagger 解析器
├── template/        # 代码模板
│   ├── gin/        # Gin 框架模板
│   ├── go-zero/    # Go-zero 框架模板
│   └── web/        # Web 前端模板
└── testdata/        # 测试数据
```

## 参数说明

- `-f` - 输入文件路径
- `-t` - 模板目录路径
- `-o` - 输出目录路径
- `-n` - 输出文件名格式
- `-c` - 上下文包路径
- `-s` - 数据库连接字符串或 SQL 文件路径
- `-m` - 模块名称
- `-i` - 接口响应类型

## 技术特点

- 基于 Cobra 构建命令行工具
- 支持自定义模板
- 支持多种代码生成场景
- 提高开发效率，减少重复代码编写
