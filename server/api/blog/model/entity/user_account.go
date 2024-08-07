package entity

import "time"

// TableNameUserAccount return the table name of <user_account>
const TableNameUserAccount = "user_account"

// UserAccount mapped from table <user_account>
type UserAccount struct {
	Id           int64     `gorm:"column:id" json:"id" `                       // id
	Username     string    `gorm:"column:username" json:"username" `           // 用户名
	Password     string    `gorm:"column:password" json:"password" `           // 密码
	Status       int64     `gorm:"column:status" json:"status" `               // 状态: -1删除 0正常 1禁用
	RegisterType string    `gorm:"column:register_type" json:"register_type" ` // 注册方式
	IpAddress    string    `gorm:"column:ip_address" json:"ip_address" `       // 注册ip
	IpSource     string    `gorm:"column:ip_source" json:"ip_source" `         // 注册ip 源
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at" `       // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at" `       // 更新时间
}

// TableName UserAccount 's table name
func (*UserAccount) TableName() string {
	return TableNameUserAccount
}
