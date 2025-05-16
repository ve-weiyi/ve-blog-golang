package stompws

import (
	"github.com/go-stomp/stomp/v3/frame"

	"github.com/ve-weiyi/ve-blog-golang/stompws/server/client"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/topic"
)

type TopicHook interface {
	OnTopicSubscribe(t *topic.Topic, sub *client.Subscription) error   // 订阅消息时调用
	OnTopicUnsubscribe(t *topic.Topic, sub *client.Subscription) error // 取消订阅时调用
	OnTopicPublish(t *topic.Topic, f *frame.Frame) error               // 消息发布到 Broker 时
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
