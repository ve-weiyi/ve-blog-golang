package initialize

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ve-weiyi/ve-blog-golang/server/config"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
)

func ConnectRabbitMq(c config.RabbitMQConf) (*rabbitmq.RabbitmqConn, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", c.Username, c.Password, c.Host, c.Port)

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
		return nil, fmt.Errorf("rabbitmq 初始化失败: %v", err)
	}

	return mq, nil
}

// 订阅消息
func SubscribeMessage(c config.Config) {
	r := c.RabbitMQ

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Username, r.Password, r.Host, r.Port)
	// 消息订阅者需要声明交换机和队列
	mq := rabbitmq.NewRabbitmqConn(url,
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
	err := mq.Connect(nil)
	if err != nil {
		log.Fatal("rabbitmq 初始化失败!", err)
	}

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

	//订阅消息队列，发送邮件
	err = mq.SubscribeMessage(func(message []byte) (err error) {
		var msg mail.EmailMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			return err
		}

		err = emailSender.DeliveryEmail(msg)
		if err != nil {
			log.Println("邮件发送失败!", err)
		}
		return err
	})
	if err != nil {
		log.Fatal("订阅消息失败!", err)
	}
}
