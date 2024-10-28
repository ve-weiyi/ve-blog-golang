package initialize

import (
	"fmt"
	"log"

	"github.com/ve-weiyi/ve-blog-golang/server/config"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
)

func ConnectRabbitMq(c *config.Config) (*mail.MqEmailDeliver, error) {
	e := c.Email
	emailSender := mail.NewEmailDeliver(
		mail.WithHost(e.Host),
		mail.WithPort(e.Port),
		mail.WithUsername(e.Username),
		mail.WithPassword(e.Password),
		mail.WithNickname(e.Nickname),
		mail.WithDeliver(e.Deliver),
		mail.WithIsSSL(e.IsSSL),
	)

	r := c.RabbitMQ
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Username, r.Password, r.Host, r.Port)
	// 消息发布者只需要声明交换机
	mq := rabbitmq.NewRabbitmqConn(url,
		rabbitmq.Exchange(rabbitmq.ExchangeOptions{
			Name:    constant.EmailExchange,
			Type:    rabbitmq.ExchangeTypeFanout,
			Durable: true,
		}),
		rabbitmq.DisableAutoAck(),
		rabbitmq.Requeue(),
	)

	err := mq.Connect(nil)
	if err != nil {
		log.Fatal("rabbitmq 初始化失败!", err)
	}

	// 消息订阅者需要声明交换机和队列
	sb := rabbitmq.NewRabbitmqConn(url,
		rabbitmq.Queue(rabbitmq.QueueOptions{
			Name:    constant.EmailQueue,
			Durable: true,
			Args:    nil,
		}),
		rabbitmq.Exchange(rabbitmq.ExchangeOptions{
			Name:    constant.EmailExchange,
			Type:    rabbitmq.ExchangeTypeFanout,
			Durable: true,
		}),
		rabbitmq.Key("email"),
	)
	err = sb.Connect(nil)
	if err != nil {
		log.Fatal("rabbitmq 初始化失败!", err)
	}

	deliver := mail.NewMqEmailDeliver(emailSender, sb, sb)

	return deliver, nil
}
