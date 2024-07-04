package entity

import "time"

// TableNameCategory return the table name of <category>
const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	Id           int64     `gorm:"column:id" json:"id" `                       // id
	CategoryName string    `gorm:"column:category_name" json:"category_name" ` // 分类名
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at" `       // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at" `       // 更新时间
}

// TableName Category 's table name
func (*Category) TableName() string {
	return TableNameCategory
}
