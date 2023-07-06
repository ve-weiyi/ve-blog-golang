package initialize

import (
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rabbitmq/handler"
)

// 订阅消息
func RabbitMqSubscribe() {
	main()
}

func main() {
	emailMq := handler.NewEmailHandler(global.CONFIG.RabbitMQ.GetUrl())
	emailSender := mail.NewEmailSender(&global.CONFIG.Email)

	emailMq.ReceiveMessage(func(message string) {
		var msg mail.EmailMessage
		jsonconv.JsonToObject(message, &msg)
		err := emailSender.SendEmailMessage(msg)
		if err != nil {
			global.LOG.Error("邮件发送失败!", err)
		}
	})
}
