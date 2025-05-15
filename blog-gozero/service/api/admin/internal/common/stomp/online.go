package stomp

import (
	"fmt"
	"time"

	"github.com/go-stomp/stomp/v3/frame"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/client"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/topic"

	"github.com/ve-weiyi/ve-blog-golang/stompws/logws"
)

type OnlineTopicHook struct {
	logger logws.Logger
	count  int64
}

func NewOnlineTopicHook(logger logws.Logger) *OnlineTopicHook {
	return &OnlineTopicHook{
		logger: logger,
	}
}

func (h *OnlineTopicHook) OnTopicSubscribe(t *topic.Topic, sub *client.Subscription) error {
	h.logger.Infof("Topic subscribed :%s", sub.Destination())
	h.count++

	f := frame.New(frame.MESSAGE)
	f.Header.Set(frame.Subscription, sub.Id())
	f.Header.Set(frame.Destination, sub.Destination())
	f.Header.Set(frame.MessageId, fmt.Sprintf("%d", time.Now().UnixMilli()))
	f.Body = []byte(jsonconv.AnyToJsonNE(OnlineEvent{
		Count:    h.count,
		IsOnline: true,
	}))
	t.Enqueue(f)
	return nil
}

func (h *OnlineTopicHook) OnTopicUnsubscribe(t *topic.Topic, sub *client.Subscription) error {
	h.logger.Infof("Topic unsubscribed :%s", sub.Destination())
	h.count--
	f := frame.New(frame.MESSAGE)
	f.Header.Set(frame.Subscription, sub.Id())
	f.Header.Set(frame.Destination, sub.Destination())
	f.Header.Set(frame.MessageId, fmt.Sprintf("%d", time.Now().UnixMilli()))
	f.Body = []byte(jsonconv.AnyToJsonNE(OnlineEvent{
		Count:    h.count,
		IsOnline: false,
	}))
	t.Enqueue(f)
	return nil
}

func (h *OnlineTopicHook) OnTopicPublish(t *topic.Topic, f *frame.Frame) error {
	h.logger.Infof("Topic published :%s, msg:%s", f.Header.Get(frame.Destination), string(f.Command))
	return nil
}
