package mq

import (
	"fmt"

	"github.com/ve-weiyi/vkit/adapter/mqx"
	"github.com/ve-weiyi/vkit/adapter/mqx/rabbitmqx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

const (
	LoginExchange   = "blog-login-log-exchange"
	LoginQueue      = "blog-login-log-queue"
	LoginRoutingKey = "login"
)

func initLoginMq(svcCtx *svc.ServiceContext) (mqx.MessageQueue, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		svcCtx.Config.RabbitMQConf.Username,
		svcCtx.Config.RabbitMQConf.Password,
		svcCtx.Config.RabbitMQConf.Host,
		svcCtx.Config.RabbitMQConf.Port)

	return rabbitmqx.NewRabbitMQ(&rabbitmqx.Config{
		URL:          url,
		ExchangeName: LoginExchange,
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
