package entity

import "time"

// TableNameUserLoginHistory return the table name of <user_login_history>
const TableNameUserLoginHistory = "user_login_history"

// UserLoginHistory mapped from table <user_login_history>
type UserLoginHistory struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // id
	UserId    int64     `gorm:"column:user_id" json:"user_id" `       // 用户id
	LoginType string    `gorm:"column:login_type" json:"login_type" ` // 登录类型
	Agent     string    `gorm:"column:agent" json:"agent" `           // 代理
	IpAddress string    `gorm:"column:ip_address" json:"ip_address" ` // ip host
	IpSource  string    `gorm:"column:ip_source" json:"ip_source" `   // ip 源
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName UserLoginHistory 's table name
func (*UserLoginHistory) TableName() string {
	return TableNameUserLoginHistory
}
