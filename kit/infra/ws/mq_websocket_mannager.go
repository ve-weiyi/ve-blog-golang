package ws

import (
	"github.com/gorilla/websocket"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
)

type MqWebsocketManager struct {
	// 用于升级 HTTP 连接
	websocket.Upgrader
	// 用于存储连接
	connections map[string]*websocket.Conn

	// 消息发布者和订阅者
	Publisher  rabbitmq.MessagePublisher
	Subscriber rabbitmq.MessageSubscriber
}
