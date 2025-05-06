package stompws

import (
	"context"
)

// BroadcastMessage 广播消息结构
type BroadcastMessage struct {
	Topic   string
	Message string
	Sender  *Client
	Context context.Context
}
