package entity

import "time"

// TableNameWebsiteConfig return the table name of <website_config>
const TableNameWebsiteConfig = "website_config"

// WebsiteConfig mapped from table <website_config>
type WebsiteConfig struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // id
	Key       string    `gorm:"column:key" json:"key" `               // 关键词
	Config    string    `gorm:"column:config" json:"config" `         // 配置信息
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName WebsiteConfig 's table name
func (*WebsiteConfig) TableName() string {
	return TableNameWebsiteConfig
}
