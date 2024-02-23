package request

type ChatMessage struct {
	ChatID  int    `json:"chat_id"`
	Content string `json:"content"`
}

type ChatHistory struct {
	ChatID int   `json:"chat_id"` // 聊天ID
	After  int64 `json:"after"`   // 从这个时间点之后的消息
	Before int64 `json:"before"`  // 从这个时间点之前的消息
}
