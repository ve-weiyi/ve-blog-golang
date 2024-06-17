package entity

import "time"

// TableNameTag return the table name of <tag>
const TableNameTag = "tag"

// Tag mapped from table <tag>
type Tag struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // id
	TagName   string    `gorm:"column:tag_name" json:"tag_name" `     // 标签名
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName Tag 's table name
func (*Tag) TableName() string {
	return TableNameTag
}
