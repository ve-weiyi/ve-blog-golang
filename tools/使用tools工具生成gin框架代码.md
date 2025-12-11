# Tools - 代码生成工具

ve-blog-golang 项目的代码生成工具集，用于从 API 定义文件自动生成后端、前端和数据库模型代码，提高开发效率。

## 功能概述

| 功能         | 说明                          | 生成内容                       |
|------------|-----------------------------|----------------------------|
| API 代码生成   | 从 `.api` 文件生成 Gin 框架代码      | types/logic/handler/router |
| Model 代码生成 | 从数据库或 SQL 文件生成 GORM 模型      | 数据库模型文件                    |
| Web 代码生成   | 从 `.api` 文件生成 TypeScript 代码 | API 接口调用代码和类型定义            |

## 快速开始

### 使用 Makefile（推荐）

```bash
# 查看所有可用命令
make help

# 生成 Gin 框架代码
make gen-gin              # 生成所有 Gin 代码
make gen-gin-blog         # 仅生成博客前台代码
make gen-gin-admin        # 仅生成管理后台代码

# 生成数据库模型代码
make gen-model

# 生成 TypeScript 代码
make gen-ts               # 生成所有 TS 代码
make gen-ts-blog          # 仅生成博客前台 TS 代码
make gen-ts-admin         # 仅生成管理后台 TS 代码

# 生成所有代码
make gen-all

# 清理生成的代码
make clean
```

### 使用命令行

**生成 API 代码**

```bash
# 生成 types 文件
go run main.go api types -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -o='../blog-gin/api/blog' -n='%s.go'

# 生成 logic 文件
go run main.go api logic -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/blog' -n='%s.lg.go'

# 生成 handler 文件
go run main.go api handler -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/blog' -n='%s.hdl.go'

# 生成 router 文件
go run main.go api router -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/blog' -n='%s.rt.go'
```

**生成 Model 代码**

```bash
# 从数据库生成
go run main.go model dsn -t=./template/go-zero/model.tpl -n='%v_model.go' -o='./runtime/model' -s='root:password@(host:3306)/database?charset=utf8mb4&parseTime=True&loc=Local'

# 从 SQL 文件生成
go run main.go model ddl -t=./template/go-zero/model.tpl -n='%v_model.go' -o='./runtime/model' -s='./testdata/test.sql'
```

**生成 TypeScript 代码**

```bash
go run main.go web typescript -n='%v.ts' -t='./template/web' -o='./runtime/blog/api' -m='api' -i='IApiResponse' -f='../blog-gozero/service/api/blog/proto/blog.api'
```

## 命令参数说明

| 参数   | 说明                 | 示例                                                  |
|------|--------------------|-----------------------------------------------------|
| `-f` | 输入文件路径             | `../blog-gozero/service/api/blog/proto/blog.api`    |
| `-t` | 模板目录路径             | `./template/gin`                                    |
| `-o` | 输出目录路径             | `../blog-gin/api/blog`                              |
| `-n` | 输出文件名格式            | `%s.go` / `%v_model.go`                             |
| `-c` | 上下文包路径             | `github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx` |
| `-s` | 数据库连接字符串或 SQL 文件路径 | `root:password@(host:3306)/database`                |
| `-m` | 模块名称               | `api`                                               |
| `-i` | 接口响应类型             | `IApiResponse`                                      |

## 目录结构

```
tools/
├── cmd/                    # 命令行工具实现
│   ├── api/               # API 代码生成命令
│   ├── model/             # Model 代码生成命令
│   └── web/               # Web 代码生成命令
├── parserx/               # 解析器
│   ├── apiparser/         # API 文件解析器
│   └── swagparser/        # Swagger 解析器
├── template/              # 代码模板
│   ├── gin/              # Gin 框架模板
│   ├── go-zero/          # Go-zero 框架模板
│   └── web/              # Web 前端模板
├── testdata/              # 测试数据
├── Makefile               # 构建脚本
├── generate_gin.sh        # Gin 代码生成脚本
```

## 工作原理

### API 解析器

**ApiParser** - 使用 Go AST 语法树解析 `.api` 文件

- 提取接口路径、方法、参数、返回值
- 生成对应的 types、logic、handler、router 代码

**SwaggerParser** - 解析 `swagger.json` 文件

- 提取接口定义信息
- 支持从 Swagger 文档生成代码

### 代码生成流程

1. 解析输入文件（`.api` 或 `swagger.json`）
2. 提取接口定义信息
3. 根据模板生成目标代码
4. 输出到指定目录

## 技术特点

- ✅ 基于 Cobra 构建命令行工具
- ✅ 支持自定义模板，灵活扩展
- ✅ 支持多种代码生成场景
- ✅ 自动化生成，减少重复代码编写
- ✅ 统一代码风格，提高代码质量

## 使用场景

1. **新增接口** - 修改 `.api` 文件后，运行 `make gen-gin` 自动生成代码
2. **数据库变更** - 修改 SQL 文件后，运行 `make gen-model` 更新模型
3. **前后端协作** - 运行 `make gen-ts` 生成前端 API 调用代码

## 注意事项

1. 生成代码前建议先备份或提交现有代码
2. 生成的代码可能需要手动调整部分逻辑
3. 模板文件位于 `template/` 目录，可根据需求自定义
4. 生成的文件会覆盖同名文件，请谨慎操作
