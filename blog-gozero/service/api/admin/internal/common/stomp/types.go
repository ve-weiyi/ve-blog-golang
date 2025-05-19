package stomp

type MessageTypeEnum int

const (
	OnlineCount MessageTypeEnum = iota + 1 // 在线人数
)

type (
	MessageEvent struct {
		Type MessageTypeEnum `json:"type"` // 消息类型
		Data string          `json:"data"` // 消息内容
	}

	OnlineEvent struct {
		Count    int64 `json:"count"`
		IsOnline bool  `json:"is_online"`
	}
)
