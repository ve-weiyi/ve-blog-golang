package entity

import "time"

// TableNameChatMessage return the table name of <chat_message>
const TableNameChatMessage = "chat_message"

// ChatMessage mapped from table <chat_message>
type ChatMessage struct {
	Id         int64     `gorm:"column:id" json:"id" `                     // 主键
	ChatId     string    `gorm:"column:chat_id" json:"chat_id" `           // 聊天id
	UserId     int64     `gorm:"column:user_id" json:"user_id" `           // 用户id
	ReplyMsgId int64     `gorm:"column:reply_msg_id" json:"reply_msg_id" ` // 回复消息id
	Content    string    `gorm:"column:content" json:"content" `           // 聊天内容
	IpAddress  string    `gorm:"column:ip_address" json:"ip_address" `     // ip地址
	IpSource   string    `gorm:"column:ip_source" json:"ip_source" `       // ip来源
	Type       int64     `gorm:"column:type" json:"type" `                 // 类型
	Status     int64     `gorm:"column:status" json:"status" `             // 0正常 1撤回 2已编辑
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at" `     // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at" `     // 更新时间
}

// TableName ChatMessage 's table name
func (*ChatMessage) TableName() string {
	return TableNameChatMessage
}
