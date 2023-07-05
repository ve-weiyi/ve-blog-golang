package rabbitmq

type MessagePublisher interface {
	// 发送消息
	SendMessage(message string) error
}

type MessageSubscriber interface {
	// 接收消息
	ReceiveMessage(handler func(message string)) error
}
