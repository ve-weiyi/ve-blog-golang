package ws

// 发送消息的类型
type ReceiveMsg struct {
	Type      int    `json:"type"`      // 类型
	Content   string `json:"content"`   // 数据
	Timestamp int64  `json:"timestamp"` //时间戳
}

// 回复的消息
type ReplyMsg struct {
	Type      int    `json:"type"`      // 类型
	Content   string `json:"content"`   // 数据
	Timestamp int64  `json:"timestamp"` //时间戳
}

// 广播类，包括广播内容和源用户
type BroadcastMsg struct {
	From    string    `json:"from"`
	Message *ReplyMsg `json:"message"`
}

// 在线、离线事件
type OnlineEvent struct {
	Who         string `json:"who"`
	Online      bool   `json:"online"`
	OnlineCount int    `json:"online_count"`
}
