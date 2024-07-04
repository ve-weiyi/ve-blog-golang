package entity

import "time"

// TableNameFriendLink return the table name of <friend_link>
const TableNameFriendLink = "friend_link"

// FriendLink mapped from table <friend_link>
type FriendLink struct {
	Id          int64     `gorm:"column:id" json:"id" `                     // id
	LinkName    string    `gorm:"column:link_name" json:"link_name" `       // 链接名
	LinkAvatar  string    `gorm:"column:link_avatar" json:"link_avatar" `   // 链接头像
	LinkAddress string    `gorm:"column:link_address" json:"link_address" ` // 链接地址
	LinkIntro   string    `gorm:"column:link_intro" json:"link_intro" `     // 链接介绍
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" `     // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" `     // 更新时间
}

// TableName FriendLink 's table name
func (*FriendLink) TableName() string {
	return TableNameFriendLink
}
