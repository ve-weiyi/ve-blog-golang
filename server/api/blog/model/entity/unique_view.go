package entity

import "time"

// TableNameUniqueView return the table name of <unique_view>
const TableNameUniqueView = "unique_view"

// UniqueView mapped from table <unique_view>
type UniqueView struct {
	Id         int64     `gorm:"column:id" json:"id" `                   // id
	ViewsCount int64     `gorm:"column:views_count" json:"views_count" ` // 访问量
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at" `   // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at" `   // 更新时间
}

// TableName UniqueView 's table name
func (*UniqueView) TableName() string {
	return TableNameUniqueView
}
