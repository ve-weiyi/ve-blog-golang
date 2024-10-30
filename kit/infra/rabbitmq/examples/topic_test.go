package RabbitMQ

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
)

// 模拟一个topic模式：聊天室

const ExchangeChat = "exchange_chat"

// 发布者
const TopicUserOnline = "blog.chat.room_id.username.online"

const TopicUserMsg = "blog.chat.room_id.username.msg"

// 订阅者
const TopicAllUserOnline = "blog.chat.room_id.*.online"

const TopicAllUserMsg = "blog.chat.room_id.*.msg"

func Test_Topic_Publish(t *testing.T) {
	conn, err := rabbitmq.NewRabbitmqConn(
		"amqp://veweiyi:rabbitmq7914@localhost:5672",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	exchange := &rabbitmq.ExchangeOptions{
		Name:    ExchangeChat,
		Kind:    rabbitmq.ExchangeTypeTopic,
		Durable: true, // 是否持久化
	}

	err = conn.Declare(nil, exchange, nil)
	if err != nil {
		log.Fatal(err)
	}

	//  topic 模式只需要声明交换机和Topic
	mq1 := rabbitmq.NewPublisher(conn,
		rabbitmq.WithPublisherExchange(exchange.Name),
		rabbitmq.WithPublisherRoutingKey(TopicUserOnline),
		rabbitmq.WithPublisherMandatory(true),
	)

	mq2 := rabbitmq.NewPublisher(conn,
		rabbitmq.WithPublisherExchange(exchange.Name),
		rabbitmq.WithPublisherRoutingKey(TopicUserMsg),
		rabbitmq.WithPublisherMandatory(true),
	)

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		err = mq1.PublishMessage([]byte("user online: " + strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
		err = mq2.PublishMessage([]byte("user msg: " + strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func Test_Topic_Subscribe1(t *testing.T) {
	conn, err := rabbitmq.NewRabbitmqConn(
		"amqp://veweiyi:rabbitmq7914@localhost:5672",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queue := &rabbitmq.QueueOptions{
		Name: TopicUserOnline, //不填是随机队列名称
	}

	exchange := &rabbitmq.ExchangeOptions{
		Name:    ExchangeChat,
		Kind:    rabbitmq.ExchangeTypeTopic,
		Durable: true, // 是否持久化
	}

	binding := &rabbitmq.BindingOptions{
		RoutingKey: TopicAllUserOnline,
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

func Test_Topic_Subscribe2(t *testing.T) {
	conn, err := rabbitmq.NewRabbitmqConn(
		"amqp://veweiyi:rabbitmq7914@localhost:5672",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queue := &rabbitmq.QueueOptions{
		Name:    TopicUserMsg, //不填是随机队列名称
		Durable: true,         // 是否持久化
	}

	exchange := &rabbitmq.ExchangeOptions{
		Name:    ExchangeChat,
		Kind:    rabbitmq.ExchangeTypeTopic,
		Durable: true, // 是否持久化
	}

	binding := &rabbitmq.BindingOptions{
		RoutingKey: TopicAllUserMsg,
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
