package entity

import "time"

// TableNamePage return the table name of <page>
const TableNamePage = "page"

// Page mapped from table <page>
type Page struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // 页面id
	PageName  string    `gorm:"column:page_name" json:"page_name" `   // 页面名
	PageLabel string    `gorm:"column:page_label" json:"page_label" ` // 页面标签
	PageCover string    `gorm:"column:page_cover" json:"page_cover" ` // 页面封面
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName Page 's table name
func (*Page) TableName() string {
	return TableNamePage
}
