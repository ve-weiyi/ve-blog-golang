# RPC 设计规范 (gRPC + Protobuf)

基于 Google AIP、gRPC 官方风格指南和行业最佳实践整理，作为团队 gRPC 服务设计与评审的参考标准。

---

## 目录

- [1. 资源导向设计](#1-资源导向设计)
- [2. Proto 文件组织](#2-proto-文件组织)
- [3. 包与选项命名](#3-包与选项命名)
- [4. 标准方法](#4-标准方法)
- [5. 自定义方法](#5-自定义方法)
- [6. 消息命名](#6-消息命名)
- [7. 字段命名与类型](#7-字段命名与类型)
- [8. 请求与响应设计](#8-请求与响应设计)
- [9. 枚举与错误处理](#9-枚举与错误处理)
- [10. 版本管理](#10-版本管理)
- [11. 反模式清单](#11-反模式清单)
- [12. 合规检查清单](#12-合规检查清单)
- [参考文献](#参考文献)

---

## 1. 资源导向设计

> 参考：AIP-121（Resource-Oriented Design）

### 1.1 核心原则

gRPC API 设计以**资源**（Resource）为中心，而非以**动作**（Action）为中心。设计时先确定资源是什么，再定义对资源的操作。

```
正确思路: "这个服务管理'书籍'资源，需要对它做 CRUD"
错误思路: "我要写一个查询书籍的接口，名字就叫 QueryBooks"
```

### 1.2 资源命名

| 规则 | 正确 | 错误 |
|------|------|------|
| 名词而非动词 | `Book`、`Order` | `GetBook`、`CreateOrder` |
| 单数形式（资源名本身） | `Book`、`User` | `Books`、`Users` |
| 大驼峰（PascalCase） | `BookSummary` | `book_summary`、`bookSummary` |
| 避免随意缩写 | `Configuration`（公认的 `Config`、`Id` 除外） | `Cfg`、`Mgr`、`Attr` |

### 1.3 服务划分

一个 Service 管理一组紧密相关的资源，按**业务领域**而不是按**数据表**划分：

```
✅ 推荐：按领域聚合
LibraryService  → Book, Author, Publisher        // 都属于"图书领域"

✅ 推荐：单一资源独立成服务
UserService  → User                               // 用户是独立领域

❌ 避免：跨领域混合
MiscService → Book, Notification, Payment         // 三者毫无关联
```

---

## 2. Proto 文件组织

> 参考：AIP-122（Resource Names）、Protobuf Style Guide

### 2.1 文件命名

| 规则 | 正确 | 错误 |
|------|------|------|
| 全小写 + 下划线 | `book_service.proto`、`order.proto` | `BookService.proto`、`order-v1.proto` |
| 不含版本号 | 版本通过 package 或目录隔离 | `book_v1.proto` |
| 一个领域一个文件 | `library.proto`（含 Book + Author + Publisher） | 每张表一个文件 |

### 2.2 文件内部布局顺序

```protobuf
syntax = "proto3";                    // 1. syntax 声明
package xxx.v1;                       // 2. 包名
option go_package = "...";            // 3. Go 包路径

// ==============================
// 共享消息（分页、通用类型等）
// ==============================

// ==============================
// 核心数据模型（资源定义）
// ==============================

// ==============================
// 请求/响应消息（按资源分组）
// ==============================

// ==============================
// 服务定义
// ==============================
```

### 2.3 服务内 RPC 排序

service 内的 RPC 方法按以下顺序排列，通过代码顺序体现，无需额外注释：

```
标准方法（CRUD）→ 批量方法 → 自定义方法
```

具体顺序：

1. `List{Resources}` — 分页列表
2. `Get{Resource}` — 单资源查询
3. `Create{Resource}` — 创建
4. `Update{Resource}` — 更新
5. `Delete{Resource}` — 删除
6. `Batch{Action}{Resources}` — 批量操作
7. 特殊业务方法 — 按"核心业务 → 辅助查询"排列

---

## 3. 包与选项命名

> 参考：Protobuf Style Guide

### 3.1 proto package

包名使用有意义的小写名称，通常包含版本信息：

```protobuf
// ✅ 推荐
package library.v1;
package identity.auth.v1;

// ❌ 避免
package mypackage;
package service;           // 太泛，容易冲突
```

### 3.2 go_package

Go 生成路径应与 proto package 对应：

```protobuf
option go_package = "github.com/example/project/api/gen/go/library/v1;libraryv1";
```

格式：`{full_module_path};{alias}`，别名用 camelCase。

### 3.3 Go 生成选项

`go_package` 之外，protoc 支持通过 `--go_opt` 和 `--go-grpc_opt` 控制生成行为：

```bash
# 生成 pb.go 和 _grpc.pb.go 到指定目录
protoc --go_out=gen/go --go_opt=paths=source_relative \
       --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
       library.proto
```

`paths=source_relative` 按 proto 文件相对路径输出，避免生成与模块路径绑定的深层目录结构。

---

## 4. 标准方法

> 参考：AIP-131（Standard Methods）

### 4.1 五个标准方法

| 方法 | HTTP 映射 | RPC 命名格式 | 语义 |
|------|----------|-------------|------|
| **List** | GET | `List{Resources}` | 分页查询集合，`Resources` 用**复数** |
| **Get** | GET | `Get{Resource}` | 按主键获取单个资源 |
| **Create** | POST | `Create{Resource}` | 创建新资源 |
| **Update** | PATCH/PUT | `Update{Resource}` | 更新已有资源 |
| **Delete** | DELETE | `Delete{Resource}` | 删除（支持批量） |

### 4.2 List — 集合名称必须用复数

> AIP-132: "List methods **must** use the plural form of the resource name."

```protobuf
// ✅ 正确
rpc ListBooks(ListBooksRequest) returns (ListBooksResponse);

// ❌ 错误
rpc ListBook(ListBookRequest) returns (ListBookResponse);
rpc QueryBooks(QueryBooksRequest) returns (QueryBooksResponse);  // 非标准动词
```

**为什么必须用复数**：

1. **语义**：`ListBooks` 表达"列出书籍集合"；`ListBook` 暗示只返回一本书
2. **可读性**：`client.ListBooks(ctx, req)` 自然；`client.ListBook(ctx, req)` 产生歧义
3. **工具链**：grpc-gateway 自动映射 `ListBooks` → `GET /books`、`ListBook` → `GET /book`（不符合 REST 惯例）
4. **行业惯例**：Google Cloud API、Kubernetes API、etcd API 等全部使用复数 List

### 4.3 Get — 按主键获取

```protobuf
// ✅ 正确
rpc GetBook(GetBookRequest) returns (GetBookResponse);

message GetBookRequest {
  string name = 1;  // 资源主键（如 "books/123"）
}
```

Get 的参数是资源的**主键标识符**，不接受筛选条件。如果需要按条件查询，用 List + filter。

### 4.4 Create — 返回资源标识

```protobuf
// ✅ 正确
rpc CreateBook(CreateBookRequest) returns (Book);

message CreateBookRequest {
  Book book = 1;
}

// ✅ 或返回创建后的完整资源
rpc CreateBook(CreateBookRequest) returns (CreateBookResponse);

message CreateBookResponse {
  Book book = 1;
}
```

AIP-133 推荐返回完整资源对象而非仅 ID，让调用方无需立即再发 Get 请求。

### 4.5 Update — 标准更新与字段级更新

```protobuf
// 标准全量更新
rpc UpdateBook(UpdateBookRequest) returns (Book);

message UpdateBookRequest {
  Book book = 1;
  google.protobuf.FieldMask update_mask = 2;  // 指定要更新的字段
}
```

使用 `FieldMask` 实现部分更新是行业标准做法，比拆分多个 `Update{Resource}{Field}` 方法更具扩展性。

### 4.6 Delete — 单个资源删除

```protobuf
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty);

message DeleteBookRequest {
  string name = 1;  // 单个资源标识
}
```

AIP-135 的 Delete 针对单个资源。批量删除应定义为自定义方法（见 5.2 节）。

---

## 5. 自定义方法

> 参考：AIP-136（Custom Methods）

### 5.1 命名格式

```
Verb + Noun
```

**动词在前，名词在后**。动词表达动作或状态变更，名词是被操作的资源。

```protobuf
// ✅ Verb + Noun
rpc TranslateText(TranslateTextRequest) returns (TranslateTextResponse);
rpc PublishBook(PublishBookRequest) returns (PublishBookResponse);
rpc BatchDeleteBooks(BatchDeleteBooksRequest) returns (BatchDeleteBooksResponse);

// ❌ Noun + Verb（反了）
rpc TextTranslate(TextTranslateRequest) returns (TextTranslateResponse);
rpc BookPublish(BookPublishRequest) returns (BookPublishResponse);
```

### 5.2 自定义方法分类

| 类别 | 格式 | 示例 |
|------|------|------|
| 执行动作 | `{Action}{Resource}` | `TranslateText`、`SendEmail`、`PublishBook` |
| 状态转换 | `{Action}{Resource}` | `CancelOrder`、`ApproveRequest`、`ArchiveDocument` |
| 批量操作 | `Batch{Action}{Resources}` | `BatchDeleteBooks`、`BatchUpdateOrders` |
| 关联操作 | `{Action}{Resource1}{Resource2}` | `AddBookToShelf`、`RemoveUserFromGroup` |
| 认证授权 | `LoginByPassword`、`Logout`、`SignUp` | 认证方法同样遵循 Verb+Noun |

### 5.3 常用动词清单

| 动词 | 含义 | 示例 |
|------|------|------|
| `Get` | 获取（标准） | `GetUser` |
| `List` | 列表（标准） | `ListUsers` |
| `Create` | 创建（标准） | `CreateOrder` |
| `Update` | 更新（标准） | `UpdateProfile` |
| `Delete` | 删除（标准） | `DeleteComment` |
| `Batch{Verb}` | 批量操作 | `BatchDeleteMessages` |
| `Search` | 搜索（不同于 List 的简单筛选） | `SearchProducts` |
| `Send` | 发送 | `SendNotification` |
| `Verify` | 验证 | `VerifyEmail` |
| `Publish` | 发布/上线 | `PublishArticle` |
| `Revoke` | 撤回/下线 | `RevokeToken` |
| `Archive` | 归档 | `ArchiveDocument` |
| `Export` | 导出 | `ExportReport` |
| `Import` | 导入 | `ImportData` |
| `Translate` | 翻译/转换 | `TranslateText` |
| `Apply` | 应用 | `ApplyDiscount` |
| `Cancel` | 取消 | `CancelSubscription` |
| `Approve` | 批准 | `ApproveRequest` |
| `Reject` | 拒绝 | `RejectApplication` |
| `Assign` | 分配 | `AssignRole` |
| `Unassign` | 取消分配 | `UnassignRole` |

### 5.4 避免使用的动词

| 动词 | 原因 | 替代方案 |
|------|------|---------|
| `Add`、`Remove` | 不应用于 CRUD 场景 | 创建/删除资源本身用 `Create`/`Delete`；关联操作用 `Add`/`Remove`（如 `AddBookToShelf`） |
| `Edit`、`Modify` | 不是标准动词 | `Update` |
| `Fetch`、`Query`、`Find` | 不是标准动词 | `Get`、`List`、`Search` |
| `Do`、`Execute`、`Process`、`Handle` | 含义太笼统 | 用具体的业务动词 |
| `Set`、`Toggle` | 语义不够精确 | `Update` + FieldMask 或 `Enable`/`Disable` |

### 5.5 认证方法命名

认证方法（登录、注册、登出等）同样遵循 AIP-136 的 **VerbNoun** 格式。Google Identity Platform（Google 官方认证服务）是行业权威参考：

```protobuf
// ✅ Google Identity Platform 风格（VerbNoun）
rpc SignInWithEmailPassword(SignInWithEmailPasswordRequest) returns (SignInResponse);
rpc SignInWithPhoneNumber(SignInWithPhoneNumberRequest) returns (SignInResponse);
rpc SignUp(SignUpRequest) returns (SignUpResponse);

// ✅ 保留 "Login" 术语时，同样遵循 VerbNoun
rpc LoginByPassword(LoginByPasswordRequest) returns (LoginResponse);
rpc LoginByEmail(LoginByEmailRequest) returns (LoginResponse);

// ❌ NounVerb（违反 AIP-136）
rpc PasswordLogin(PasswordLoginRequest) returns (LoginResponse);
rpc EmailCodeLogin(EmailCodeLoginRequest) returns (LoginResponse);
```

**依据**：
- Google Identity Platform 使用 `signInWithEmailPassword`、`signInWithIdp`、`signUp`（均为 Verb(+Prep)+Noun 格式）
- 认证动作虽然没有显式的"资源"，但 AIP-136 的 VerbNoun 格式仍然适用：`Login`/`SignIn` 是动词，`ByPassword`/`ByEmailCode` 是对动作的限定
- 无论选择 `Login` 还是 `SignIn` 作为动词，**动词必须在名词之前**

---

## 6. 消息命名

> 参考：AIP-131、AIP-132

### 6.1 Request / Response 命名

每个 RPC 方法必须有**专属的** Request 和 Response 消息：

```
{RpcMethod}Request
{RpcMethod}Response
```

**禁止跨 RPC 复用 Request/Response 消息**：

```protobuf
// ✅ 正确
rpc GetBook(GetBookRequest) returns (GetBookResponse);
rpc ListBooks(ListBooksRequest) returns (ListBooksResponse);
rpc CreateBook(CreateBookRequest) returns (CreateBookResponse);

// ❌ 错误
rpc GetBook(Request) returns (Response);                     // 不带语义
rpc ListBooks(ListBooksRequest) returns (ListResponse);      // 多个 RPC 共用
```

禁止复用的原因：

1. **独立演化**：向 `CreateBookRequest` 新增字段不应影响 `UpdateBookRequest`
2. **文档质量**：代码生成工具依赖专属类型名来生成精确的 API 参考
3. **向后兼容**：废弃字段时，影响面精确到单个 RPC，不会波及无关方法

### 6.2 资源消息命名

```protobuf
// 核心模型：名词
message Book { ... }
message Author { ... }

// 派生视图：名词 + 限定词
message BookSummary { ... }         // 摘要/列表项
message BookDetail { ... }          // 详情（含关联数据）
message BookStatistics { ... }      // 统计聚合
```

### 6.3 空消息

```protobuf
// 即使当前无参数，也要定义独立空消息
rpc ListAllBooks(ListAllBooksRequest) returns (ListAllBooksResponse);

message ListAllBooksRequest {}

message ListAllBooksResponse {
  repeated Book books = 1;
}
```

---

## 7. 字段命名与类型

> 参考：Protobuf Style Guide、AIP-142

### 7.1 命名风格

```protobuf
// ✅ snake_case
message Book {
  string display_name = 1;
  int64 page_count = 2;
  string isbn = 3;             // 首字母缩写词全小写
  int64 created_at = 4;
}

// ❌ 错误
message Book {
  string displayName = 1;      // camelCase
  int64 PageCount = 2;         // PascalCase
  string ISBN = 3;             // 首字母缩写词全大写（isbn 才是标准）
}
```

### 7.2 首字母缩写词

按照 Protobuf 惯例，字段名中的首字母缩写词全小写：

```protobuf
string user_id = 1;       // ✅
string user_uid = 2;      // ✅
string user_id_url = 3;   // ✅ 两个缩写词连用也全小写
string user_ID = 4;       // ❌
string userId = 5;        // ❌ camelCase
```

消息名和 RPC 方法名中的首字母缩写词保持首字母大写，其余小写：

```protobuf
message RpcStatus { ... }          // ✅
message GetRpcStatusRequest { ... } // ✅
message GetRPCStatusRequest { ... } // ❌

message OauthToken { ... }         // ❌ OAuth 作为一个词时，保留大写：OAuthToken
message OAuthToken { ... }         // ✅
```

一致性问题：公认的多字母缩写词（如 `OAuth`）在消息名中保留原有大写形式。

### 7.3 时间类型

```protobuf
// 方案 A（推荐）：int64 毫秒时间戳
int64 created_at = 1;   // 毫秒时间戳
int64 updated_at = 2;

// 方案 B：google.protobuf.Timestamp（纳秒精度，序列化体积更大）
import "google/protobuf/timestamp.proto";
google.protobuf.Timestamp created_at = 1;
```

**推荐方案 A**：`int64` 毫秒时间戳与主流数据库的时间存储一致，零序列化开销，跨语言无歧义。仅在需要纳秒精度或与 Google Cloud 服务集成时使用方案 B。

**字段名约定**：时间字段以 `_at` 结尾（`created_at`、`updated_at`、`deleted_at`、`published_at`）。

### 7.4 List Response 中 repeated 字段的命名

> AIP-132: 使用资源名的复数形式。

```protobuf
message ListBooksResponse {
  repeated Book books = 1;              // ✅ 复数，与类型名 Book 对应
  string next_page_token = 2;
  int32 total_size = 3;
}

// ❌ 错误
message ListBooksResponse {
  repeated Book results = 1;            // 应命名为 books
  repeated Book items = 1;              // 同上
}
```

字段名应与资源类型名保持可推导关系：`repeated Book` → `books`，`repeated OrderItem` → `order_items`。

### 7.5 分页

**方案 A（推荐）：cursor 分页**

> 参考：AIP-158（Pagination）

```protobuf
message ListBooksRequest {
  int32 page_size = 1;                   // 每页数量
  string page_token = 2;                // 分页游标（首次请求为空）
}

message ListBooksResponse {
  repeated Book books = 1;
  string next_page_token = 2;           // 下一页游标，为空表示最后一页
  int32 total_size = 3;                // 总数（可选，COUNT 代价较高）
}
```

cursor 分页将上一页最后一条记录的排序值编码进 `page_token`，服务端据此定位下一页起点。面对数据实时变更（插入/删除）不会出现重复或遗漏记录，且大偏移量下性能无衰减——这是 Google Cloud API、Kubernetes API 等头部项目的标准做法。

`total_size` 代价较高，仅在管理后台等对总数有强需求的场景加入。

**方案 B（备选）：offset 分页**

```protobuf
message ListBooksRequest {
  int64 page = 1;                        // 页码（从 1 开始）
  int64 page_size = 2;                  // 每页数量
}

message ListBooksResponse {
  repeated Book books = 1;
  int64 total = 2;                      // 总记录数
}
```

offset 分页直接使用 `LIMIT/OFFSET`，实现简单，天然支持跳页和总页数展示。缺点是大偏移量下数据库扫描行数增多导致性能衰减，数据实时变更时可能出现记录重复或遗漏。

**选择建议**：C 端无限滚动 / 消息列表等对数据一致性敏感的流式场景选择方案 A；后台管理表格、需要分页跳转和总数展示的场景选择方案 B。一旦选定则全局统一。

### 7.6 可选字段

使用 `optional` 关键字标记可区分"未传"和"传了零值"的字段：

```protobuf
message ListBooksRequest {
  string parent = 1;                    // 必填
  optional string filter = 2;          // 可选筛选
  optional Status status = 3;          // 可选，可区分"未传"和"传了 0"
}
```

### 7.7 注释规范

注释保持极简，避免注释腐化：

| 类型 | 要求 | 示例 |
|------|------|------|
| 模块（文件头） | 仅标注模块名称 | `// 用户服务` |
| 字段 | 说明含义 + 枚举值（如有） | `int64 status = 1;  // 状态: 0=未知 1=正常 2=禁用` |
| RPC 方法 | 仅写核心业务说明 | `// 按邮箱搜索用户` |
| 消息类型 | 不写注释（命名自解释） | — |

禁止的注释：
- 设计规则说明（如"该字段使用 int64 是因为..."）
- 权限用途标注（如"仅管理员可调用"）
- 冗余分隔线（如 `// ===== 用户相关 =====`）
- 状态提示（如"必填"、"可选"——用 `optional` 关键字表达）

---

## 8. 请求与响应设计

### 8.1 Request 设计

| 原则 | 说明 |
|------|------|
| 分页参数在前 | `page_size` 和 `page_token` 放在靠前位置 |
| 资源标识在路径中 | AIP-122 使用 `string name` 表示 `{collection}/{id}` |
| 筛选条件用 `optional` | 方便调用方只传需要的条件 |
| 每个 RPC 独立 Request | 不跨 RPC 复用 |
| 必填字段不标记 optional | proto3 默认即为必填 |

### 8.2 Response 设计

| 原则 | 说明 |
|------|------|
| Create 返回完整资源 | AIP-133 推荐返回创建后的完整对象 |
| Update 返回完整资源 | 同理，返回更新后的完整对象 |
| Delete 返回 Empty | AIP-135 标准方式 |
| List 返回资源集合 | 包含分页游标 |
| 禁止空壳 Response | 如需明确的操作结果，用 `bool success` 或自定义状态 |
| 长耗时操作用 LRO | AIP-151：长时间运行操作返回 Operation 对象 |

### 8.3 资源标识（Resource Name）

AIP-122 推荐使用 `string name` 而非 `int64 id` 作为资源主键：

```protobuf
message Book {
  string name = 1;          // "publishers/123/books/456"
  string display_name = 2;  // 人类可读的显示名
}

message GetBookRequest {
  string name = 1;          // 完整的资源标识
}
```

优点：支持层级关系、便于权限校验、与 Google Cloud API 风格一致。
不过 `int64 id` 在内部微服务中代价更低且更简单，两种方案根据场景选择，一旦选定则全局统一。

---

## 9. 枚举与错误处理

### 9.1 状态字段方案对比

两种方案各有适用场景，根据项目实际情况选择，一旦选定则全局统一。

**方案 A：int32 + 注释**

```protobuf
// 0=未知 1=正常 2=禁用 3=已删除
int32 status = 1;
```

优点：与数据库存储天然一致（存整数），JSON 序列化为数字，无零值语义问题。
缺点：缺少编译期类型检查，调用方需要查阅注释才知道有效值范围。

**方案 B：proto3 enum**

```protobuf
enum Status {
  STATUS_UNSPECIFIED = 0;   // proto3 强制要求 0 值存在
  ACTIVE = 1;
  DISABLED = 2;
}
```

优点：编译期类型安全，IDE 自动补全，API 文档自动生成枚举值说明，Google Cloud API 和 Kubernetes API 等头部项目广泛使用。
缺点：零值必须存在且不能作为有效业务值；JSON 序列化为字符串（与数据库 int 存储不一致）；旧客户端可能收到未知枚举值。

**enum 使用规范**（选择方案 B 时遵循）：

- 0 值命名为 `{ENUM_NAME}_UNSPECIFIED`，仅作默认占位，不作为业务有效值
- 结合 `optional Status` 使用，区分"未设置"与"UNSPECIFIED"
- 服务端对 UNSPECIFIED 做防御性校验，返回明确的参数错误

### 9.2 业务校验错误放在 Response 中

```protobuf
message CreateBookResponse {
  Book book = 1;
  repeated FieldError errors = 2;  // 业务校验错误
}

message FieldError {
  string field = 1;      // 出错字段名
  string message = 2;    // 错误描述
}
```

gRPC status code 用于表达**系统级错误**（网络、权限、内部错误）。**业务校验失败**（如邮箱格式不对、库存不足）放在 Response body 中，便于调用方区分处理逻辑。

---

## 10. 版本管理

> 参考：AIP-181（Versioning）

### 10.1 兼容的变更（可安全进行）

- 新增 RPC 方法
- 新增 Message 字段（分配新编号，不重用废弃编号）
- 新增 `optional` 字段
- 废弃字段（使用 `[deprecated = true]`，保留编号）

### 10.2 不兼容的变更（需要新版本）

- 重命名 RPC 方法或 Message
- 重命名字段
- 修改字段类型
- 修改或删除字段编号
- 删除 RPC 方法

### 10.3 版本隔离方案

在 proto package 层面做版本隔离：

```protobuf
// v1
syntax = "proto3";
package library.v1;

// v2
syntax = "proto3";
package library.v2;
```

两个版本在同一个服务端共存，逐步迁移调用方后下线 v1。Google Cloud API 和 Kubernetes API 均采用此方案。

---

## 11. 反模式清单

| 反模式 | 问题 | 正确做法 |
|--------|------|---------|
| `List` + 单数名词 | `ListBook` | `ListBooks` |
| Noun + Verb | `BookCreate`、`TextTranslate` | `CreateBook`、`TranslateText` |
| 跨 RPC 复用消息 | `Get` 和 `List` 共用同一个 Response | 每个 RPC 独立 Request/Response |
| 空 Response | `returns (Empty)` 形式上空消息 | 返回完整资源或明确的结果状态 |
| 笼统动词 | `Process`、`Handle`、`Do` | 用具体的 `Publish`、`Archive`、`Approve` |
| 非标准动词 | `Modify`、`Edit`、`Fetch`、`Query` | `Update`、`Get`、`List`；`Add`/`Remove` 仅用于关联操作（如 `AddBookToShelf`） |
| 首字母缩写词大小写错误（消息名） | `GetOauthUrl` → `Oauth` | `GetOAuthUrl` |
| 首字母缩写词大小写错误（字段名） | `user_ID` | `user_id` |
| List Response 字段名与类型不对齐 | `repeated Book items` | `repeated Book books` |
| 字段编号不连续或重排 | 删除字段后立即重用编号 | 保留废弃字段编号，新字段用新编号 |
| enum 使用不当 | 未定义 UNSPECIFIED 零值、零值直接作为业务有效状态 | 遵循 enum 规范：0 值为占位、结合 optional 使用（参见 9.1 节） |
| 用 gRPC error 返回业务校验失败 | `return status.Error(...)` | 业务错误放在 Response body 中 |

---

## 12. 合规检查清单

用于 Code Review 时逐项检查：

| # | 检查项 | 合规标准 |
|---|--------|---------|
| 1 | **方法命名** | `List{Resources}`(复数) / `Get{Resource}` / `Create{Resource}` / `Update{Resource}` / `Delete{Resource}`，自定义用 VerbNoun |
| 2 | **消息命名** | 每个 RPC 有专属 `{Method}Request` / `{Method}Response`，不跨 RPC 复用 |
| 3 | **字段命名** | snake_case，首字母缩写词全小写（`user_id`），时间字段以 `_at` 结尾 |
| 4 | **字段类型** | 时间用 `int64` 毫秒时间戳；List Response 中 `repeated` 字段用资源名复数 |
| 5 | **分页** | cursor 分页用 `page_token` + `page_size`，`next_page_token` 为空表示最后一页；offset 分页用 `page` + `page_size` + `total`。全局统一 |
| 6 | **文件结构** | syntax → package → go_package → 共享消息 → 核心模型 → 请求/响应 → service |
| 7 | **RPC 排序** | List → Get → Create → Update → Delete → Batch → 自定义 |
| 8 | **响应设计** | Create/Update 返回完整资源；Delete 返回 Empty；List 返回集合 + 分页游标 |
| 9 | **enum 使用** | 0 值为 `UNSPECIFIED` 占位，结合 `optional` 使用，零值不作为业务有效值 |
| 10 | **错误处理** | gRPC status code 仅用于系统级错误，业务校验失败放 Response body |
| 11 | **注释** | 极简：字段标含义+枚举值，RPC 标核心逻辑，禁止设计说明/权限标注类注释 |
| 12 | **版本兼容** | 新增字段用新编号不重用废弃编号，不兼容变更走新 package 版本（`xxx.v2`） |

---

## 参考文献

| 文档 | 链接 |
|------|------|
| Google AIP 总览 | https://google.aip.dev/ |
| AIP-121 资源导向设计 | https://google.aip.dev/121 |
| AIP-122 资源命名 | https://google.aip.dev/122 |
| AIP-131 标准方法 | https://google.aip.dev/131 |
| AIP-132 List | https://google.aip.dev/132 |
| AIP-133 Create | https://google.aip.dev/133 |
| AIP-134 Update | https://google.aip.dev/134 |
| AIP-135 Delete | https://google.aip.dev/135 |
| AIP-136 自定义方法 | https://google.aip.dev/136 |
| AIP-155 批量方法 | https://google.aip.dev/155 |
| AIP-181 版本管理 | https://google.aip.dev/181 |
| Protobuf 风格指南 | https://protobuf.dev/programming-guides/style/ |
| gRPC API 设计 | https://grpc.io/blog/api-design/ |
| Google Cloud APIs 设计指南 | https://cloud.google.com/apis/design |
