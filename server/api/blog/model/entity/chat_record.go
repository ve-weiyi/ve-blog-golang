package entity

import "time"

// TableNameChatRecord return the table name of <chat_record>
const TableNameChatRecord = "chat_record"

// ChatRecord mapped from table <chat_record>
type ChatRecord struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // 主键
	UserId    int64     `gorm:"column:user_id" json:"user_id" `       // 用户id
	Nickname  string    `gorm:"column:nickname" json:"nickname" `     // 昵称
	Avatar    string    `gorm:"column:avatar" json:"avatar" `         // 头像
	Content   string    `gorm:"column:content" json:"content" `       // 聊天内容
	IpAddress string    `gorm:"column:ip_address" json:"ip_address" ` // ip地址
	IpSource  string    `gorm:"column:ip_source" json:"ip_source" `   // ip来源
	Type      int64     `gorm:"column:type" json:"type" `             // 类型
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName ChatRecord 's table name
func (*ChatRecord) TableName() string {
	return TableNameChatRecord
}
