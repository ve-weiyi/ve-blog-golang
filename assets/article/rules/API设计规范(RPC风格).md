# API 设计规范（RPC 风格）

采用 **RPC 风格**：URL 表达动作（动词 + 名词），查询类用 GET，其余统一用 POST，参数放 Body。

**选择理由：**

- 接口名即文档，URL 直接表达"做什么"，调用方无需理解资源模型
- 适合大量非 CRUD 操作（触发、重放、执行、导出等），无需纠结 REST 资源建模
- 团队上手快，接口语义零歧义
- 国内部分主流平台（如抖音开放平台）采用类似的 RPC 风格，URL 直接表达动作语义

> 核心原则：**URL 代表「调用哪个方法」，而非「操作哪个资源」。**

---

## 目录

- [1. 设计原则](#1-设计原则)
- [2. URL 设计](#2-url-设计)
- [3. HTTP 方法](#3-http-方法)
- [4. 请求设计](#4-请求设计)
- [5. 响应设计](#5-响应设计)
- [6. 错误处理](#6-错误处理)
- [7. 认证与鉴权](#7-认证与鉴权)
- [8. 分页](#8-分页)
- [9. 排序与过滤](#9-排序与过滤)
- [10. 版本管理](#10-版本管理)
- [11. 安全规范](#11-安全规范)
- [12. 反模式清单](#12-反模式清单)
- [13. 监控与告警](#13-监控与告警)
- [附录A：完整请求/响应示例](#附录a完整请求响应示例)
- [参考文献](#参考文献)

---

## 1. 设计原则

### 1.1 动作为中心

RPC 风格以**动作**为中心，URL 直接描述操作，而非间接通过 HTTP 方法表达语义。

```
RPC 思路: "我要创建一篇文章 → POST /api/v1/article/create_article"
REST 思路: "我要操作文章资源 → POST /api/v1/articles"
```

RPC 风格 URL 自带文档属性——看到路径就清楚功能，无需查阅 HTTP 方法映射规则。

### 1.2 核心约束

| 原则 | 说明 |
|------|------|
| **URL 即文档** | 路径直接表达动作，语义零歧义 |
| **统一 POST** | 所有非查询操作使用 POST，参数统一放 Body |
| **GET 可缓存** | 查询接口用 GET，可缓存、可重试、幂等 |
| **无状态** | 每个请求包含完成所需的全部信息 |
| **可演进** | 新增字段、新增接口不破坏已有调用方 |

### 1.3 与 REST 风格对比

| 维度 | REST 风格 | RPC 风格（本规范） |
|------|----------|-----------------|
| 核心理念 | 操作资源 | 调用方法 |
| URL 格式 | 名词（`/users`） | 动词 + 名词（`/article/create_article`） |
| HTTP 方法 | GET/POST/PUT/PATCH/DELETE | GET（查询）+ POST（其他） |
| 语义来源 | HTTP 方法 + URL 资源 | URL 直接表达 |
| 学习成本 | 需要理解 REST 资源模型 | 看 URL 即懂 |
| 非 CRUD 操作 | 需要自定义命名映射 | 自然表达（`/order/cancel_order`） |
| 适用场景 | 资源 CRUD 为主 | 操作类接口多的业务系统 |

---

## 2. URL 设计

### 2.1 核心格式

```
{Method}  {BasePath}/{Version}/{Module}/{Action}

GET    /api/v1/article/get_article
POST   /api/v1/article/create_article
POST   /api/v1/article/update_article
POST   /api/v1/article/delete_article
POST   /api/v1/article/query_article_list
```

### 2.2 命名规则

| 规则 | 正确 | 错误 |
|------|------|------|
| **动词在前，名词在后** | `create_article`、`get_user` | `article_create`、`user_get` |
| 全小写 + 下划线 | `get_article_list`、`user_profile` | `getArticleList`、`user-profile` |
| 动宾结构 | `create_order`、`send_notification` | `order`、`notification` |
| 按模块分组前缀 | `/article/get_article`、`/user/get_user` | 全放根路径 |
| 版本号在路径中 | `/api/v1/article/create_article` | `/api/article/create_article?v=1` |

> URL 中使用下划线 `_` 而非连字符 `-`，原因是 RPC 方法名本身就是 identifier（snake_case），保持一致性。

### 2.3 动词词汇表

| 动词 | 含义 | 示例 |
|------|------|------|
| `get` | 获取单个 | `get_article`、`get_user` |
| `query` | 查询列表 | `query_article_list`、`query_order_list` |
| `create` | 创建 | `create_article`、`create_order` |
| `update` | 更新 | `update_article`、`update_config` |
| `delete` | 删除 | `delete_article`、`delete_comment` |
| `batch_delete` | 批量删除 | `batch_delete_article` |
| `send` | 发送 | `send_notification`、`send_email` |
| `verify` | 验证 | `verify_email`、`verify_phone` |
| `publish` | 发布 | `publish_article` |
| `revoke` | 撤回 | `revoke_article` |
| `cancel` | 取消 | `cancel_order` |
| `export` | 导出 | `export_report` |
| `import` | 导入 | `import_data` |
| `archive` | 归档 | `archive_document` |

### 2.4 动词选择优先级

对于同一资源的不同读取方式：

```bash
GET  /api/v1/article/get_article         # 按 ID 获取单个
GET  /api/v1/article/query_article       # 按条件查询单个（返回第一个匹配）
GET  /api/v1/article/query_article_list  # 分页查询列表
GET  /api/v1/article/search_article      # 全文搜索
```

- `get` → 按主键获取，确定返回 0 或 1 条
- `query` → 按条件查询，可能是单个或列表
- `search` → 全文搜索，走搜索引擎

### 2.5 模块分组

按业务模块加前缀组织，避免 URL 命名冲突：

```bash
/api/v1/article/
├── get_article
├── query_article_list
├── create_article
├── update_article
├── delete_article
├── publish_article
└── archive_article

/api/v1/user/
├── get_user
├── create_user
├── update_user
└── query_user_list

/api/v1/order/
├── get_order
├── create_order
├── cancel_order
└── query_order_list
```

层级不宜超过 3 层：`/api/{version}/{module}/{action}`。（如有子模块：`/api/{version}/{module}/{sub_module}/{action}`）

---

## 3. HTTP 方法

### 3.1 方法约定

| 场景 | HTTP 方法 | 原因 |
|------|----------|------|
| 查询（无副作用，不修改数据） | `GET` | 可缓存、可重试、幂等 |
| 其他所有操作 | `POST` | 统一、Body 可携带复杂参数 |

### 3.2 GET 接口约束

- 不能有副作用（不修改服务端状态）
- 参数放 Query String
- 对于响应变化不大的查询，可添加缓存头 `Cache-Control: max-age=60`
- 涉及敏感数据的 GET 接口不加缓存

### 3.3 POST 接口约束

- 参数统一放 JSON Body
- `Content-Type: application/json`
- 不通过 URL 传递业务参数

### 3.4 不使用的方法

- 不使用 `PUT`、`PATCH`、`DELETE` — 全部统一为 POST + 动词 URL
- 不使用 `HEAD`、`OPTIONS` — 由框架/网关自动处理

---

## 4. 请求设计

### 4.1 请求头

```http
Content-Type: application/json
ceAccessKey: ak-xxxxxxxx                # 认证 Key（HMAC-SHA256 签名）
timestamp: 1705289400000                # 请求时间戳（防重放）
sign: <HMAC-SHA256 签名>                # 请求签名
X-Trace-Id: req-abc123                  # 全链路追踪 ID
```

### 4.2 GET 请求参数

参数放 Query String，签名时需包含 Query String：

```bash
GET /api/v1/article/query_article_list?page=1&page_size=20&status=published&keyword=微服务
```

### 4.3 POST 请求参数

参数放 JSON Body：

```json
{
  "title": "Go 微服务实战",
  "content": "本文介绍...",
  "author_id": "user_001",
  "is_published": false,
  "tags": ["go", "microservice"]
}
```

### 4.4 字段命名规范

```json
// ✅ snake_case
{
  "user_id": "user_001",
  "display_name": "张三",
  "phone_number": "+8613800000000",
  "is_active": true,
  "has_children": false,
  "start_time": 1762324323000,
  "end_time": 1762410723000
}

// ❌ 避免
{
  "userId": "user_001",              // camelCase
  "displayName": "张三",             // camelCase
  "active": true,                    // 布尔字段缺少前缀
  "children": false                  // 同上
}
```

| 规则 | 正确 | 错误 |
|------|------|------|
| 通用字段 snake_case | `user_id`、`display_name` | `userId`、`displayName` |
| 布尔字段加 `is_`/`has_` 前缀 | `is_active`、`has_children` | `active`、`children` |
| 时间字段毫秒时间戳 | `"created_at": 1700000000000` | `"2024-01-01T00:00:00Z"` |

### 4.5 批量操作

```bash
POST /api/v1/article/batch_delete_article
{
  "article_ids": ["art_001", "art_002", "art_003"]
}

POST /api/v1/order/batch_create_order
{
  "orders": [
    { "product_id": "prod_001", "quantity": 2 },
    { "product_id": "prod_002", "quantity": 1 }
  ]
}
```

---

## 5. 响应设计

### 5.1 HTTP 状态码

| 状态码 | 场景 |
|--------|------|
| `200 OK` | 请求被正确处理 |
| `400 Bad Request` | 请求格式错误、参数校验失败 |
| `403 Forbidden` | 认证/鉴权失败 |
| `404 Not Found` | 接口路径不存在 |
| `429 Too Many Requests` | 触发限流 |
| `500 Internal Server Error` | 服务端系统错误 |

业务错误（资源不存在、状态冲突等）的 HTTP 状态码存在两种实践方案，详见 6.1 节。一旦选定则项目内全局统一。

### 5.2 统一响应结构

#### 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "article_id": "art_789",
    "created_at": 1705289400000
  },
  "trace_id": "req-xyz789",
  "timestamp": 1705289400124
}
```

#### 分页列表响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      { "article_id": "art_001", "title": "Go 微服务实战" },
      { "article_id": "art_002", "title": "深入理解 Go" }
    ],
    "total": 100,
    "page": 1,
    "page_size": 20
  },
  "trace_id": "req-xyz789",
  "timestamp": 1705289400124
}
```

#### 字段说明

| 字段 | 类型 | 说明 |
|------|------|------|
| `code` | int | 业务状态码，0=成功，非 0=失败（见第 6 章） |
| `message` | string | 人类可读的状态描述 |
| `data` | object/array/null | 业务数据；单个资源放对象，列表放 `{ list, total, page, page_size }` |
| `trace_id` | string | 全链路追踪 ID，每次请求唯一 |
| `timestamp` | int64 | 服务端响应时间（毫秒时间戳） |

### 5.3 规范要点

- 每个响应必须包含 `trace_id`，用于排查问题
- `data` 为 `null` 时表示无数据（非"未设置"）
- 时间字段统一用毫秒时间戳（`int64`），字段名以 `_at` 结尾
- 空列表返回 `{ "list": [], "total": 0, "page": 1, "page_size": 20 }`，不返回 `null`

---

## 6. 错误处理

### 6.1 双轨错误体系

**HTTP 状态码**（系统级） + **业务错误码 code**（业务级）。业务错误的 HTTP 状态码有两种实践方案：

#### 方案 A：业务错误返回 HTTP 200

```
业务成功/失败 → HTTP 200，通过 code 区分（0=成功，非 0=失败）
系统错误      → HTTP 400/500
```

| HTTP 状态码 | 业务错误码范围 | 含义 |
|-------------|--------------|------|
| `200` | `0` | 成功 |
| `200` | `2000–2999` | 业务错误（资源不存在、状态冲突、数据不合法） |
| `400` | `1000–1999` | 通用错误（参数校验、签名错误、限流） |
| `403` | `3000–3999` | 权限错误（AK 无权限、IP 白名单拒绝） |
| `500` | `5000–5999` | 系统错误（DB 异常、依赖服务故障） |

**优点**：调用方只需判断 `code != 0` 即可统一处理错误；监控仅对 4xx/5xx 告警，不会因正常业务失败误报。
**缺点**：与 HTTP 语义不完全一致（"资源不存在"返回 200），网关/代理层无法根据状态码区分业务失败。

#### 方案 B：业务错误返回 HTTP 400

```
业务成功      → HTTP 200
业务失败      → HTTP 400，通过 code 区分具体原因
系统错误      → HTTP 500
```

| HTTP 状态码 | 业务错误码范围 | 含义 |
|-------------|--------------|------|
| `200` | `0` | 成功 |
| `400` | `1000–1999` | 通用错误（参数校验、签名错误、限流） |
| `400` | `2000–2999` | 业务错误（资源不存在、状态冲突、数据不合法） |
| `403` | `3000–3999` | 权限错误（AK 无权限、IP 白名单拒绝） |
| `500` | `5000–5999` | 系统错误（DB 异常、依赖服务故障） |

**优点**：符合 HTTP 语义，网关/CDN 可直接利用状态码做差异化处理。
**缺点**：HTTP 400 同时承载参数错误和业务错误，监控侧需按 `code` 二次过滤（如排除 2xxx 再告警），否则"文章不存在"这类正常业务结果也会触发告警。

#### 选择建议

| 场景 | 推荐方案 |
|------|---------|
| 仅服务端渲染/内部调用，无 CDN/网关依赖 HTTP 状态码 | 方案 A（200 + code） |
| 有 CDN/网关/代理层需要根据状态码做差异化处理 | 方案 B（400 + code） |
| 团队已习惯某一种 | 延续现有习惯 |

**两种方案的共同约束**：
- 错误码范围划分保持一致（1xxx 通用、2xxx 业务、3xxx 权限、5xxx 系统）
- 错误响应格式相同（`{ code, message, data, trace_id, timestamp }`）
- 项目内全局统一，不允许混用

### 6.2 错误响应格式

```json
{
  "code": 10001,
  "message": "参数校验失败：title 不能为空",
  "data": null,
  "trace_id": "req-error-001",
  "timestamp": 1705289400000
}
```

### 6.3 错误码命名规范

按模块划分错误码区间，每个模块预留 100 个子码：

| 模块 | code 范围 | 示例 |
|------|----------|------|
| 通用（参数、签名、限流） | 1000–1099 | `10001`=参数缺失，`10002`=签名错误 |
| 文章（article） | 2000–2099 | `20001`=文章不存在，`20002`=文章已发布 |
| 用户（user） | 2100–2199 | `21001`=用户不存在，`21002`=邮箱已注册 |
| 订单（order） | 2200–2299 | `22001`=订单不存在，`22002`=订单已取消 |
| 权限 | 3000–3099 | `30001`=AK 无权限，`30002`=IP 白名单拒绝 |
| 系统 | 5000–5099 | `50001`=DB 错误，`50002`=Redis 不可用 |

### 6.4 字段级校验错误

```json
{
  "code": 10001,
  "message": "参数校验失败",
  "data": {
    "errors": [
      { "field": "title", "reason": "required", "message": "title 为必填项" },
      { "field": "email", "reason": "format_invalid", "message": "邮箱格式不正确" },
      { "field": "price", "reason": "out_of_range", "message": "价格必须在 0.01–99999.99 之间" }
    ]
  },
  "trace_id": "req-val-001",
  "timestamp": 1705289400000
}
```

字段级错误放在 `data.errors` 数组中，包含出错字段名、错误原因码和人类可读描述。

### 6.5 错误处理原则

- 错误 `message` 面向调用方，不暴露内部细节（如 SQL 语句、堆栈信息）
- 内部细节记录在服务端日志中，通过 `trace_id` 关联
- 方案 A（200 + code）：监控以 `code != 0` 为失败判定，HTTP 4xx/5xx 为系统级告警
- 方案 B（400 + code）：监控侧按 `code` 过滤，排除业务错误（2xxx）后再对 4xx 告警，避免"资源不存在"误报

---

## 7. 认证与鉴权

### 7.1 签名认证机制

采用 **AccessKey + HMAC-SHA256 签名**：

**请求头：**

```
ceAccessKey: ak-xxxxxxxx
timestamp: 1705289400000
sign: <HMAC-SHA256 签名>
```

### 7.2 签名计算规则

**POST 请求：**

```
签名内容 = accessKey + timestamp + requestPath + bodyHash
sign = HMAC-SHA256(签名内容, accessSecret)
```

**GET 请求：**

```
签名内容 = accessKey + timestamp + requestPath + queryString
sign = HMAC-SHA256(签名内容, accessSecret)
```

签名规则：
- `timestamp` 为毫秒时间戳，允许 5 分钟偏差（防重放攻击）
- `requestPath` 为完整路径（不含 Host），如 `/api/v1/article/create_article`
- `bodyHash` 为请求体的 SHA-256 哈希值（十六进制）
- `queryString` 按参数名字母序排序后拼接（`key1=val1&key2=val2`）
- 大文件上传场景：不对完整 Body 计算 hash，改为对 `metadata + 文件URL` 签名，文件本身单独校验 MD5

### 7.3 权限控制

- 每个 AccessKey 配置独立的接口权限列表、IP 白名单、调用频率配额
- 不同调用方持有不同 AccessKey，互不越权
- 权限校验失败返回 HTTP 403 + code 30001

### 7.4 免签名单

以下接口不参与签名验证、不限流：

| 接口 | 用途 |
|------|------|
| `GET /health` | K8s 存活探针 |
| `GET /ready` | K8s 就绪探针（检查 DB、Redis 等依赖） |

### 7.5 Token 认证（可选，面向浏览器客户端）

对于浏览器端调用的接口，可选用 Bearer Token 代替 AK/SK 签名：

```http
Authorization: Bearer <access_token>
```

Token 有效期短（15–60 分钟），通过 Refresh Token 续期。

---

## 8. 分页

### 8.1 基于页码的分页（标准方式）

```bash
GET /api/v1/article/query_article_list?page=1&page_size=20
```

响应：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [ ... ],
    "total": 100,
    "page": 1,
    "page_size": 20
  },
  "trace_id": "req-xyz789",
  "timestamp": 1705289400124
}
```

### 8.2 参数约定

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `page` | int | 1 | 页码，从 1 开始 |
| `page_size` | int | 20 | 每页条数，上限 100（超出自动截断并返回 `warning` 字段） |

```json
{
  "code": 0,
  "message": "success",
  "warning": "page_size 超过上限，已自动截断为 100",
  "data": { ... },
  "trace_id": "req-xyz789",
  "timestamp": 1705289400124
}
```

### 8.3 游标分页（动态列表场景）

数据变更频繁的场景（如实时动态流），支持游标分页：

```bash
GET /api/v1/feed/query_feed_list?cursor=xxx&limit=20
```

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [ ... ],
    "next_cursor": "yyy",
    "has_more": true
  },
  "trace_id": "req-xyz789",
  "timestamp": 1705289400124
}
```

`next_cursor` 为空时表示最后一批。

---

## 9. 排序与过滤

### 9.1 排序

```bash
GET /api/v1/article/query_article_list?order_by=created_at&order_dir=desc
GET /api/v1/article/query_article_list?order_by=title&order_dir=asc
```

| 参数 | 说明 |
|------|------|
| `order_by` | 排序字段名（snake_case） |
| `order_dir` | `asc`（升序）或 `desc`（降序），默认 `desc` |

多字段排序用逗号分隔：

```bash
GET /api/v1/article/query_article_list?order_by=priority,created_at&order_dir=desc,asc
```

### 9.2 过滤

```bash
# 等值过滤
GET /api/v1/article/query_article_list?status=published&author_id=user_001

# 范围过滤
GET /api/v1/article/query_article_list?start_time=1700000000000&end_time=1705000000000

# 模糊搜索
GET /api/v1/article/query_article_list?keyword=微服务

# 多选过滤
GET /api/v1/article/query_article_list?status=published,draft
```

参数同名时，用逗号分隔多个值（如 `status=published,draft`）。

### 9.3 全文搜索与过滤分离

- **过滤**：基于索引的简单条件筛选 → `query_article_list?keyword=xxx`
- **全文搜索**：走搜索引擎 → `POST /api/v1/article/search_article`

```bash
POST /api/v1/article/search_article
{
  "query": "微服务架构",
  "filters": {
    "status": "published"
  },
  "page": 1,
  "page_size": 20
}
```

---

## 10. 版本管理

### 10.1 版本策略

版本号放 URL 路径：

```bash
POST /api/v1/order/create_order
POST /api/v2/order/create_order
```

### 10.2 兼容性规则

**不需要新版本**（可在当前版本直接变更）：

- 新增接口
- 新增可选请求字段
- 新增响应字段
- 必填字段改为可选
- 放宽校验规则

**需要新版本**（不兼容变更）：

- 删除或重命名接口/字段
- 修改字段类型
- 修改 URL 结构
- 修改认证方式
- 修改错误码语义

### 10.3 版本生命周期

| 阶段 | 状态 | 迁移期 | 说明 |
|------|------|--------|------|
| **Beta** | 内测，不稳定 | — | 仅限内部调用方试用 |
| **GA** | 正式发布 | — | 生产可用，保证兼容性 |
| **Deprecated** | 废弃 | ≥ 6 个月 | 仍可使用，但建议迁移 |
| **EOL** | 完全下线 | — | 接口不再可用 |

### 10.4 废弃响应头

```
HTTP/1.1 200 OK
Deprecation: true
Sunset: Sat, 31 Dec 2026 23:59:59 GMT
Link: </api/v2/order/create_order>; rel="successor-version"
```

参考 RFC 8594（Sunset Header）。

---

## 11. 安全规范

### 11.1 基础要求

| 要求 | 说明 |
|------|------|
| **HTTPS** | 生产环境强制 HTTPS |
| **输入校验** | 服务端校验所有输入，不依赖调用方 |
| **参数化查询** | 使用 ORM / 参数化，禁止拼接 SQL |
| **CORS** | 严格配置允许来源（对外接口按需配置） |
| **请求体大小** | 普通接口 ≤ 1MB，批量 ≤ 10MB，文件上传按需评估 |
| **日志脱敏** | 不打印 AK/SK、密码、身份证号、银行卡号 |

### 11.2 签名防重放

- `timestamp` 为请求发起时的毫秒时间戳
- 服务端校验 `timestamp` 与当前时间的偏差 ≤ 5 分钟
- 超过偏差的请求直接拒绝（HTTP 400，code 10002）

### 11.3 限流

#### 限流维度

| 维度 | 粒度 | 说明 |
|------|------|------|
| 按 AccessKey | 每分钟/每日最大调用次数 | 核心控制维度 |
| 按接口 | 计算密集型接口单独限流 | 保护系统资源 |
| 按 IP | 防 DDoS | 兜底防护 |

#### 推荐算法：令牌桶

| 参数 | 说明 | 示例 |
|------|------|------|
| `capacity` | 桶容量，允许的短时峰值 | 200 |
| `rate` | 填充速率 = 稳态 QPS | 100（允许短时 2 倍突发） |

#### 限流响应

```http
HTTP/1.1 429 Too Many Requests
Retry-After: 60
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 0
X-RateLimit-Reset: 1705289460
```

```json
{
  "code": 10003,
  "message": "请求频率超限，请 60 秒后重试",
  "data": null,
  "trace_id": "req-limit-xxx",
  "timestamp": 1705289400000
}
```

#### 免限流接口

以下接口不限流、不参与签名验证：

| 端点 | 用途 |
|------|------|
| `GET /health` | K8s 存活探针 |
| `GET /ready` | K8s 就绪探针 |

### 11.4 健康检查接口

| 端点 | 用途 | 返回 |
|------|------|------|
| `GET /health` | K8s 存活探针（liveness） | `{"status": "ok"}` |
| `GET /ready` | K8s 就绪探针（readiness） | `{"status": "ready", "deps": {"db": "ok", "redis": "ok"}}` |

健康检查接口特性：不参与签名验证、不限流、不记录业务日志（避免日志刷屏）。

### 11.5 全链路追踪

每次请求记录以下信息：

- `trace_id`（请求头 `X-Trace-Id` 或响应体中）：跨服务链路唯一 ID，在整个调用链中传递不变
- `accessKey`、`apiPath`、`statusCode`、`costMillis`、`errMsg`

排查流程：通过 `trace_id` 串联所有服务节点的日志。

```json
{
  "trace_id": "req-abc123",
  "access_key": "ak-xxx",
  "api_path": "/api/v1/article/create_article",
  "status_code": 200,
  "cost_millis": 47,
  "err_msg": "",
  "timestamp": "2026-01-15T10:30:00Z"
}
```

---

## 12. 反模式清单

| 反模式 | 错误示例 | 正确做法 |
|--------|---------|---------|
| URL 只有名词（无动词） | `/api/v1/article` | `/api/v1/article/create_article` |
| 名词在前，动词在后 | `article_create` | `create_article` |
| URL 用连字符 | `/article/create-article` | `/article/create_article` |
| 所有操作都用 GET | `GET /article/create_article` | `POST /article/create_article` |
| 所有操作都用 POST | `POST /article/get_article` | `GET /article/get_article` |
| 所有错误返回 200 | `200 { code: 50001 }` | `500 { code: 50001 }` |
| GET 接口有副作用 | `GET /article/delete_article` | `POST /article/delete_article` |
| 密码/Token 放 URL | `/login?token=xxx` | Token 放 Header |
| 暴露内部错误信息 | `{ "message": "sql: table 'x' not found" }` | 返回通用错误，详情记日志 |
| 缺失 trace_id | 响应中无 trace_id 字段 | 每个响应带 trace_id |
| 分页参数不一致 | 有的用 `pageSize`，有的用 `page_size` | 统一用 `page_size` |
| 时间格式不统一 | `"2024-01-01"`和`1704067200`混用 | 统一毫秒时间戳 |
| 布尔字段缺少前缀 | `{ "active": true, "children": false }` | `{ "is_active": true, "has_children": false }` |
| 错误码无区间划分 | `100`、`200`、`201` 分散无规律 | 按模块划分范围（2xxx=文章、3xxx=权限） |
| 健康检查参与认证 | `/health` 也要签名 | 健康检查免签名、免限流 |

---

## 13. 监控与告警

### 13.1 核心指标（Google SRE 黄金四项）

| 指标 | 字段 | 说明 |
|------|------|------|
| 延迟 | avgCostMillis、P99 | 请求响应时间分布 |
| 流量 | totalCallCount、QPS | 每秒请求数 |
| 错误率 | failCallCount / totalCallCount | 失败请求占比 |
| 饱和度 | maxConcurrency vs 阈值 | 系统资源使用程度 |

### 13.2 告警分类

| 类型 | 触发条件 | 用途 |
|------|---------|------|
| 实时错误预警 | 出现 5xx 立即通知 | 捕获突发故障 |
| 统计预警 | 成功率低于阈值（周期检测） | 发现渐进式劣化 |

### 13.3 告警分级

| 级别 | 通知方式 | 适用场景 |
|------|---------|---------|
| LOW | 邮件 | 非紧急问题，如 Deprecated 版本调用量异常 |
| MEDIUM | IM（企微/钉钉/飞书） | 需要及时关注，如成功率下降 1% |
| HIGH | 电话 | 紧急故障，如 5xx 错误率突增、服务不可用 |

### 13.4 关键监控点

- **接口维度**：按 AccessKey + API Path 统计 QPS、错误率、P99 延迟
- **依赖维度**：DB 连接池、Redis 命中率、上游服务延迟
- **系统维度**：CPU、内存、goroutine 数量、GC 暂停时间
- 监控以 `code != 0`（方案 A）或 `HTTP 4xx/5xx 过滤业务错误`（方案 B）为失败判定
- 日志中记录 `trace_id`，监控平台通过 `trace_id` 串联调用链

---

## 附录A：完整请求/响应示例

### 请求示例（创建文章）

```http
POST /api/v1/article/create_article HTTP/1.1
Host: api.example.com
Content-Type: application/json
ceAccessKey: ak-123456
timestamp: 1705289400000
sign: abc123def456...

{
  "title": "API 设计最佳实践",
  "content": "本文介绍...",
  "author_id": "user_001",
  "is_published": false
}
```

### 响应示例（成功）

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "article_id": "art_789",
    "created_at": 1705289400000
  },
  "trace_id": "req-xyz789",
  "timestamp": 1705289400124
}
```

### 响应示例（参数错误）

```json
{
  "code": 10001,
  "message": "参数校验失败：title 不能为空",
  "data": null,
  "trace_id": "req-error-001",
  "timestamp": 1705289400000
}
```

---

## 参考文献

### 核心参考（直接依据）

| 文档 | 链接 | 说明 |
|------|------|------|
| RFC 9110 HTTP 语义 | https://datatracker.ietf.org/doc/html/rfc9110 | HTTP 方法语义、状态码定义 |
| RFC 8594 Sunset Header | https://datatracker.ietf.org/doc/html/rfc8594 | API 废弃的 Sunset / Deprecation 响应头 |
| HMAC-SHA256（RFC 2104） | https://datatracker.ietf.org/doc/html/rfc2104 | 签名认证的 HMAC 算法定义 |

### 参考与借鉴（理念相近，细节有差异）

| 文档 | 链接 | 说明 |
|------|------|------|
| 抖音开放平台 API 接口调用约定 | https://developer.open-douyin.com/docs/resource/zh-CN/local-life/develop/preparation/openapiinterfacecallconvention | 动作导向 URL、下划线命名、版本号在路径中 |
| 腾讯企业微信 开发前必读 | https://developer.work.weixin.qq.com/document/path/90664 | RPC 风格的动作语义 URL，但鉴权方式（access_token 在 Query String）和响应格式（errcode/errmsg）与本规范不同 |
| JSON-RPC 2.0 规范 | https://www.jsonrpc.org/specification | 同为 RPC 思想，但 JSON-RPC 将方法名放在 Body 中且使用单一端点，本规范将方法放在 URL 路径中 |
| Google AIP-136（自定义方法） | https://google.aip.dev/136 | 自定义方法的动词+名词命名思路，但 AIP 使用 `资源:动词` 冒号格式 |
| Google AIP-181（稳定性级别） | https://google.aip.dev/181 | Beta/Stable/Deprecated 生命周期模型的理念来源 |

### 延伸阅读

| 文档 | 链接 |
|------|------|
| Google SRE Book（分布式系统监控） | https://sre.google/sre-book/monitoring-distributed-systems/ |

> **说明**：本规范的 URL 格式为 `{module}/{verb}_{noun}`（RPC-over-HTTP 风格），与 JSON-RPC 2.0（方法名在 Body 中、单端点）和 Google AIP-136（冒号格式）属于不同的 RPC 实现路径，不应混淆。规范中的平台示例仅表示理念相近，各平台的具体鉴权方式、响应格式和 URL 结构各有不同。
