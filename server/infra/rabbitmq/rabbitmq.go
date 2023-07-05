package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 连接信息amqp://kuteng:kuteng@127.0.0.1:5672/kuteng这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host
const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

const (
	Simple = ""       //Simple模式（简单模式，一个生产者对应多个消费者）
	Worker = ""       //Worker模式（工作模式，一个生产者对应多个消费者，一个消息只能被一个消费者获取）
	Fanout = "fanout" //Publish模式（订阅模式，消息被路由投递给多个队列，一个消息被多个消费者获取）
	Direct = "direct" //Direct模式（路由模式，消息被路由投递给符合路由规则的队列，一个消息被一个消费者获取）
	Topic  = "topic"  //Topic模式（主题模式，消息被路由投递给符合通配符的队列，一个消息被一个消费者获取）
)

// rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	//工作模式
	Type string
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Keys []string

	//连接信息
	Mqurl string
}

// 创建结构体实例
func NewRabbitMQ(url string) *RabbitMQ {
	var err error
	r := &RabbitMQ{Mqurl: url}
	//获取connection
	conn, err := amqp.Dial(r.Mqurl)
	r.failOnErr("failed to connect rabbitmq!", err)
	//获取channel
	channel, err := conn.Channel()
	r.failOnErr("failed to open a channel", err)

	r.conn = conn
	r.channel = channel
	//创建RabbitMQ实例
	return r
}

// 绑定队列。simple和working模式只需要声明队列
func (r *RabbitMQ) BindQueue(queueName string) *RabbitMQ {
	var err error
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		// 为空表示随机生产队列名称
		queueName,
		//是否持久化
		true,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	r.failOnErr("Failed to declare a queue: %s", err)
	r.QueueName = q.Name
	log.Println("创建队列成功 ", r.QueueName)
	return r
}

// 绑定交换机。生产者可以只声明交换机
func (r *RabbitMQ) BindExchange(kind string, exchange string, keys ...string) *RabbitMQ {
	var err error
	r.Type = kind
	r.Exchange = exchange
	r.Keys = keys
	if r.Keys == nil {
		r.Keys = []string{""}
	}
	//2.尝试创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange, // name
		r.Type,     // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	r.failOnErr("Failed to declare an exchange: %s", err)
	log.Println("创建交换机成功 ", r.Exchange)

	if r.QueueName != "" {

		// 3.绑定队列到交换机
		for _, s := range r.Keys {
			err = r.channel.QueueBind(
				r.QueueName, // queue name
				//在pub/sub模式下，这里的key要为空
				s,          // routing key
				r.Exchange, // exchange
				false,
				nil)
			log.Printf("Binding queue %s to exchange %s with routing key %s", r.QueueName, r.Exchange, s)
			r.failOnErr("Failed to bind a queue: %s", err)
		}

	}

	return r
}

// 直接模式队列生产
func (r *RabbitMQ) PublishMessage(message string) (err error) {
	if r.channel == nil {
		return fmt.Errorf("channel is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//2.发送消息
	err = r.channel.PublishWithContext(
		ctx,
		r.Exchange,
		//要设置
		r.Keys[0],
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	return err
}

// 话题模式接受消息
// 要注意key,规则
// 其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
// 匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
func (r *RabbitMQ) SubscribeMessage(consume func(message string)) error {
	if r.channel == nil {
		return fmt.Errorf("channel is nil")
	}
	//接收消息
	msgs, err := r.channel.Consume(
		r.QueueName, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	var forever chan struct{}
	//启用协程处理消息
	go func() {
		for d := range msgs {
			fmt.Println("111")
			//消息逻辑处理，可以自行设计逻辑
			consume(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. BindExchange exit press CTRL+C")
	<-forever
	return nil
}

// 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(message string, err error) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
	}
}
