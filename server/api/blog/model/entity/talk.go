package entity

import "time"

// TableNameTalk return the table name of <talk>
const TableNameTalk = "talk"

// Talk mapped from table <talk>
type Talk struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // 说说id
	UserId    int64     `gorm:"column:user_id" json:"user_id" `       // 用户id
	Content   string    `gorm:"column:content" json:"content" `       // 说说内容
	Images    string    `gorm:"column:images" json:"images" `         // 图片
	IsTop     int64     `gorm:"column:is_top" json:"is_top" `         // 是否置顶
	Status    int64     `gorm:"column:status" json:"status" `         // 状态 1.公开 2.私密
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName Talk 's table name
func (*Talk) TableName() string {
	return TableNameTalk
}
