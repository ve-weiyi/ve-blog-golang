package rabbitmq

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *RabbitmqConn
	opts PublisherOptions

	mu sync.Mutex
}

func NewPublisher(conn *RabbitmqConn, opts ...PublisherOption) *Publisher {
	opt := PublisherOptions{
		Exchange:   "",
		RoutingKey: "",
		Mandatory:  false,
		Immediate:  false,
	}

	for _, o := range opts {
		o(&opt)
	}

	return &Publisher{
		conn: conn,
		opts: opt,
	}
}

func (r *Publisher) PublishMessage(message []byte) error {
	err := r.conn.Publish(amqp.Publishing{
		ContentType: "text/plain",
		Body:        message,
	}, r.opts)

	if err != nil {
		log.Printf("rabbitmq publish message failed: %v", err)
		return err
	}

	log.Printf("rabbitmq publish message success")
	return nil
}
