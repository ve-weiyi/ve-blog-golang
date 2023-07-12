package response

import "time"

// CommentDTO 评论
type CommentDTO struct {
	ID             int         `json:"id"`             // 评论id
	UserID         int         `json:"userId"`         // 用户id
	Nickname       string      `json:"nickname"`       // 用户昵称
	Avatar         string      `json:"avatar"`         // 用户头像
	WebSite        string      `json:"webSite"`        // 个人网站
	CommentContent string      `json:"commentContent"` // 评论内容
	LikeCount      int         `json:"likeCount"`      // 点赞数
	CreatedAt      time.Time   `json:"createdAt"`      // 评论时间
	ReplyCount     int64       `json:"replyCount"`     // 回复量
	ReplyDTOList   []*ReplyDTO `json:"replyDTOList"`   // 回复列表
}

// ReplyDTO 回复
type ReplyDTO struct {
	ID             int       `json:"id"`             // 评论id
	ParentID       int       `json:"parentId"`       // 父评论id
	UserID         int       `json:"userId"`         // 用户id
	Nickname       string    `json:"nickname"`       // 用户昵称
	Avatar         string    `json:"avatar"`         // 用户头像
	WebSite        string    `json:"webSite"`        // 个人网站
	ReplyUserID    int       `json:"replyUserId"`    // 被回复用户id
	ReplyNickname  string    `json:"replyNickname"`  // 被回复用户昵称
	ReplyWebSite   string    `json:"replyWebSite"`   // 被回复个人网站
	CommentContent string    `json:"commentContent"` // 评论内容
	LikeCount      int       `json:"likeCount"`      // 点赞数
	CreatedAt      time.Time `json:"createdAt"`      // 评论时间
}

type CommentBackDTO struct {
	ID             int       `json:"id"`
	Avatar         string    `json:"avatar"`
	Nickname       string    `json:"nickname"`
	ReplyNickname  string    `json:"replyNickname"`
	ArticleTitle   string    `json:"articleTitle"`
	CommentContent string    `json:"commentContent"`
	Type           int       `json:"type"`
	IsReview       bool      `json:"isReview"`
	CreatedAt      time.Time `json:"createdAt"`
}
