# 博客系统开发工具包

## 业务工具包 [infra](infra)

### 用户认证

- [captcha](infra/captcha) - 图形验证码生成与验证
- [jjwt](infra/jjwt) - JWT Token 签发与验证
- [oauth](infra/oauth) - 第三方登录集成（GitHub/Google/微信）

### 消息服务

- [mail](infra/mail) - 邮件发送服务
- [rabbitmq](infra/mq/rabbitmq) - 消息队列服务

### 文件服务

- [upload](infra/oss) - 文件上传与存储

## 代码生成工具 [quickstart](quickstart)

### 代码分析

- [astx](quickstart/astx) - AST 语法树分析工具
- [inject](quickstart/inject) - 代码注入工具
- [invent](quickstart/invent) - 模板代码生成工具

## 通用工具包 [util](util)

- 字符串处理
- 时间处理
- 加密解密
- 文件操作
- 类型转换
