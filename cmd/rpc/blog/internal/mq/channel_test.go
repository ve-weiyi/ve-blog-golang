package mq

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/ve-weiyi/vkit/adapter/mqx"
)

type fakePublisher struct {
	topic string
	body  []byte
}

func (f *fakePublisher) Publish(_ context.Context, topic string, msg *mqx.Message, _ ...mqx.PublishOption) error {
	f.topic = topic
	f.body = msg.Body
	return nil
}

func (f *fakePublisher) PublishBatch(context.Context, string, []*mqx.Message, ...mqx.PublishOption) error {
	return nil
}

func (f *fakePublisher) Close() error { return nil }

func TestChannelPublish(t *testing.T) {
	event := &EmailMessageEvent{Email: "a@example.com", Title: "hi"}

	t.Run("发布器不可用时返回 ErrUnavailable", func(t *testing.T) {
		ch := &Channel[EmailMessageEvent]{routingKey: EmailRoutingKey}
		if err := ch.Publish(context.Background(), event); !errors.Is(err, ErrUnavailable) {
			t.Fatalf("err = %v, want ErrUnavailable", err)
		}
	})

	t.Run("按路由键投递序列化后的事件", func(t *testing.T) {
		pub := &fakePublisher{}
		ch := &Channel[EmailMessageEvent]{routingKey: EmailRoutingKey, pub: pub}

		if err := ch.Publish(context.Background(), event); err != nil {
			t.Fatalf("Publish: %v", err)
		}
		if pub.topic != EmailRoutingKey {
			t.Errorf("topic = %q, want %q", pub.topic, EmailRoutingKey)
		}

		var got EmailMessageEvent
		if err := json.Unmarshal(pub.body, &got); err != nil {
			t.Fatalf("unmarshal body: %v", err)
		}
		if got.Email != event.Email || got.Title != event.Title {
			t.Errorf("body = %+v, want %+v", got, *event)
		}
	})
}

func TestHandleMessage(t *testing.T) {
	var handled *EmailMessageEvent
	handler := handleMessage[EmailMessageEvent](EmailExchange,
		func(_ context.Context, event *EmailMessageEvent) error {
			handled = event
			return nil
		})

	t.Run("解析成功后交给业务处理", func(t *testing.T) {
		body, _ := json.Marshal(EmailMessageEvent{Email: "a@example.com"})
		if err := handler(context.Background(), &mqx.Message{Body: body}); err != nil {
			t.Fatalf("handle: %v", err)
		}
		if handled == nil || handled.Email != "a@example.com" {
			t.Fatalf("handled = %+v, want email a@example.com", handled)
		}
	})

	t.Run("毒消息直接丢弃", func(t *testing.T) {
		handled = nil
		if err := handler(context.Background(), &mqx.Message{Body: []byte("not json")}); err != nil {
			t.Fatalf("handle poison message: %v, want nil", err)
		}
		if handled != nil {
			t.Fatalf("handled poison message: %+v", handled)
		}
	})
}
