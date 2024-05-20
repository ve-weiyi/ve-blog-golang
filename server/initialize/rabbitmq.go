package initialize

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

// 订阅消息
func RabbitMq() {
	m := global.CONFIG.RabbitMQ
	url := m.GetUrl()
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

	global.EmailMQ = mq

	go SubscribeMessage()
}

func SubscribeMessage() {
	m := global.CONFIG.RabbitMQ
	url := m.GetUrl()
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

	e := global.CONFIG.Email
	emailSender := mail.NewEmailSender(
		mail.WithHost(e.Host),
		mail.WithPort(e.Port),
		mail.WithUsername(e.Username),
		mail.WithPassword(e.Password),
		mail.WithNickname(e.Nickname),
		mail.WithDeliver(strings.Split(e.Deliver, ",")),
		mail.WithIsSSL(e.IsSSL),
	)

	//订阅消息队列，发送邮件
	err = mq.SubscribeMessage(func(message []byte) (err error) {
		var msg mail.EmailMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			return err
		}

		err = emailSender.SendEmailMessage(msg)
		if err != nil {
			global.LOG.Error("邮件发送失败!", err)
		}
		return err
	})
	if err != nil {
		log.Fatal("订阅消息失败!", err)
	}
}
