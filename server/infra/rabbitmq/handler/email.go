package handler

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rabbitmq"
)

const (
	// email交换机
	EMAIL_EXCHANGE = "email_exchange"
	// email队列
	EMAIL_QUEUE = "email_queue"
)

type EmailHandler struct {
	mq *rabbitmq.RabbitMQ
}

func NewEmailHandler(url string) *EmailHandler {
	mq := rabbitmq.NewRabbitMQ(url)
	mq.BindQueue(EMAIL_QUEUE).BindExchange(rabbitmq.Fanout, EMAIL_EXCHANGE, "email")
	return &EmailHandler{
		mq: mq,
	}
}

func (s *EmailHandler) Publisher() rabbitmq.MessagePublisher {
	return s
}

func (s *EmailHandler) Subscriber() rabbitmq.MessageSubscriber {
	return s
}

func (s *EmailHandler) SendMessage(message string) error {
	err := s.mq.PublishMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmailHandler) ReceiveMessage(handler func(message string)) error {
	//handler := func(message string) {
	//	log.Printf("Received a message: %s", message)
	//}

	err := s.mq.SubscribeMessage(handler)
	if err != nil {
		return err
	}

	return nil
}
