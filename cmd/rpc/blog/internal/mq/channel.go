package mq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/vkit/adapter/mqx"
	"github.com/ve-weiyi/vkit/adapter/mqx/rabbitmqx"
)

// ErrUnavailable 消息队列未就绪（RabbitMQ 不可用），消息未投递。
// 调用方据此区分「投递能力缺失」与「投递失败」，前者通常降级处理。
var ErrUnavailable = errors.New("mq: broker unavailable")

// Channel 绑定一个交换机、队列及其收发器，T 为消息体类型。
type Channel[T any] struct {
	exchange   string
	routingKey string
	queue      string
	mq         mqx.MessageQueue
	pub        mqx.Publisher
	sub        mqx.Subscriber
}

// newChannel 建立交换机连接并派生发布器与订阅器。
// 连接失败时通道仍然可用：发布返回 ErrUnavailable，订阅直接退出。
func newChannel[T any](amqpURL, exchange, routingKey, queue string) *Channel[T] {
	ch := &Channel[T]{
		exchange:   exchange,
		routingKey: routingKey,
		queue:      queue,
	}

	conn, err := rabbitmqx.New(&rabbitmqx.Config{
		URL:          amqpURL,
		ExchangeName: exchange,
		ExchangeType: rabbitmqx.ExchangeFanout,
		Durable:      true,
		AutoDelete:   false,
		LazyConnect:  true,
		Logger:       &mqx.LogxLogger{Logger: logx.WithContext(context.Background())},
	})
	if err != nil {
		logx.Errorf("[mq] init %s: %v", exchange, err)
		return ch
	}
	ch.mq = conn

	if ch.pub, err = conn.Publisher(); err != nil {
		logx.Errorf("[mq] create publisher %s: %v", exchange, err)
	}
	if ch.sub, err = conn.Subscriber(); err != nil {
		logx.Errorf("[mq] create subscriber %s: %v", exchange, err)
	}
	return ch
}

// Publish 序列化事件并投递到交换机。
func (ch *Channel[T]) Publish(ctx context.Context, event *T) error {
	if ch.pub == nil {
		return ErrUnavailable
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("mq: marshal %s: %w", ch.routingKey, err)
	}

	err = ch.pub.Publish(ctx, ch.routingKey, &mqx.Message{Body: body})
	if errors.Is(err, mqx.ErrNotConnected) {
		return ErrUnavailable
	}
	return err
}

// Subscribe 在独立 goroutine 中启动消费循环。
// 订阅建立失败只记录日志，不中断服务：其余交换机与 RPC 服务不受影响。
func (ch *Channel[T]) Subscribe(handler func(context.Context, *T) error) {
	go ch.consume(handler)
}

func (ch *Channel[T]) consume(handler func(context.Context, *T) error) {
	if ch.sub == nil {
		logx.Errorf("[mq] %s subscriber unavailable, consume disabled", ch.exchange)
		return
	}

	logx.Infof("[mq] %s subscriber starting", ch.exchange)
	if err := ch.sub.Subscribe(context.Background(), ch.routingKey, chain(handleMessage(ch.exchange, handler)),
		mqx.WithGroup(ch.queue),
		mqx.WithPrefetch(10),
	); err != nil {
		logx.Errorf("[mq] %s subscriber stopped: %v", ch.exchange, err)
	}
}

func (ch *Channel[T]) close() {
	if ch.sub != nil {
		_ = ch.sub.Close()
	}
	if ch.pub != nil {
		_ = ch.pub.Close()
	}
	if ch.mq != nil {
		_ = ch.mq.Close()
	}
}

// handleMessage 解析消息体并交给业务处理。解析失败视为毒消息丢弃，避免无限重投。
func handleMessage[T any](exchange string, handler func(context.Context, *T) error) mqx.Handler {
	return func(ctx context.Context, msg *mqx.Message) error {
		var event T
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			logx.Errorf("[mq] unmarshal message on %s: %v", exchange, err)
			return nil
		}
		return handler(ctx, &event)
	}
}

// chain 为消息处理加上 panic 兜底与重试：先兜 panic 保证消费循环不退出，
// 再重试 3 次。重试耗尽后返回错误，由 mqx 按 AutoAck 语义丢弃该消息，
// 避免毒消息无限重投。
func chain(next mqx.Handler) mqx.Handler {
	return mqx.Chain(
		mqx.Recovery(&mqx.LogxLogger{Logger: logx.WithContext(context.Background())}),
		mqx.Retry(3),
	)(next)
}

// closer 让不同事件类型的通道可以统一关闭。
type closer interface {
	close()
}
