# StompWS - STOMP WebSocket 服务器

基于 STOMP 协议和 WebSocket 实现的高性能聊天室服务器，完整支持 STOMP 1.0/1.1/1.2 协议规范。

## 📚 项目简介

StompWS 是 ve-blog-golang 项目的 WebSocket 聊天室模块，提供基于 STOMP 协议的实时通信能力，支持群聊、私聊、消息队列等功能。

## ✨ 核心特性

### STOMP 协议支持

- ✅ **标准命令**：CONNECT/STOMP、SUBSCRIBE/UNSUBSCRIBE、SEND/MESSAGE、ACK/NACK、BEGIN/COMMIT/ABORT、RECEIPT、ERROR、DISCONNECT
- ✅ **版本协商**：支持 STOMP 1.0/1.1/1.2 版本自动协商
- ✅ **心跳机制**：读写超时检测，自动清理死连接
- ✅ **事务支持**：完整的事务管理（BEGIN/COMMIT/ABORT）
- ✅ **消息确认**：支持 auto/client/client-individual 三种确认模式

### 消息模式

- **Topic（广播）**：`/topic/*` - 发布订阅模式，一对多消息广播
- **Queue（点对点）**：`/queue/*` - 消息队列模式，负载均衡分发
- **私聊**：`/user/{username}` - 一对一私密通信

### 架构特性

- **高并发设计**：三协程模型（读/写/处理分离），非阻塞消息发送
- **可靠性保证**：心跳超时检测、未确认消息重入队、优雅断开连接
- **可扩展性**：插件化设计，支持自定义认证、事件钩子、日志系统
- **安全性**：支持身份验证、防暴力破解、连接状态管理

## 🏗️ 项目结构

```
stompws/
├── main.go                    # 服务入口
├── logws/                     # 日志系统
│   └── logger.go             # Logger 接口与实现
├── server/                    # 服务端实现
│   ├── client/               # 客户端连接管理
│   │   ├── server.go        # STOMP 服务器
│   │   ├── client.go        # 客户端连接
│   │   ├── handlers.go      # STOMP 命令处理
│   │   ├── auth.go          # 身份验证
│   │   ├── hook.go          # 事件钩子
│   │   ├── subscription.go  # 订阅管理
│   │   ├── tx_store.go      # 事务存储
│   │   └── validator.go     # 帧验证器
│   ├── topic/               # Topic 管理器（广播）
│   │   ├── manager.go       # Topic 管理
│   │   └── topic.go         # Topic 实现
│   └── queue/               # Queue 管理器（点对点）
│       ├── manager.go       # Queue 管理
│       ├── queue.go         # Queue 实现
│       └── storage.go       # 消息存储接口
└── web/                      # Web 客户端
    └── client.html          # 测试页面
```

## 🚀 快速开始

### 环境要求

- Go 1.20+
- 支持 WebSocket 的浏览器

### 启动服务

```bash
# 进入项目目录
cd stompws

# 安装依赖
go mod tidy

# 启动服务
go run main.go
```

服务将在 `http://localhost:9091` 启动，访问 `http://localhost:9091` 可打开测试页面。

### 基本使用

```go
package main

import (
	"net/http"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/client"
	"github.com/ve-weiyi/ve-blog-golang/stompws/logws"
)

func main() {
	// 创建服务器
	server := client.NewStompHubServer(
		client.WithEventHooks(client.NewDefaultEventHook()),
		client.WithAuthenticator(client.NewNoAuthenticator()),
		client.WithLogger(logws.NewDefaultLogger()),
	)

	// 注册 WebSocket 路由
	http.HandleFunc("/websocket", server.HandleWebSocket)

	// 启动服务
	http.ListenAndServe(":9091", nil)
}
```

## 🔧 配置选项

### 身份验证

```go
// 无需认证（默认）
server := client.NewStompHubServer(
client.WithAuthenticator(client.NewNoAuthenticator()),
)

// 密码认证
auth := client.NewPasswordAuthenticator()
auth.AddUser("alice", "password123")
auth.AddUser("bob", "secret456")

server := client.NewStompHubServer(
client.WithAuthenticator(auth),
)
```

### 事件钩子

```go
type MyEventHook struct {
client.DefaultEventHook
}

func (h *MyEventHook) OnConnect(server *client.StompHubServer, c *client.Client) {
// 用户连接时的自定义逻辑
}

func (h *MyEventHook) OnSend(server *client.StompHubServer, c *client.Client, message *frame.Frame) bool {
// 消息发送前的拦截处理
return true // 返回 false 可阻止消息发送
}

server := client.NewStompHubServer(
client.WithEventHooks(&MyEventHook{}),
)
```

### 自定义日志

```go
type MyLogger struct{}

func (l *MyLogger) Infof(format string, args ...interface{}) {
// 自定义日志实现
}

server := client.NewStompHubServer(
client.WithLogger(&MyLogger{}),
)
```

## 📖 使用示例

### JavaScript 客户端

```javascript
// 连接服务器
const socket = new WebSocket('ws://localhost:9091/websocket');
const client = Stomp.over(socket);

// 连接并订阅
client.connect({}, function (frame) {
    console.log('Connected:', frame);

    // 订阅 Topic
    client.subscribe('/topic/chat', function (message) {
        console.log('Received:', message.body);
    });

    // 发送消息
    client.send('/topic/chat', {}, JSON.stringify({
        username: 'Alice',
        content: 'Hello World!'
    }));
});
```

### 事务示例

```javascript
// 开始事务
const tx = 'tx-' + Date.now();
client.begin(tx);

// 批量发送（原子操作）
client.send('/queue/orders', {transaction: tx}, order1);
client.send('/queue/orders', {transaction: tx}, order2);
client.send('/queue/notifications', {transaction: tx}, notification);

// 提交事务
client.commit(tx);

// 或回滚事务
// client.abort(tx);
```

## 🎯 技术亮点

1. **完整的 STOMP 协议实现**：严格遵循 STOMP 1.0/1.1/1.2 规范，支持所有标准命令
2. **高并发架构**：三协程模型（读/写/处理分离）+ Channel 通信，避免锁竞争
3. **可靠性保证**：心跳检测、消息重入队、事务支持、优雅断开
4. **插件化设计**：可自定义认证器、事件钩子、日志系统、消息存储
5. **代码质量**：清晰的职责分离、完善的错误处理、详细的代码注释

## 📦 依赖

```go
require (
github.com/go -stomp/stomp/v3 v3.1.5
github.com/gorilla/websocket v1.5.3
)
```

## 🔗 相关链接

- [STOMP 协议规范](https://stomp.github.io/)
- [ve-blog-golang 主项目](https://github.com/ve-weiyi/ve-blog-golang)
- [详细文档](./server/README.md)

## 📄 开源协议

MIT License

## 🙏 致谢

本项目参考了 [go-stomp/stomp](https://github.com/go-stomp/stomp) 的设计思想。
