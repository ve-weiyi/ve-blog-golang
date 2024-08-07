package entity

import "time"

// TableNameUserInformation return the table name of <user_information>
const TableNameUserInformation = "user_information"

// UserInformation mapped from table <user_information>
type UserInformation struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // id
	UserId    int64     `gorm:"column:user_id" json:"user_id" `       // 用户id
	Email     string    `gorm:"column:email" json:"email" `           // 用户邮箱
	Nickname  string    `gorm:"column:nickname" json:"nickname" `     // 用户昵称
	Avatar    string    `gorm:"column:avatar" json:"avatar" `         // 用户头像
	Phone     string    `gorm:"column:phone" json:"phone" `           // 用户手机号
	Intro     string    `gorm:"column:intro" json:"intro" `           // 个人简介
	Website   string    `gorm:"column:website" json:"website" `       // 个人网站
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName UserInformation 's table name
func (*UserInformation) TableName() string {
	return TableNameUserInformation
}
