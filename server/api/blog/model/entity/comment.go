package entity

import "time"

// TableNameComment return the table name of <comment>
const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	Id             int64     `gorm:"column:id" json:"id" `                           // 主键
	TopicId        int64     `gorm:"column:topic_id" json:"topic_id" `               // 主题id
	ParentId       int64     `gorm:"column:parent_id" json:"parent_id" `             // 父评论id
	SessionId      int64     `gorm:"column:session_id" json:"session_id" `           // 会话id
	UserId         int64     `gorm:"column:user_id" json:"user_id" `                 // 评论用户id
	ReplyUserId    int64     `gorm:"column:reply_user_id" json:"reply_user_id" `     // 评论回复用户id
	CommentContent string    `gorm:"column:comment_content" json:"comment_content" ` // 评论内容
	LikeCount      int64     `gorm:"column:like_count" json:"like_count" `           // 评论点赞数量
	Type           int64     `gorm:"column:type" json:"type" `                       // 评论类型 1.文章 2.友链 3.说说
	Status         int64     `gorm:"column:status" json:"status" `                   // 状态 0.正常 1.已编辑 2.已删除
	IsReview       int64     `gorm:"column:is_review" json:"is_review" `             // 是否审核
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at" `           // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at" `           // 更新时间
}

// TableName Comment 's table name
func (*Comment) TableName() string {
	return TableNameComment
}
