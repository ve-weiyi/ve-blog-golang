# HTTP REST API 设计规范

基于 Google AIP（REST 映射部分）、RFC 7231/9110（HTTP 语义）、Microsoft REST API Guidelines 和行业最佳实践整理。

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
- [附录A：业务领域端点示例](#附录a业务领域端点示例)
- [参考文献](#参考文献)

---

## 1. 设计原则

### 1.1 资源导向

API 以**资源**（Resource）为中心设计，而非以**操作**（Operation）为中心。每个资源有唯一的 URL 标识。

```
设计思路: "系统管理'用户'资源，通过 HTTP 方法对它执行 CRUD"
而非: "我要写一个查询用户接口，URL 就叫 /api/getUser"
```

### 1.2 核心约束

| 原则 | 说明 |
|------|------|
| **统一接口** | 通过 HTTP 方法（GET/POST/PUT/PATCH/DELETE）统一操作语义 |
| **无状态** | 每个请求包含完成请求所需的所有信息，服务端不保存客户端会话状态 |
| **可寻址** | 每个资源有唯一 URL，可作为标识符被引用和链接 |
| **可演进** | 向后兼容变更不应破坏已有客户端 |
| **HATEOAS**（可选） | 响应中包含相关资源的链接，提升可发现性 |

### 1.3 标准方法优先

Google AIP 定义五大标准方法，覆盖 70%+ 场景：

| 方法 | HTTP 映射 | 功能 |
|------|----------|------|
| List | GET | 分页查询资源集合 |
| Get | GET | 获取单个资源 |
| Create | POST | 创建资源 |
| Update | PUT/PATCH | 全量/部分更新资源 |
| Delete | DELETE | 删除资源 |

当标准方法无法表达时，使用自定义方法（见 3.3 节）。

---

## 2. URL 设计

### 2.1 核心规则

| 规则 | 正确 | 错误 |
|------|------|------|
| 资源名用**名词复数** | `/users`、`/articles` | `/user`、`/getUser` |
| 全小写 + 连字符 | `/blog-posts`、`/order-items` | `/blogPosts`、`/OrderItems` |
| 不使用文件扩展名 | `/users` | `/users.json`、`/users.xml` |
| 不使用动词 | `/articles` | `/getArticles`、`/createArticle` |
| 层级表达从属关系 | `/users/123/orders` | `/orders?userId=123` |
| 末尾不加斜杠 | `/users` | `/users/` |

### 2.2 资源命名

```
✅ 推荐
/users                    # 用户集合
/users/123                # 单个用户
/users/123/orders         # 用户的订单集合
/users/123/orders/456     # 用户的单个订单

❌ 避免
/getUsers                 # URL 中包含动词
/users/123/               # 末尾斜杠
/userList                 # camelCase
/usrs                     # 随意缩写
```

### 2.3 层级深度

资源嵌套不宜超过 3 层。深层嵌套可通过根路径的独立端点访问。

```
✅ 推荐（最多 3 层）
/users/123/orders/456

❌ 避免（过深）
/users/123/orders/456/items/789/comments

✅ 替代方案
/comments?orderId=456      # 通过查询参数绕过深层嵌套
```

### 2.4 父子资源嵌套原则

子资源 URI 必须包含父资源 ID：

```
GET /users/{userId}/orders          # 获取用户的订单（父子）
GET /orders/{orderId}               # 若订单可独立存在，需提供独立入口
```

关键规则：如果子资源可以脱离父资源独立存在，必须同时提供独立入口端点。

### 2.5 版本前缀

```bash
# ✅ 推荐：路径前缀
GET /v1/users
GET /v2/users

# ✅ 可接受：Header
GET /users
Accept: application/vnd.api+json;version=1

# ❌ 避免：查询参数
GET /users?version=1
```

优先使用路径前缀方式（`/v1/`），最直观且兼容性最好。

---

## 3. HTTP 方法

### 3.1 标准方法

| 方法 | 语义 | 幂等 | 安全 | 示例 |
|------|------|------|------|------|
| `GET` | 获取资源 | 是 | 是 | `GET /v1/users/123` |
| `POST` | 创建资源 | 否 | 否 | `POST /v1/users` |
| `PUT` | 全量替换 | 是 | 否 | `PUT /v1/users/123` |
| `PATCH` | 部分更新 | 否 | 否 | `PATCH /v1/users/123` |
| `DELETE` | 删除资源 | 是 | 否 | `DELETE /v1/users/123` |
| `HEAD` | 获取元数据 | 是 | 是 | `HEAD /v1/users/123` |
| `OPTIONS` | 获取支持的方法 | 是 | 是 | `OPTIONS /v1/users` |

- **幂等**：多次相同的请求产生相同的效果
- **安全**：不会修改服务端状态

### 3.2 方法使用规范

**GET** — 获取资源

```http
# 获取集合
GET /v1/users?page_size=20&page_token=xxx

# 获取单个资源
GET /v1/users/123

# 响应
200 OK
{
  "data": {
    "id": "123",
    "name": "张三",
    "email": "zhangsan@example.com"
  }
}
```

**POST** — 创建资源

```http
POST /v1/users
Content-Type: application/json

{
  "name": "张三",
  "email": "zhangsan@example.com"
}

# 响应：返回完整资源 + Location 头
201 Created
Location: /v1/users/123
{
  "data": {
    "id": "123",
    "name": "张三",
    "email": "zhangsan@example.com",
    "created_at": 1700000000000
  }
}
```

**PUT** — 全量替换（需传完整资源）

```http
PUT /v1/users/123
Content-Type: application/json

{
  "name": "张三丰",
  "email": "zhangsanfeng@example.com"
}

# 响应
200 OK
{
  "data": {
    "id": "123",
    "name": "张三丰",
    "email": "zhangsanfeng@example.com",
    "updated_at": 1700000001000
  }
}
```

**PATCH** — 部分更新（仅传需要修改的字段）

```http
PATCH /v1/users/123
Content-Type: application/json

{
  "name": "张三丰"
}

# 响应
200 OK
{
  "data": {
    "id": "123",
    "name": "张三丰",
    "email": "zhangsan@example.com",
    "updated_at": 1700000001000
  }
}
```

**DELETE** — 删除资源

```http
DELETE /v1/users/123

# 响应
204 No Content
```

### 3.3 自定义方法

当标准 CRUD 方法无法表达时，使用 `资源:动作` 格式的自定义方法。

#### 为什么用 `:` 而非 `/`

使用 `/action` 格式会产生与子资源路径的歧义：

```
❌ POST /v1/users/batch-delete
语义歧义：系统可能将 batch-delete 解析为 users 的子资源（如存在名为 batch-delete 的资源）

✅ POST /v1/users:batch-delete
语义清晰：对 users 资源执行 batch-delete 操作
```

`:` 语法显式声明这是一个**动作（动词）**，而非**资源（名词）**。这是 Google AIP-136 的核心设计规则。

#### 操作类型决策表

| 操作类型 | 示例 | URI 格式 | HTTP 方法 |
|---------|------|---------|----------|
| 自定义方法（非标准操作） | 批量删除、发布、恢复、取消 | `POST /res:action` | POST |
| 自定义方法（只读查询） | 全文搜索 | `GET /res:action` | GET |
| 标准查询（过滤/排序） | 条件过滤、分页列表 | `GET /res?filter=...` | GET |
| 子资源操作 | 获取用户的订单 | `GET /users/{id}/orders` | GET/POST 等 |

#### 自定义方法分类目录

**计算与处理类**
```
POST /v1/orders:batch-delete         # 批量删除
POST /v1/invoices/{id}:calculate-tax # 计算税费
POST /v1/sales:generate-report       # 生成报表
POST /v1/files/{id}:restore          # 恢复已删除文件
POST /v1/documents/{id}:translate    # 翻译文档
```

**状态与流程管理**
```
POST /v1/blogs/{id}:publish          # 发布文章
POST /v1/tasks/{id}:approve          # 审批任务
POST /v1/orders/{id}:cancel          # 取消订单
POST /v1/users/{id}:disable          # 停用用户
POST /v1/users/{id}:reset-password   # 重置密码
```

**批量与聚合操作**
```
POST /v1/users:import                # 批量导入
POST /v1/reports:export              # 导出数据
POST /v1/calendars/{id}:sync         # 同步数据
```

**资源生命周期扩展**
```
POST /v1/projects/{id}:undelete      # 恢复逻辑删除
POST /v1/documents/{id}:archive      # 归档资源
POST /v1/forms/{id}:validate         # 校验资源
```

**关联资源操作**
```
POST /v1/files/{id}/labels:attach    # 关联标签到文件
POST /v1/templates/{id}:clone        # 复制模板
POST /v1/files/{id}:move             # 移动文件
POST /v1/contacts:merge              # 合并联系人
```

#### 自定义方法设计要点

- 优先考虑标准方法，仅在标准语义无法满足时采用
- 动词命名使用 kebab-case，不含介词（"for"、"with" 等）
- 动词不得包含标准方法名（Get、List、Create、Update、Delete）
- 只读操作用 `GET`，有副作用或修改资源的操作用 `POST`
- 当操作可建模为资源操作时，关联到具体资源（如 `POST /v1/blogs/{id}:publish`）

---

## 4. 请求设计

### 4.1 请求头

```http
# 标准请求头
Content-Type: application/json                     # 请求体格式
Accept: application/json                           # 期望的响应格式
Authorization: Bearer <token>                      # 认证
Accept-Language: zh-CN                             # 国际化
Idempotency-Key: <uuid>                           # 幂等键（POST 场景）
X-Request-ID: <uuid>                              # 请求追踪
```

### 4.2 查询参数

| 参数 | 用途 | 示例 |
|------|------|------|
| `page_size` | 每页条数 | `page_size=20` |
| `page_token` | 分页游标 | `page_token=xxx` |
| `filter` | 筛选条件 | `filter=status=active` |
| `order_by` | 排序 | `order_by=created_at desc` |
| `fields` | 字段选择 | `fields=id,name,email` |

### 4.3 请求体

```json
// ✅ 推荐：snake_case 字段名
{
  "display_name": "张三",
  "phone_number": "+8613800000000",
  "date_of_birth": "1990-01-01"
}

// ❌ 避免：camelCase
{
  "displayName": "张三",
  "phoneNumber": "+8613800000000"
}
```

JSON 字段统一使用 `snake_case`，与数据库字段和 Protobuf 风格保持一致。

### 4.4 批量操作

```http
# 批量创建
POST /v1/users:batch-create
{
  "users": [
    { "name": "张三" },
    { "name": "李四" }
  ]
}

# 批量删除
POST /v1/users:batch-delete
{
  "ids": ["123", "456", "789"]
}

# 批量更新
POST /v1/users:batch-update
{
  "users": [
    { "id": "123", "name": "张三丰" },
    { "id": "456", "name": "李四光" }
  ]
}
```

---

## 5. 响应设计

### 5.1 HTTP 状态码

#### 成功 2xx

| 状态码 | 场景 | 说明 |
|--------|------|------|
| `200 OK` | GET、PUT、PATCH 成功 | 返回资源数据 |
| `201 Created` | POST 创建成功 | 返回新资源 + Location 头 |
| `202 Accepted` | 异步操作已受理 | 返回任务 ID 供轮询 |
| `204 No Content` | DELETE 成功 | 无响应体 |

#### 客户端错误 4xx

| 状态码 | 场景 |
|--------|------|
| `400 Bad Request` | 请求格式/参数错误 |
| `401 Unauthorized` | 未认证（缺少或无效 Token） |
| `403 Forbidden` | 已认证但无权限 |
| `404 Not Found` | 资源不存在 |
| `405 Method Not Allowed` | HTTP 方法不支持 |
| `409 Conflict` | 资源冲突（重复创建等） |
| `422 Unprocessable Entity` | 业务校验失败 |
| `429 Too Many Requests` | 触发限流 |

#### 服务端错误 5xx

| 状态码 | 场景 |
|--------|------|
| `500 Internal Server Error` | 未知内部错误 |
| `502 Bad Gateway` | 上游服务错误 |
| `503 Service Unavailable` | 服务暂不可用（维护中） |
| `504 Gateway Timeout` | 上游超时 |

### 5.2 统一响应体

#### 成功响应

```json
{
  "data": {
    "id": "123",
    "name": "张三",
    "email": "zhangsan@example.com",
    "created_at": 1700000000000
  }
}
```

#### 列表响应

```json
{
  "data": [
    { "id": "123", "name": "张三" },
    { "id": "456", "name": "李四" }
  ],
  "next_page_token": "eyJsYXN0SWQiOiA0NTZ9",
  "total_size": 100
}
```

#### 错误响应

```json
{
  "error": {
    "code": "INVALID_ARGUMENT",
    "message": "邮箱格式不正确",
    "details": [
      {
        "field": "email",
        "reason": "format_invalid",
        "message": "请输入有效的邮箱地址"
      }
    ]
  }
}
```

### 5.3 响应格式原则

- 统一使用 JSON；二进制内容用 `Content-Type: application/octet-stream`
- `data` 字段：单个资源返回对象，集合返回数组
- 时间字段统一用毫秒时间戳（`int64`），字段名 `_at` 结尾
- 空字段：`null` 而非省略；删除操作返回 `204 No Content` 且无响应体

### 5.4 字段选择

支持调用方按需获取字段，减少不必要的数据传输：

```http
GET /v1/users/123?fields=id,name,email

# 响应仅包含请求的字段
{
  "data": {
    "id": "123",
    "name": "张三",
    "email": "zhangsan@example.com"
  }
}
```

参考 Google AIP-157（Partial Responses）。

---

## 6. 错误处理

### 6.1 错误码设计

采用与 gRPC 标准状态码对齐的字符串错误码：

| 错误码 | HTTP 状态码 | 含义 |
|--------|-----------|------|
| `OK` | 200 | 成功 |
| `INVALID_ARGUMENT` | 400 | 参数错误 |
| `FAILED_PRECONDITION` | 400 | 前置条件不满足 |
| `OUT_OF_RANGE` | 400 | 参数超出范围 |
| `UNAUTHENTICATED` | 401 | 未认证 |
| `PERMISSION_DENIED` | 403 | 无权限 |
| `NOT_FOUND` | 404 | 资源不存在 |
| `ALREADY_EXISTS` | 409 | 资源已存在 |
| `RESOURCE_EXHAUSTED` | 429 | 资源耗尽（限流） |
| `CANCELLED` | 499 | 客户端取消请求 |
| `DATA_LOSS` | 500 | 数据丢失 |
| `UNKNOWN` | 500 | 未知错误 |
| `INTERNAL` | 500 | 内部错误 |
| `UNAVAILABLE` | 503 | 服务不可用 |
| `DEADLINE_EXCEEDED` | 504 | 超时 |

参考 google.rpc.Code。

### 6.2 业务校验错误

业务逻辑的校验失败放在 Response body 而非仅依赖 HTTP 状态码：

```json
// 422 Unprocessable Entity
{
  "error": {
    "code": "INVALID_ARGUMENT",
    "message": "请求参数校验失败",
    "details": [
      { "field": "email", "reason": "required", "message": "邮箱为必填项" },
      { "field": "password", "reason": "too_short", "message": "密码长度不能少于 8 位" }
    ]
  }
}
```

HTTP 状态码表达**系统级错误**（网络、认证、权限），业务校验失败通过 Response body 中的 `details` 数组表达。

---

## 7. 认证与鉴权

### 7.1 认证方式

```http
# Bearer Token（推荐）
Authorization: Bearer eyJhbGciOiJSUzI1NiIs...

# API Key（服务间调用）
X-API-Key: your-api-key
```

### 7.2 Token 实践

- Access Token 有效期短（15-60 分钟），通过 Refresh Token 续期
- Token 放在 `Authorization` 头，不放在 URL 或 Cookie 中
- Token 失效返回 `401 Unauthorized`，客户端静默刷新

### 7.3 权限模型

```http
# 资源级权限检查
GET /v1/users/123/orders
# 能访问即代表有权限，无权限返回 403

# 403 响应
{
  "error": {
    "code": "PERMISSION_DENIED",
    "message": "无权访问该资源"
  }
}
```

---

## 8. 分页

### 8.1 游标分页（推荐）

基于 `page_token` 的游标分页，应对数据变更时更稳定：

```http
# 请求
GET /v1/articles?page_size=20&page_token=eyJsYXN0SWQiOiA0NTZ9

# 响应
{
  "data": [ ... ],
  "next_page_token": "eyJsYXN0SWQiOiA0NzZ9",
  "total_size": 100      // 可选，代价较高
}
```

### 8.2 偏移分页（管理后台可接受）

```http
GET /v1/articles?page=1&page_size=20

{
  "data": [ ... ],
  "page": 1,
  "page_size": 20,
  "total": 100,
  "total_pages": 5
}
```

### 8.3 规则

- `page_size` 设置上限（如 max 100），防止客户端一次拉取过多数据
- 游标分页为首选方案；偏移分页仅在数据变更不频繁的管理后台使用
- `total_size` 可选，在数据量大的场景计算总数代价较高

---

## 9. 排序与过滤

### 9.1 排序

```http
GET /v1/articles?order_by=created_at desc
GET /v1/articles?order_by=title asc
GET /v1/articles?order_by=priority desc,created_at asc   # 多字段排序
```

格式：`order_by={field} {direction}`，多字段用逗号分隔，direction 为 `asc` 或 `desc`，默认 `asc`。

### 9.2 过滤

```http
# 简单过滤（标准 List + filter 参数）
GET /v1/articles?filter=status=published
GET /v1/articles?filter=author_id=123&filter=status=published

# 高级过滤（Google AIP-160 风格）
GET /v1/articles?filter=status="published" AND created_at>"2024-01-01"
GET /v1/articles?filter=title:"Go" AND (status="draft" OR status="published")
```

### 9.3 搜索

全文搜索与条件过滤是两种不同场景，需区分对待：

```http
# 条件过滤：走标准 List 方法
GET /v1/articles?filter=status=published&order_by=created_at desc

# 全文搜索：走自定义方法（搜索引擎）
GET /v1/articles:search?q=微服务架构&page_size=20
```

| 场景 | 方式 | URI |
|------|------|-----|
| 条件过滤/筛选 | 标准 List + filter | `GET /v1/articles?filter=status=published` |
| 全文搜索 | 自定义方法 | `GET /v1/articles:search?q=keyword` |
| 复杂搜索（需请求体） | 自定义方法 | `POST /v1/articles:search` |

`search` 属于自定义方法（AIP-136），因为其语义不同于标准 List（List 返回集合中所有匹配过滤条件的资源，Search 涉及全文索引、相关性评分、模糊匹配等搜索引擎行为）。

---

## 10. 版本管理

### 10.1 版本策略

```http
GET /v1/users      # v1
GET /v2/users      # v2（breaking change）
```

### 10.2 兼容性规则

**兼容的变更**（可在当前版本进行）：

- 新增 API 端点
- 新增可选请求字段
- 新增响应字段
- 新增 HTTP 方法（如原来只有 GET，新增 POST）
- 放宽校验规则

**不兼容的变更**（需要新版本）：

- 删除或重命名 API 端点
- 删除或重命名字段
- 修改字段类型
- 修改字段语义（如 `age` 从"岁"改为"月"）
- 收紧校验规则（原来不校验的现在校验）
- 修改认证方式

### 10.3 废弃流程

```
v1 标记废弃 → 发布 v2 → 通知调用方迁移 → 观察期（3-6 个月）→ 下线 v1
```

被废弃的 API 在响应中增设 `Deprecation: true` 头和 `Sunset: <date>` 头（RFC 8594）。

---

## 11. 安全规范

### 11.1 基础要求

| 要求 | 说明 |
|------|------|
| **HTTPS** | 生产环境强制 HTTPS，禁用 HTTP |
| **输入校验** | 所有用户输入必须在服务端校验，不可仅依赖前端校验 |
| **参数化查询** | 使用 ORM 或参数化查询，防止 SQL 注入 |
| **输出编码** | JSON 序列化自动处理 XSS，二进制上下文单独评估 |
| **CORS** | 严格配置允许的来源，不使用 `Access-Control-Allow-Origin: *` |
| **限流** | 关键端点配置 Rate Limit，返回 `429 Too Many Requests` |
| **日志脱敏** | 不在日志中打印密码、Token、身份证号等敏感信息 |

### 11.2 幂等性

非幂等的 POST 操作通过 `Idempotency-Key` 请求头保障：

```http
POST /v1/orders
Idempotency-Key: 550e8400-e29b-41d4-a716-446655440000
```

服务端以 Key 去重，重复请求返回首次结果。参考 IETF draft-idempotency-key-header。

### 11.3 并发控制

使用 ETag / If-Match 防止更新冲突（乐观锁）：

```http
# 获取资源 + ETag
GET /v1/articles/123
200 OK
ETag: "abc123"

# 更新时携带
PUT /v1/articles/123
If-Match: "abc123"

# 冲突
412 Precondition Failed
```

### 11.4 请求体大小限制

| 端点类型 | 建议上限 | 说明 |
|---------|---------|------|
| 普通 API | 1 MB | 大部分 JSON 请求 |
| 文件上传 | 根据业务 | 单独评估；使用 multipart/form-data |
| 批量操作 | 10 MB | 大量数据的批量创建/更新 |

---

## 12. 反模式清单

| 反模式 | 错误示例 | 正确做法 |
|--------|---------|---------|
| URL 中使用动词 | `POST /getUsers` | `GET /users` |
| 资源名用单数 | `/user`、`/user/123` | `/users`、`/users/123` |
| 深层嵌套 | `/a/b/c/d/e/f` | 查询参数替代或拆分端点 |
| 返回裸数组 | `[{...}, {...}]` | `{ "data": [{...}, {...}] }` |
| 所有错误返回 200 | `200 OK { error: "..." }` | 使用恰当的状态码 |
| 所有错误返回 500 | `500 { error: "..." }` | 区分 4xx 和 5xx |
| 在请求体中传 Token | `{ "token": "..." }` | 使用 `Authorization` 头 |
| 暴露内部细节 | `{ "sql_error": "Table 'x' doesn't exist" }` | 返回通用错误信息，详情记日志 |
| 分页无上限 | `page_size=10000` | 设置 max limit（如 100） |
| 时间格式不统一 | `"2024-01-01"` / `"2024/01/01"` / `1704067200` | 统一用毫秒时间戳（int64） |
| 布尔字段前缀不一致 | `is_active` / `active` / `has_permission` | 统一使用 `is_`、`has_` 前缀 |
| 用 `/verb` 代替 `:verb` | `POST /users/batch-delete` | `POST /users:batch-delete`（`/` 表示子资源，`:` 表示动作） |

---

## 附录A：业务领域端点示例

以下示例覆盖常见业务场景，遵循本规范的所有约定。

### 用户管理

```
GET    /v1/users                        # 分页查询用户列表
POST   /v1/users                        # 创建新用户
GET    /v1/users/{userId}               # 获取用户详情
PUT    /v1/users/{userId}               # 全量更新用户信息
PATCH  /v1/users/{userId}               # 部分更新用户
DELETE /v1/users/{userId}               # 删除用户
GET    /v1/users/{userId}/permissions   # 获取用户权限
POST   /v1/users/{userId}/avatar        # 上传用户头像
DELETE /v1/users/{userId}/sessions      # 终止用户所有活跃会话
GET    /v1/users/{userId}/roles/{roleId}# 查询用户绑定的特定角色
POST   /v1/users:batch-create           # 批量创建用户
POST   /v1/users:batch-delete           # 批量删除用户
POST   /v1/users:batch-update           # 批量更新用户
POST   /v1/users:import                 # 批量导入用户
POST   /v1/users/{userId}:disable       # 停用用户
POST   /v1/users/{userId}:reset-password# 重置密码
```

### 产品目录

```
GET    /v1/products                              # 过滤产品
POST   /v1/products                              # 创建新产品
GET    /v1/products/{productId}                  # 查询产品详情
PUT    /v1/products/{productId}                  # 全量更新产品
PATCH  /v1/products/{productId}/stock            # 更新库存
DELETE /v1/products/{productId}                  # 下架产品
GET    /v1/products/{productId}/reviews          # 获取产品评价
POST   /v1/products/{productId}/reviews          # 添加评价
DELETE /v1/products/{productId}/reviews/{id}     # 删除评价
GET    /v1/products/{productId}/versions         # 查询历史版本
POST   /v1/products/{productId}/attachments      # 上传附件
GET    /v1/categories/{categoryId}/products      # 按分类查询产品
PATCH  /v1/products/{productId}/pricing          # 更新产品价格
GET    /v1/brands/{brandId}/products             # 按品牌查询产品
POST   /v1/products:batch-delete                 # 批量删除产品
POST   /v1/products:batch-create                 # 批量创建产品
```

### 订单系统

```
POST   /v1/users/{userId}/orders               # 为用户创建订单
GET    /v1/users/{userId}/orders/{orderId}     # 查询用户订单
DELETE /v1/users/{userId}/orders/{orderId}     # 取消订单
GET    /v1/orders?filter=status=shipped        # 全局过滤订单
PATCH  /v1/orders/{orderId}/status             # 更新订单状态
POST   /v1/orders/{orderId}/items              # 添加商品项
DELETE /v1/orders/{orderId}/items/{itemId}     # 移除商品项
GET    /v1/orders/{orderId}/invoice            # 下载发票
POST   /v1/orders/{orderId}/refunds            # 发起退款
PUT    /v1/orders/{orderId}/shipping-address   # 更新配送地址
GET    /v1/users/{userId}/orders/history       # 查询历史订单
PATCH  /v1/orders/{orderId}/priority           # 调整订单优先级
GET    /v1/warehouses/{warehouseId}/orders     # 查询仓库关联订单
POST   /v1/orders/{orderId}/notifications      # 发送订单通知
DELETE /v1/orders/{orderId}/attachments/{id}   # 删除订单附件
GET    /v1/orders/{orderId}/timeline           # 获取订单状态时间线
PUT    /v1/orders/{orderId}/coupon             # 应用优惠券
POST   /v1/orders/{orderId}:split              # 拆分订单
GET    /v1/orders/stats?period=monthly         # 获取订单统计
POST   /v1/orders:batch-create                 # 批量创建订单
POST   /v1/orders:batch-delete                 # 批量删除订单
POST   /v1/orders/{orderId}:cancel             # 取消订单（自定义方法）
```

### 博客系统

```
POST   /v1/blogs                              # 创建文章
GET    /v1/blogs/{blogId}                     # 获取文章
PUT    /v1/blogs/{blogId}                     # 全量更新文章
DELETE /v1/blogs/{blogId}                     # 删除文章
GET    /v1/blogs/{blogId}/comments            # 获取评论
POST   /v1/blogs/{blogId}/comments            # 新增评论
DELETE /v1/blogs/{blogId}/comments/{id}       # 删除评论
PATCH  /v1/blogs/{blogId}/likes               # 更新点赞数
GET    /v1/users/{userId}/blogs               # 查询用户所有文章
PUT    /v1/blogs/{blogId}/content             # 全量更新文章内容
POST   /v1/blogs/{blogId}/tags                # 添加标签
DELETE /v1/blogs/{blogId}/tags/{tagId}        # 移除标签
GET    /v1/tags/{tagName}/blogs               # 按标签查询文章
POST   /v1/blogs/{blogId}:publish             # 发布文章
GET    /v1/blogs/{blogId}/versions/{id}       # 获取历史版本
POST   /v1/blogs/{blogId}/attachments         # 上传附件
DELETE /v1/blogs/{blogId}/subscriptions       # 取消订阅
GET    /v1/blogs/trending?top=10              # 获取热门文章
GET    /v1/blogs:search?q=微服务架构           # 全文搜索
POST   /v1/blogs:batch-delete                 # 批量删除文章
```

### 文件存储

```
POST   /v1/files                              # 上传文件
GET    /v1/files/{fileId}                     # 下载文件
DELETE /v1/files/{fileId}                     # 删除文件
GET    /v1/folders/{folderId}/files           # 列出文件夹内文件
POST   /v1/folders/{folderId}/files           # 向文件夹上传文件
PATCH  /v1/files/{fileId}/permissions         # 修改文件权限
PUT    /v1/files/{fileId}/metadata            # 更新元数据
GET    /v1/files/{fileId}/history             # 获取版本历史
POST   /v1/files/{fileId}:restore             # 恢复到指定版本
DELETE /v1/folders/{folderId}                 # 删除文件夹
PATCH  /v1/shared-links/{linkId}              # 更新分享链接
GET    /v1/users/{userId}/trash/files         # 查询回收站文件
POST   /v1/files/{fileId}:move                # 移动文件
GET    /v1/files:search?name=report.docx      # 全局文件搜索
POST   /v1/files:zip                          # 批量打包下载
```

### 组织架构

```
POST   /v1/companies                                  # 创建公司
GET    /v1/companies/{companyId}/departments          # 查询部门列表
POST   /v1/companies/{companyId}/departments          # 新增部门
GET    /v1/departments/{deptId}/employees             # 获取部门员工
POST   /v1/departments/{deptId}/employees             # 向部门添加员工
DELETE /v1/departments/{deptId}/employees/{empId}     # 从部门移除员工
PUT    /v1/employees/{empId}/manager                  # 更新直属经理
GET    /v1/companies/{companyId}/projects             # 查询公司项目
POST   /v1/projects/{projectId}/teams                 # 为项目分配团队
DELETE /v1/companies/{companyId}/departments/{deptId} # 删除部门
PATCH  /v1/employees/{empId}/status                   # 更新员工状态
GET    /v1/teams/{teamId}/members                     # 获取团队成员
POST   /v1/teams/{teamId}/tasks                       # 创建任务
GET    /v1/employees/{empId}/reports                  # 获取下属
PUT    /v1/projects/{projectId}/budget                # 更新项目预算
POST   /v1/companies/{companyId}/locations            # 添加办公地点
DELETE /v1/companies/{companyId}/locations/{locId}    # 移除办公地点
GET    /v1/employees/{empId}/attendance?month=2025-03 # 查询考勤
POST   /v1/employees:batch-update                     # 批量更新员工
PATCH  /v1/tasks/{taskId}/progress                    # 更新任务进度
GET    /v1/companies/{companyId}/analytics/salaries   # 获取薪资分析
POST   /v1/teams/{teamId}/meetings                    # 安排团队会议
DELETE /v1/tasks/{taskId}/assignments/{empId}         # 解除任务分配
PUT    /v1/companies/{companyId}/policies/{policyId}  # 更新公司政策
GET    /v1/companies/{companyId}/hierarchy            # 获取组织架构树
POST   /v1/companies:batch-create                     # 批量创建公司
```

---

## 参考文献

| 文档 | 链接 |
|------|------|
| Google AIP 总览 | https://google.aip.dev/ |
| AIP-121 资源导向设计 | https://google.aip.dev/121 |
| AIP-122 资源命名 | https://google.aip.dev/122 |
| AIP-131 标准方法（Get） | https://google.aip.dev/131 |
| AIP-132 标准方法（List） | https://google.aip.dev/132 |
| AIP-133 标准方法（Create） | https://google.aip.dev/133 |
| AIP-134 标准方法（Update） | https://google.aip.dev/134 |
| AIP-135 标准方法（Delete） | https://google.aip.dev/135 |
| AIP-136 自定义方法 | https://google.aip.dev/136 |
| AIP-157 部分响应 | https://google.aip.dev/157 |
| AIP-158 分页 | https://google.aip.dev/158 |
| AIP-160 过滤 | https://google.aip.dev/160 |
| AIP-181 版本管理 | https://google.aip.dev/181 |
| RFC 7231 HTTP/1.1 语义 | https://datatracker.ietf.org/doc/html/rfc7231 |
| RFC 9110 HTTP 语义 | https://datatracker.ietf.org/doc/html/rfc9110 |
| RFC 8594 Sunset Header | https://datatracker.ietf.org/doc/html/rfc8594 |
| Microsoft REST API Guidelines | https://github.com/microsoft/api-guidelines |
| Zalando REST API Guidelines | https://opensource.zalando.com/restful-api-guidelines/ |
| JSON API 规范 | https://jsonapi.org/ |
| google.rpc.Code | https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto |
