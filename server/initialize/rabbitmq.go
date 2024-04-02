package initialize

import (
	"fmt"
	"log"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mq/rabbitmqx"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
)

func InitEmailDeliver(c *config.Config) (*mail.MqEmailDeliver, error) {
	e := c.Email
	emailSender := mail.NewEmailDeliver(
		mail.WithHost(e.Host),
		mail.WithPort(e.Port),
		mail.WithUsername(e.Username),
		mail.WithPassword(e.Password),
		mail.WithNickname(e.Nickname),
		mail.WithDeliver(e.Deliver),
		mail.WithSSL(e.IsSSL),
	)

	r := c.RabbitMQ
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Username, r.Password, r.Host, r.Port)
	// 创建连接
	conn, err := rabbitmqx.NewRabbitmqConn(url, nil)
	if err != nil {
		log.Fatal("rabbitmq 初始化失败!", err)
	}

	queue := &rabbitmqx.QueueOptions{
		Name:    constant.EmailQueue,
		Durable: true, // 是否持久化
	}

	exchange := &rabbitmqx.ExchangeOptions{
		Name:    constant.EmailExchange,
		Kind:    rabbitmqx.ExchangeTypeFanout,
		Durable: true, // 是否持久化
	}

	binding := &rabbitmqx.BindingOptions{
		RoutingKey: "",
	}

	err = conn.Declare(queue, exchange, binding)
	if err != nil {
		log.Fatal(err)
	}

	// pub/sub模式 消息发布者只需要声明交换机
	pb := rabbitmqx.NewRabbitmqProducer(conn,
		rabbitmqx.WithPublisherExchange(exchange.Name),
		rabbitmqx.WithPublisherMandatory(true),
	)

	// pub/sub模式 消息订阅者需要声明交换机和队列
	sb := rabbitmqx.NewRabbitmqConsumer(
		conn,
		rabbitmqx.WithConsumerQueue(queue.Name),
		rabbitmqx.WithConsumerAutoAck(true),
	)

	deliver := mail.NewMqEmailDeliver(emailSender, pb, sb)
	// 订阅消息
	go deliver.SubscribeEmail()

	return deliver, nil
}
