package response

import "time"

type TalkDetailsDTO struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"user_id"`       // 用户ID
	Nickname     string    `json:"nickname"`      // 用户昵称
	Avatar       string    `json:"avatar"`        // 用户头像
	Content      string    `json:"content"`       // 评论内容
	ImgList      []string  `json:"img_list"`      // 图片Url列表
	IsTop        int64     `json:"is_top"`        // 是否置顶
	Status       int64     `json:"status"`        // 状态
	LikeCount    int64     `json:"like_count"`    // 点赞量
	CommentCount int64     `json:"comment_count"` // 评论量
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
}
