package mq

import (
	"fmt"

	"github.com/ve-weiyi/vkit/adapter/mqx"
	"github.com/ve-weiyi/vkit/adapter/mqx/rabbitmqx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

const (
	InboxExchange   = "blog-inbox-exchange"
	InboxQueue      = "blog-inbox-queue"
	InboxRoutingKey = "inbox"
)

func initInboxMq(svcCtx *svc.ServiceContext) (mqx.MessageQueue, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		svcCtx.Config.RabbitMQConf.Username,
		svcCtx.Config.RabbitMQConf.Password,
		svcCtx.Config.RabbitMQConf.Host,
		svcCtx.Config.RabbitMQConf.Port)

	return rabbitmqx.NewRabbitMQ(&rabbitmqx.Config{
		URL:          url,
		ExchangeName: InboxExchange,
		ExchangeType: "fanout",
		Durable:      true,
		AutoDelete:   false,
		QueueConfig: &rabbitmqx.QueueConfig{
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
		},
	})
}
