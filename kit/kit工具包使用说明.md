# 工具包使用说明

本目录包含博客系统开发所需的各类工具包，分为业务工具包、代码生成工具和通用工具包三大类。

## 📦 业务工具包 [infra](infra)

### 认证授权

| 工具包                      | 说明     | 主要功能                     |
|--------------------------|--------|--------------------------|
| [captcha](infra/captcha) | 图形验证码  | 验证码生成、Redis 存储、验证        |
| [jwtx](infra/jwtx)       | JWT 认证 | Token 签发、解析、验证           |
| [oauth](infra/oauth)     | 第三方登录  | GitHub/QQ/微信/微博/飞书/Gitee |

### 消息通信

| 工具包                | 说明        | 主要功能                       |
|--------------------|-----------|----------------------------|
| [mail](infra/mail) | 邮件服务      | 邮件发送、模板渲染、消息队列集成           |
| [mq](infra/mq)     | 消息队列      | RabbitMQ/Kafka/Redis MQ 封装 |
| [ws](infra/ws)     | WebSocket | 实时通信、聊天室                   |

### 存储服务

| 工具包                  | 说明   | 主要功能             |
|----------------------|------|------------------|
| [oss](infra/oss)     | 对象存储 | 阿里云 OSS/七牛云/本地存储 |
| [nacos](infra/nacos) | 配置中心 | Nacos 配置读取与监听    |

### 数据处理

| 工具包                            | 说明       | 主要功能         |
|--------------------------------|----------|--------------|
| [excel](infra/excel)           | Excel 处理 | 导入导出、数据解析    |
| [gsm](infra/gsm)               | 数据库迁移    | SQL 生成、表结构管理 |
| [gormlogger](infra/gormlogger) | GORM 日志  | SQL 日志格式化    |

### 其他服务

| 工具包                      | 说明         | 主要功能           |
|--------------------------|------------|----------------|
| [chatgpt](infra/chatgpt) | ChatGPT 集成 | AI 对话、提示词管理    |
| [music](infra/music)     | 音乐服务       | 网易云音乐 API 封装   |
| [logz](infra/logz)       | 日志服务       | Zap 日志封装       |
| [random](infra/random)   | 随机生成       | UID/UUID 生成    |
| [biz](infra/biz)         | 业务错误       | 业务异常、HTTP 错误封装 |

## 🛠️ 代码生成工具 [quickstart](quickstart)

| 工具包                         | 说明      | 主要功能          |
|-----------------------------|---------|---------------|
| [astx](quickstart/astx)     | AST 语法树 | Go 代码解析、语法树操作 |
| [inject](quickstart/inject) | 代码注入    | 自动注入导入、函数、结构体 |
| [invent](quickstart/invent) | 模板生成    | 基于模板快速生成代码    |

**参考文档**：

- [Go 抽象语法树和 ast 库](../assets/article/guide/Go-抽象语法树和ast库.md)
- [快速开始指南](quickstart/README.md)

## 🔧 通用工具包 [utils](utils)

### 加密解密

| 工具包                    | 说明   | 支持算法                                |
|------------------------|------|-------------------------------------|
| [crypto](utils/crypto) | 加密工具 | AES/RSA/ECC/ECDSA/MD5/SHA256/Bcrypt |

### 数据处理

| 工具包                        | 说明      | 主要功能                  |
|----------------------------|---------|-----------------------|
| [jsonconv](utils/jsonconv) | JSON 转换 | 驼峰/蛇形转换、JSON 序列化      |
| [slicex](utils/slicex)     | 切片操作    | 去重、过滤、映射              |
| [mark](utils/mark)         | 数据脱敏    | 手机号/邮箱/身份证脱敏          |
| [typecase](utils/typecase) | 类型转换    | MySQL/TypeScript 类型映射 |

### 网络工具

| 工具包                  | 说明       | 主要功能                |
|----------------------|----------|---------------------|
| [httpx](utils/httpx) | HTTP 客户端 | 请求封装、链式调用           |
| [ipx](utils/ipx)     | IP 查询    | IP 归属地查询（百度/IP-API） |

### 文件操作

| 工具包                  | 说明   | 主要功能      |
|----------------------|------|-----------|
| [filex](utils/filex) | 文件工具 | 文件读写、压缩解压 |

### 数据库

| 工具包              | 说明    | 主要功能         |
|------------------|-------|--------------|
| [dbx](utils/dbx) | 数据库工具 | MySQL 连接、表清理 |

### 其他工具

| 工具包                        | 说明   | 主要功能       |
|----------------------------|------|------------|
| [color](utils/color)       | 终端颜色 | 彩色输出（跨平台）  |
| [patternx](utils/patternx) | 模式匹配 | 正则表达式、版本检查 |
| [timer](utils/timer)       | 定时任务 | 定时器封装      |
| [tempx](utils/tempx)       | 模板工具 | 模板渲染       |
| [system](utils/system)     | 系统工具 | 服务管理、热重载   |

## 📚 使用示例

### JWT 认证

```go
import "github.com/ve-weiyi/ve-blog-golang/kit/infra/jwtx"

// 生成 Token
token, _ := jwtx.GenerateToken(userId, secret, expireTime)

// 解析 Token
claims, _ := jwtx.ParseToken(token, secret)
```

### 图形验证码

```go
import "github.com/ve-weiyi/ve-blog-golang/kit/infra/captcha"

// 生成验证码
id, b64s, _ := captcha.Generate()

// 验证
ok := captcha.Verify(id, code)
```

### 文件上传

```go
import "github.com/ve-weiyi/ve-blog-golang/kit/infra/oss"

// 上传文件
url, _ := oss.UploadFile(file, "aliyun")
```

### 数据脱敏

```go
import "github.com/ve-weiyi/ve-blog-golang/kit/utils/mark"

// 手机号脱敏
phone := mark.MaskPhone("13800138000") // 138****8000
```

## 🔗 相关文档

- [消息队列使用说明](infra/mq/消息队列.md)
- [数据库迁移工具](infra/gsm/README.md)
- [数据脱敏说明](utils/mark/README.md)

## 📝 注意事项

1. 使用前请先安装依赖：`go mod tidy`
2. 部分工具需要配置相关服务（如 Redis、MySQL）
3. 第三方服务需要配置对应的 AppKey 和 Secret
4. 生产环境请注意密钥安全，不要硬编码在代码中
