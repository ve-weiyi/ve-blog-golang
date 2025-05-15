package stompws

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-stomp/stomp/v3/frame"
	"github.com/go-stomp/stomp/v3/server/client"

	"github.com/ve-weiyi/ve-blog-golang/stompws/logws"
)

type TopicHook interface {
	OnTopicSubscribe(sub *client.Subscription) error               // 订阅消息时调用
	OnTopicUnsubscribe(sub *client.Subscription) error             // 取消订阅时调用
	OnTopicPublish(sub *client.Subscription, f *frame.Frame) error // 消息发布到 Broker 时
}

type HookManager struct {
	hooks map[string][]TopicHook
}

// NewManager creates a new topic manager.
func NewHookManager() *HookManager {
	tm := &HookManager{hooks: make(map[string][]TopicHook)}
	return tm
}

func (tm *HookManager) RegisterHooks(destination string, hooks ...TopicHook) {
	tm.hooks[destination] = hooks
}

// Finds the topic for the given destination, and creates it if necessary.
func (tm *HookManager) Find(destination string) []TopicHook {
	t, ok := tm.hooks[destination]
	if !ok {
		return nil
	}
	return t
}

// LoggingTopicHook 是一个用于记录主题事件的示例 hook
type LoggingTopicHook struct {
	logger logws.Logger
}

// NewLoggingHook 创建一个新的日志记录 hook
func NewLoggingHook(logger logws.Logger) *LoggingTopicHook {
	return &LoggingTopicHook{
		logger: logger,
	}
}

func (h *LoggingTopicHook) OnTopicSubscribe(sub *client.Subscription) error {
	h.logger.Infof("Topic subscribed :%s", sub.Destination())
	return nil
}

func (h *LoggingTopicHook) OnTopicUnsubscribe(sub *client.Subscription) error {
	h.logger.Infof("Topic unsubscribed :%s", sub.Destination())
	return nil
}

func (h *LoggingTopicHook) OnTopicPublish(sub *client.Subscription, f *frame.Frame) error {
	h.logger.Infof("Topic published :%s, msg:%s", sub.Destination(), string(f.Command))
	return nil
}

type OnlineTopicHook struct {
	logger logws.Logger
	count  int64
}

func NewOnlineHook(logger logws.Logger) *OnlineTopicHook {
	return &OnlineTopicHook{
		logger: logger,
	}
}

func (h *OnlineTopicHook) OnTopicSubscribe(sub *client.Subscription) error {
	h.logger.Infof("Topic subscribed :%s", sub.Destination())
	h.count++

	f := frame.New(frame.MESSAGE)
	f.Header.Set(frame.Subscription, sub.Id())
	f.Header.Set(frame.Destination, sub.Destination())
	f.Header.Set(frame.MessageId, fmt.Sprintf("%d", time.Now().UnixMilli()))
	data := map[string]string{
		"count": fmt.Sprintf("%d", h.count),
		"msg":   fmt.Sprintf("%s is online", sub.Id()),
	}
	js, _ := json.Marshal(data)
	f.Body = js
	sub.SendTopicFrame(f)
	return nil
}

func (h *OnlineTopicHook) OnTopicUnsubscribe(sub *client.Subscription) error {
	h.logger.Infof("Topic unsubscribed :%s", sub.Destination())
	h.count--
	f := frame.New(frame.MESSAGE)
	f.Header.Set(frame.Subscription, sub.Id())
	f.Header.Set(frame.Destination, sub.Destination())
	f.Header.Set(frame.MessageId, fmt.Sprintf("%d", time.Now().UnixMilli()))
	data := map[string]string{
		"count": fmt.Sprintf("%d", h.count),
		"msg":   fmt.Sprintf("%s is offline", sub.Id()),
	}
	js, _ := json.Marshal(data)
	f.Body = js
	sub.SendTopicFrame(f)
	return nil
}

func (h *OnlineTopicHook) OnTopicPublish(sub *client.Subscription, f *frame.Frame) error {
	h.logger.Infof("Topic published :%s, msg:%s", sub.Destination(), string(f.Command))
	return nil
}
