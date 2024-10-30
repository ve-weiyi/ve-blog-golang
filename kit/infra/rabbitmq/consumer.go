package rabbitmq

import (
	"log"
	"sync"
)

type Consumer struct {
	conn *RabbitmqConn
	opts ConsumerOptions

	mu sync.Mutex
}

func NewConsumer(conn *RabbitmqConn, opts ...ConsumerOption) *Consumer {
	opt := ConsumerOptions{
		Queue:     "",
		Name:      "",
		AutoAck:   false,
		Exclusive: false,
		NoWait:    false,
		NoLocal:   false,
		Args:      nil,
	}

	for _, o := range opts {
		o(&opt)
	}

	return &Consumer{
		conn: conn,
		opts: opt,
	}
}

func (r *Consumer) SubscribeMessage(handler func(message []byte) error) {
	//接收消息
	msgs, err := r.conn.Consume(r.opts)
	if err != nil {
		log.Printf("Error: rabbitmq consume: %v", err)
		return
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for msg := range msgs {
			// 处理消息
			err = handler(msg.Body)
			if err != nil {
				log.Printf("Warning: rabbitmq subscribe handle error: %v", err)

				// 消费失败，根据配置决定是否重入队列
				if err = msg.Nack(false, true); err != nil {
					log.Printf("Warning: rabbitmq nack msg error: %v", err)
				}
				continue
			}

			// AutoAck为true时，不进行手动ack确认
			if r.opts.AutoAck {
				continue
			}

			// 消费成功时确认消息
			if err = msg.Ack(false); err != nil {
				log.Printf("Warning: rabbitmq ack msg error: %v", err)
				// 这里可以决定是否重新处理消息或报警
				continue
			}

			log.Printf("rabbitmq handler message success: %s", msg.MessageId)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return
}
