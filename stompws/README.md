STOMP 的核心概念
概念 说明
CONNECT 客户端连接到 STOMP 服务器（类似 HTTP 的握手）。
SUBSCRIBE 客户端订阅某个消息目的地（如 /topic/chat），接收该目的地的消息。
SEND 客户端向指定目的地（如 /app/chat）发送消息。
UNSUBSCRIBE 取消订阅，不再接收某目的地的消息。
DISCONNECT 关闭连接。
