# 用户模块 RESTful API 设计规范

# 一、通用设计原则

## 1.1 基础约定

- 协议：HTTPS；版本控制：`/api/v{版本号}/`（当前v1），迭代时升级版本避免兼容问题。

- 命名：路径用小写+连字符（kebab-case），操作名用驼峰（CamelCase）；集合资源用复数，单个资源用单数。

- 编码：UTF-8；鉴权：非公开接口需通过请求头 `Authorization: Bearer {token}` 传递JWT令牌。

## 1.2 幂等性与高可用

- 幂等性：PUT/PATCH/DELETE多次调用结果一致；POST接口需携带 `x-request-id` 去重（有效期24小时）。

- 限流：单用户/单IP每分钟最多60次请求，超量返回429；熔断：接口失败率超50%触发，30秒内返回503。

## 1.3 统一响应格式

### 1.3.1 成功响应（200/201）

基础成功响应：

```json

{
  "code": 200,
  "msg": "success",
  "data": {},
  "trace_id": "xxxx"
}
```

分页列表成功响应：

```json

{
  "code": 200,
  "msg": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "page_size": 10
  },
  "trace_id": "xxxx"
}
```

### 1.3.2 错误响应（4xx/5xx）

```json

{
  "code": 400,
  "msg": "参数错误：用户ID不能为空",
  "data": null,
  "trace_id": "xxxx"
}
```

### 1.3.3 常用状态码映射

| 状态码 | 说明    | 典型场景            |
|-----|-------|-----------------|
| 200 | 请求成功  | 查询、更新、删除成功      |
| 201 | 创建成功  | 单个用户创建、批量创建成功   |
| 400 | 参数错误  | 必填参数缺失、参数格式错误   |
| 401 | 未授权   | 未携带令牌、令牌过期、令牌无效 |
| 403 | 权限不足  | 无操作目标资源的权限      |
| 404 | 资源不存在 | 查询的用户ID不存在      |
| 409 | 资源冲突  | 用户名/邮箱已被注册      |
| 429 | 请求限流  | 单位时间内请求次数超出限制   |
| 500 | 服务端异常 | 服务端逻辑错误、数据库异常   |
| 503 | 服务不可用 | 接口熔断、服务降级       |

## 1.4 字段脱敏规范

- 脱敏：手机号（138****1234）、邮箱（zh****@163.com）等敏感信息按规则处理。

- 日志：记录请求响应详情，敏感操作日志留存90天，错误日志记录堆栈。


## 1.5 日志埋点要求

- 需记录的日志信息：请求URL、请求方法、请求参数、响应结果、用户ID、IP地址、请求耗时、请求唯一标识（requestId）

- 错误日志：额外记录错误堆栈信息，便于排查问题

- 敏感操作日志：用户创建、删除、密码修改、绑定关系变更等操作，需单独记录操作轨迹，留存至少90天

# 二、接口详细定义

## 2.1 用户基础 CRUD

| 请求方法   | 接口路径               | 接口说明     | 操作名        | 特殊说明                      |
|--------|--------------------|----------|------------|---------------------------|
| GET    | /api/v1/users      | 获取用户列表   | ListUsers  | 支持分页、筛选、排序，需管理员权限         |
| POST   | /api/v1/users      | 创建单个用户   | CreateUser | 需管理员权限，用户名/邮箱需唯一，返回201状态码 |
| GET    | /api/v1/users/{id} | 获取指定用户信息 | GetUser    | 管理员可查所有用户，普通用户仅可查自己       |
| PUT    | /api/v1/users/{id} | 完整更新用户信息 | UpdateUser | 需管理员权限，需传递所有必填字段          |
| PATCH  | /api/v1/users/{id} | 部分更新用户信息 | PatchUser  | 需管理员权限，仅传递需更新的字段          |
| DELETE | /api/v1/users/{id} | 删除用户     | DeleteUser | 需管理员权限，幂等性支持（重复删除返回成功）    |

## 2.2 批量操作

| 请求方法   | 接口路径                           | 接口说明   | 操作名                  | 特殊说明                                     |
|--------|--------------------------------|--------|----------------------|------------------------------------------|
| POST   | /api/v1/users/search           | 复杂查询用户 | SearchUsers          | 支持多条件组合筛选，返回分页结果，需管理员权限                  |
| POST   | /api/v1/users/batch            | 批量创建用户 | BatchCreateUsers     | 需管理员权限，请求体传递用户列表，单次最多创建50条               |
| PUT    | /api/v1/users/batch            | 批量完整更新 | BatchUpdateUsers     | 需管理员权限，请求体传递用户ID列表+完整字段，单次最多更新50条        |
| PATCH  | /api/v1/users/batch            | 批量部分更新 | BatchPatchUsers      | 需管理员权限，请求体传递用户ID列表+待更新字段，单次最多更新50条       |
| DELETE | /api/v1/users?id=1&id=2        | 简单批量删除 | BatchDeleteUsers     | 需管理员权限，单次最多删除50条，幂等性支持                   |
| POST   | /api/v1/users/activation/batch | 批量激活用户 | BatchActivateUsers   | 需管理员权限，请求体传递用户ID列表，激活后状态变为active         |
| POST   | /api/v1/users/activation/batch | 批量停用用户 | BatchDeactivateUsers | 需管理员权限，请求体传递用户ID列表，停用后状态变为inactive，不删除数据 |

## 2.3 当前用户（/me 简化路径）

### 2.3.1 基础信息

| 请求方法  | 接口路径                     | 接口说明       | 操作名             | 特殊说明                  |
|-------|--------------------------|------------|-----------------|-----------------------|
| GET   | /api/v1/users/me         | 获取当前用户信息   | GetMe           | 无需传递用户ID，根据令牌解析当前用户   |
| PUT   | /api/v1/users/me         | 完整更新当前用户   | UpdateMe        | 需传递所有必填字段，不可修改管理员专属字段 |
| PATCH | /api/v1/users/me         | 部分更新当前用户   | PatchMe         | 仅传递需更新的字段，不可修改管理员专属字段 |
| GET   | /api/v1/users/me/profile | 获取当前用户详细资料 | GetMeProfile    | 返回用户完整资料（含脱敏后的敏感信息）   |
| PUT   | /api/v1/users/me/profile | 更新当前用户详细资料 | UpdateMeProfile | 支持更新昵称、性别、生日等非核心字段    |

### 2.3.2 账号安全

| 请求方法 | 接口路径                      | 接口说明     | 操作名              | 特殊说明                                         |
|------|---------------------------|----------|------------------|----------------------------------------------|
| PUT  | /api/v1/users/me/avatar   | 更新当前用户头像 | UpdateMeAvatar   | 请求体传递图片文件，支持格式：jpg/png，大小不超过5MB              |
| PUT  | /api/v1/users/me/password | 修改当前用户密码 | UpdateMePassword | 请求体需传递原密码、新密码，新密码需符合复杂度要求（8-20位，含字母+数字+特殊符号） |
| GET  | /api/v1/users/me/settings | 获取当前用户设置 | GetMeSettings    | 返回通知设置、隐私设置等配置                               |
| PUT  | /api/v1/users/me/settings | 更新当前用户设置 | UpdateMeSettings | 支持修改通知开关、隐私权限等配置                             |

### 2.3.3 绑定管理

| 请求方法   | 接口路径                                 | 接口说明     | 操作名                | 特殊说明                                  |
|--------|--------------------------------------|----------|--------------------|---------------------------------------|
| GET    | /api/v1/users/me/bindings            | 获取所有绑定信息 | GetMeBindings      | 返回邮箱、手机、第三方平台的绑定状态                    |
| POST   | /api/v1/users/me/bindings/email      | 绑定邮箱     | BindMeEmail        | 需传递邮箱+验证码，邮箱需未被注册                     |
| PATCH  | /api/v1/users/me/bindings/email      | 更换邮箱     | UpdateMeEmail      | 需传递原邮箱验证码+新邮箱+新邮箱验证码                  |
| DELETE | /api/v1/users/me/bindings/email      | 解绑邮箱     | UnbindMeEmail      | 需验证密码，解绑后不可用邮箱登录                      |
| POST   | /api/v1/users/me/bindings/phone      | 绑定手机     | BindMePhone        | 需传递手机号+验证码，手机号需未被注册                   |
| PATCH  | /api/v1/users/me/bindings/phone      | 更换手机     | UpdateMePhone      | 需传递原手机号验证码+新手机号+新手机号验证码               |
| DELETE | /api/v1/users/me/bindings/phone      | 解绑手机     | UnbindMePhone      | 需验证密码，解绑后不可用手机登录                      |
| POST   | /api/v1/users/me/bindings/{platform} | 绑定第三方平台  | BindMeThirdParty   | platform枚举：wechat、github、qq，需传递第三方授权码 |
| DELETE | /api/v1/users/me/bindings/{platform} | 解绑第三方平台  | UnbindMeThirdParty | 需验证密码，解绑后不可用该平台登录                     |

### 2.3.4 权限相关

| 请求方法 | 接口路径                         | 接口说明        | 操作名              | 特殊说明                 |
|------|------------------------------|-------------|------------------|----------------------|
| GET  | /api/v1/users/me/roles       | 获取当前用户角色    | GetMeRoles       | 返回角色ID、角色名称、角色描述     |
| GET  | /api/v1/users/me/permissions | 获取当前用户权限    | GetMePermissions | 返回权限标识、权限名称、权限描述列表   |
| GET  | /api/v1/users/me/menus       | 获取当前用户菜单权限  | GetMeMenus       | 返回可见菜单的层级结构、路由、图标等信息 |
| GET  | /api/v1/users/me/apis        | 获取当前用户API权限 | GetMeApis        | 返回可访问的API路径、请求方法列表   |

### 2.3.5 活动记录

| 请求方法 | 接口路径                             | 接口说明   | 操作名                | 特殊说明                          |
|------|----------------------------------|--------|--------------------|-------------------------------|
| GET  | /api/v1/users/me/login-histories | 获取登录历史 | GetMeLoginHistory  | 支持分页、时间范围筛选，返回登录时间、IP、设备、登录状态 |
| GET  | /api/v1/users/me/activity-logs   | 获取操作日志 | GetMeActivityLogs  | 支持分页、操作类型筛选，返回操作时间、操作内容、操作结果  |
| GET  | /api/v1/users/me/notifications   | 获取通知列表 | GetMeNotifications | 支持分页、通知类型筛选，可通过参数筛选已读/未读      |

### 2.3.6 社交关系

| 请求方法 | 接口路径                       | 接口说明    | 操作名            | 特殊说明                  |
|------|----------------------------|---------|----------------|-----------------------|
| GET  | /api/v1/users/me/following | 获取我关注的人 | GetMeFollowing | 支持分页、搜索筛选             |
| GET  | /api/v1/users/me/followers | 获取关注我的人 | GetMeFollowers | 支持分页、搜索筛选             |
| GET  | /api/v1/users/me/friends   | 获取好友列表  | GetMeFriends   | 支持分页、分组筛选，返回好友备注、在线状态 |

### 2.3.7 用户内容

| 请求方法 | 接口路径                       | 接口说明   | 操作名            | 特殊说明               |
|------|----------------------------|--------|----------------|--------------------|
| GET  | /api/v1/users/me/orders    | 获取我的订单 | GetMeOrders    | 支持分页、订单状态筛选        |
| GET  | /api/v1/users/me/articles  | 获取我的文章 | GetMeArticles  | 支持分页、文章状态（草稿/发布）筛选 |
| GET  | /api/v1/users/me/favorites | 获取我的收藏 | GetMeFavorites | 支持分页、收藏类型筛选        |
| GET  | /api/v1/users/me/addresses | 获取我的地址 | GetMeAddresses | 支持分页，返回默认地址标识      |
