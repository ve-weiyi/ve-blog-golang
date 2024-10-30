package mail

import (
	"encoding/json"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
)

type MqEmailDeliver struct {
	Deliver *EmailDeliver

	Publisher  rabbitmq.MessagePublisher
	Subscriber rabbitmq.MessageSubscriber
}

func NewMqEmailDeliver(deliver *EmailDeliver, pb rabbitmq.MessagePublisher, sb rabbitmq.MessageSubscriber) *MqEmailDeliver {
	return &MqEmailDeliver{
		Deliver:    deliver,
		Publisher:  pb,
		Subscriber: sb,
	}
}

func (m *MqEmailDeliver) DeliveryEmail(msg *EmailMessage) error {
	mq := m.Publisher

	// 序列化消息
	jb, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// 发送邮件
	err = mq.PublishMessage(jb)
	if err != nil {
		return err
	}
	return nil
}

func (m *MqEmailDeliver) SubscribeEmail() {
	mq := m.Subscriber
	deliver := m.Deliver

	handler := func(message []byte) (err error) {
		var msg EmailMessage
		// 反序列化消息
		err = json.Unmarshal(message, &msg)
		if err != nil {
			return err
		}

		// 发送邮件
		err = deliver.DeliveryEmail(&msg)
		if err != nil {
			return err
		}

		return nil
	}

	// 订阅消息队列，发送邮件
	mq.SubscribeMessage(handler)
	return
}
