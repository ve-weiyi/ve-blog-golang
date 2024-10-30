package RabbitMQ

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
)

// 模拟一个fanout模式：邮件订阅

const ExchangeEmail = "email_exchange"

// 发布者
const TopicEmail = "blog.email" // fanout模式不需要routing key

// 订阅者
const QueueEmail = "email_queue" //相同队列名称会争抢消息

const QueueEmail2 = "email_queue2"

func Test_Fanout_Publish(t *testing.T) {
	conn, err := rabbitmq.NewRabbitmqConn(
		"amqp://veweiyi:rabbitmq7914@localhost:5672",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	exchange := &rabbitmq.ExchangeOptions{
		Name:    ExchangeEmail,
		Kind:    rabbitmq.ExchangeTypeFanout,
		Durable: true, // 是否持久化
	}

	err = conn.Declare(nil, exchange, nil)
	if err != nil {
		log.Fatal(err)
	}

	// pub/sub 模式发布者只需要声明交换机
	mq1 := rabbitmq.NewPublisher(conn,
		rabbitmq.WithPublisherExchange(exchange.Name),
		rabbitmq.WithPublisherMandatory(true),
	)

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		err = mq1.PublishMessage([]byte("user email: " + strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func Test_Fanout_Subscribe1(t *testing.T) {
	conn, err := rabbitmq.NewRabbitmqConn(
		"amqp://veweiyi:rabbitmq7914@localhost:5672",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queue := &rabbitmq.QueueOptions{
		Name:    QueueEmail, //不填是随机队列名称
		Durable: true,       // 是否持久化
	}

	exchange := &rabbitmq.ExchangeOptions{
		Name:    ExchangeEmail,
		Kind:    rabbitmq.ExchangeTypeFanout,
		Durable: true, // 是否持久化
	}

	binding := &rabbitmq.BindingOptions{
		RoutingKey: TopicEmail,
	}

	err = conn.Declare(queue, exchange, binding)
	if err != nil {
		log.Fatal(err)
	}

	mq := rabbitmq.NewConsumer(conn,
		rabbitmq.WithConsumerQueue(queue.Name),
		rabbitmq.WithConsumerAutoAck(true),
	)

	mq.SubscribeMessage(func(msg []byte) error {
		log.Printf("receive message: %s", string(msg))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}

func Test_Fanout_Subscribe2(t *testing.T) {
	conn, err := rabbitmq.NewRabbitmqConn(
		"amqp://veweiyi:rabbitmq7914@localhost:5672",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queue := &rabbitmq.QueueOptions{
		Name:    QueueEmail, //不填是随机队列名称
		Durable: true,       // 是否持久化
	}

	exchange := &rabbitmq.ExchangeOptions{
		Name:    ExchangeEmail,
		Kind:    rabbitmq.ExchangeTypeFanout,
		Durable: true, // 是否持久化
	}

	binding := &rabbitmq.BindingOptions{
		RoutingKey: TopicEmail,
	}

	err = conn.Declare(queue, exchange, binding)
	if err != nil {
		log.Fatal(err)
	}

	mq := rabbitmq.NewConsumer(conn,
		rabbitmq.WithConsumerQueue(queue.Name),
		rabbitmq.WithConsumerAutoAck(true),
	)

	mq.SubscribeMessage(func(msg []byte) error {
		log.Printf("receive message: %s", string(msg))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}

func Test_Fanout_Subscribe3(t *testing.T) {
	conn, err := rabbitmq.NewRabbitmqConn(
		"amqp://veweiyi:rabbitmq7914@localhost:5672",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queue := &rabbitmq.QueueOptions{
		Name:    QueueEmail2, //不填是随机队列名称
		Durable: true,        // 是否持久化
	}

	exchange := &rabbitmq.ExchangeOptions{
		Name:    ExchangeEmail,
		Kind:    rabbitmq.ExchangeTypeFanout,
		Durable: true, // 是否持久化
	}

	binding := &rabbitmq.BindingOptions{
		RoutingKey: TopicEmail,
	}

	err = conn.Declare(queue, exchange, binding)
	if err != nil {
		log.Fatal(err)
	}

	mq := rabbitmq.NewConsumer(conn,
		rabbitmq.WithConsumerQueue(queue.Name),
		rabbitmq.WithConsumerAutoAck(true),
	)

	mq.SubscribeMessage(func(msg []byte) error {
		log.Printf("receive message: %s", string(msg))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
