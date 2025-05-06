package stompws

import (
	"context"
)

// TopicHandler 定义主题处理函数
type TopicHandler struct {
	OnMessageSubscribe   func(ctx context.Context, topic string, client *Client) error            // 订阅消息时调用
	OnMessageUnsubscribe func(ctx context.Context, topic string, client *Client) error            // 取消订阅时调用
	OnMessagePublish     func(ctx context.Context, topic, message string, client *Client) error   // 发布消息时调用
	OnMessageDeliver     func(ctx context.Context, topic, message string, client *Client) error   // 推送消息时调用
	OnMessageDrop        func(ctx context.Context, topic, message string, client *Client) error   // 消息被丢弃时调用
	OnMessageAck         func(ctx context.Context, topic, messageID string, client *Client) error // 消息被确认时调用
	OnMessageNack        func(ctx context.Context, topic, messageID string, client *Client) error // 消息被拒绝时调用
}
