package entity

import "time"

// TableNameRemark return the table name of <remark>
const TableNameRemark = "remark"

// Remark mapped from table <remark>
type Remark struct {
	Id             int64     `gorm:"column:id" json:"id" `                           // 主键id
	Nickname       string    `gorm:"column:nickname" json:"nickname" `               // 昵称
	Avatar         string    `gorm:"column:avatar" json:"avatar" `                   // 头像
	MessageContent string    `gorm:"column:message_content" json:"message_content" ` // 留言内容
	IpAddress      string    `gorm:"column:ip_address" json:"ip_address" `           // 用户ip
	IpSource       string    `gorm:"column:ip_source" json:"ip_source" `             // 用户地址
	Time           int64     `gorm:"column:time" json:"time" `                       // 弹幕速度
	IsReview       int64     `gorm:"column:is_review" json:"is_review" `             // 是否审核
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at" `           // 发布时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at" `           // 更新时间
}

// TableName Remark 's table name
func (*Remark) TableName() string {
	return TableNameRemark
}
