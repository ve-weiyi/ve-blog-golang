package initialize

import (
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

const (
	// email交换机
	EMAIL_EXCHANGE = "email_exchange"
	// email队列
	EMAIL_QUEUE = "email_queue"
)

// 订阅消息
func RabbitMq() {
	url := global.CONFIG.RabbitMQ.GetUrl()
	mq := rabbitmq.NewRabbitMQ(url)
	mq.BindQueue(EMAIL_QUEUE).BindExchange(rabbitmq.Fanout, EMAIL_EXCHANGE, "email")

	global.EmailMQ = mq
	go SubscribeMessage()
}

func SubscribeMessage() {
	cfg := global.CONFIG.Email
	emailSender := &mail.EmailSender{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		Password: cfg.Password,
		Nickname: cfg.Nickname,
		Deliver:  strings.Split(cfg.Deliver, ","),
		IsSSL:    cfg.IsSSL,
	}

	//订阅消息队列，发送邮件
	global.EmailMQ.SubscribeMessage(func(message string) {
		var msg mail.EmailMessage
		jsonconv.JsonToObject(message, &msg)
		err := emailSender.SendEmailMessage(msg)
		if err != nil {
			global.LOG.Error("邮件发送失败!", err)
		}
	})
}
