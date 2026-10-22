package mq

import (
	"net"
	"net/url"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/config"
)

// Client 持有各交换机的通道，由 ServiceContext 注入使用。
type Client struct {
	Email  *Channel[EmailMessageEvent]
	Sms    *Channel[SmsMessageEvent]
	Inbox  *Channel[InboxMessageEvent]
	Login  *Channel[LoginEvent]
	Logout *Channel[LogoutEvent]
}

// NewClient 建立各交换机的连接与收发器。
// RabbitMQ 不可用时只记录错误并降级：发布返回 ErrUnavailable，订阅直接退出，不阻断服务启动。
func NewClient(c config.RabbitMQConf) *Client {
	amqpURL := buildAMQPURL(c)

	return &Client{
		Email:  newChannel[EmailMessageEvent](amqpURL, EmailExchange, EmailRoutingKey, EmailQueue),
		Sms:    newChannel[SmsMessageEvent](amqpURL, SmsExchange, SmsRoutingKey, SmsQueue),
		Inbox:  newChannel[InboxMessageEvent](amqpURL, InboxExchange, InboxRoutingKey, InboxQueue),
		Login:  newChannel[LoginEvent](amqpURL, LoginExchange, LoginRoutingKey, LoginQueue),
		Logout: newChannel[LogoutEvent](amqpURL, LogoutExchange, LogoutRoutingKey, LogoutQueue),
	}
}

// Close 关闭全部订阅器、发布器与连接。
func (c *Client) Close() {
	for _, ch := range []closer{c.Email, c.Sms, c.Inbox, c.Login, c.Logout} {
		ch.close()
	}
}

// buildAMQPURL 拼装 AMQP 连接串。口令经 url.UserPassword 转义，
// 含 @、:、/ 时不会破坏 URL 解析。
func buildAMQPURL(c config.RabbitMQConf) string {
	u := url.URL{
		Scheme: "amqp",
		User:   url.UserPassword(c.Username, c.Password),
		Host:   net.JoinHostPort(c.Host, c.Port),
		Path:   "/",
	}
	return u.String()
}
