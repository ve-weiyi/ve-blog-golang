package entity

import "time"

// TableNameUserOauth return the table name of <user_oauth>
const TableNameUserOauth = "user_oauth"

// UserOauth mapped from table <user_oauth>
type UserOauth struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // id
	UserId    int64     `gorm:"column:user_id" json:"user_id" `       // 用户id
	OpenId    string    `gorm:"column:open_id" json:"open_id" `       // 开发平台id，标识唯一用户
	Platform  string    `gorm:"column:platform" json:"platform" `     // 平台:手机号、邮箱、微信、飞书
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName UserOauth 's table name
func (*UserOauth) TableName() string {
	return TableNameUserOauth
}
