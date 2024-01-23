package rabbitmq

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 连接信息amqp://guest:guest@127.0.0.1:5672/guest这个信息是固定不变的amqp://是固定参数 后面两个是用户名密码ip地址端口号Virtual Host
const rabbitMQURL = "amqp://guest:guest@localhost:5672/"

func TestWorking_SimplePublish(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("code")
	for i := 0; i <= 100; i++ {
		err := rabbitmq.PublishMessage("Hello,RabbitMQ!" + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func TestWorking_SimpleConsumer(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("code")
	rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
}

func TestWorking_SimpleConsumer2(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("code")
	rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
}

func TestWorking_PublishTopic(t *testing.T) {
	rabbitmq1 := NewRabbitMQ(rabbitMQURL).BindExchange(Topic, "email", "email.code.login")
	rabbitmq2 := NewRabbitMQ(rabbitMQURL).BindExchange(Topic, "email", "email.code.register")
	for i := 0; i <= 100; i++ {
		err := rabbitmq1.PublishMessage("Hello,RabbitMQ 1!" + strconv.Itoa(i))
		err = rabbitmq2.PublishMessage("Hello,RabbitMQ 2!" + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func TestWorking_ConsumerTopic(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("").BindExchange(Topic, "email", "#")
	err := rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestWorking_ConsumerTopic2(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("").BindExchange(Topic, "email", "email.*.register")
	err := rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestWorking_PublishFanout(t *testing.T) {
	rabbitmq1 := NewRabbitMQ(rabbitMQURL).BindExchange(Fanout, "email_fanout")
	for i := 0; i <= 100; i++ {
		err := rabbitmq1.PublishMessage("Hello,RabbitMQ!" + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func TestWorking_ConsumerFanout(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("").BindExchange(Fanout, "email_fanout")
	err := rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestWorking_ConsumerFanout2(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("").BindExchange(Fanout, "email_fanout")
	err := rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestWorking_PublishDirect(t *testing.T) {
	rabbitmq1 := NewRabbitMQ(rabbitMQURL).BindQueue("").BindExchange(Direct, "email_direct", "email.code.login")
	rabbitmq2 := NewRabbitMQ(rabbitMQURL).BindQueue("").BindExchange(Direct, "email_direct", "email.code.register")
	for i := 0; i <= 100; i++ {
		err := rabbitmq1.PublishMessage("Hello,RabbitMQ login!" + strconv.Itoa(i))
		err = rabbitmq2.PublishMessage("Hello,RabbitMQ register!" + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func TestWorking_ConsumerDirect(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("").BindExchange(Direct, "email_direct", "email.code.login")
	err := rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestWorking_ConsumerDirect2(t *testing.T) {
	rabbitmq := NewRabbitMQ(rabbitMQURL)
	rabbitmq.BindQueue("").BindExchange(Direct, "email_direct", "email.code.register")
	err := rabbitmq.SubscribeMessage(func(message string) {
		fmt.Println(message)
	})
	if err != nil {
		fmt.Println(err)
	}
}
