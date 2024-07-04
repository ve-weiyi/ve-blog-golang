package entity

import "time"

// TableNameChatSession return the table name of <chat_session>
const TableNameChatSession = "chat_session"

// ChatSession mapped from table <chat_session>
type ChatSession struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // 主键
	ChatId    string    `gorm:"column:chat_id" json:"chat_id" `       // 聊天id
	ChatTitle string    `gorm:"column:chat_title" json:"chat_title" ` // 标题
	Type      string    `gorm:"column:type" json:"type" `             // 类型
	Status    int64     `gorm:"column:status" json:"status" `         // 0正常 1删除
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName ChatSession 's table name
func (*ChatSession) TableName() string {
	return TableNameChatSession
}
