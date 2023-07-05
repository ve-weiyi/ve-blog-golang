package response

import "time"

type TalkDetails struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`    // 用户ID
	Nickname  string    `json:"nickname"`   // 用户昵称
	Avatar    string    `json:"avatar"`     // 用户头像
	Content   string    `json:"content"`    // 评论内容
	Images    string    `json:"images"`     // 评论图片
	IsTop     int       `json:"is_top"`     // 是否置顶
	Status    int       `json:"status"`     // 状态
	LikeCount int       `json:"like_count"` // 点赞量
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
