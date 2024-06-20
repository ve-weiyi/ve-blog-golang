package dto

import (
	"time"
)

type CommentQueryReq struct {
	Page     int64  `json:"page,optional"`      // 页码
	PageSize int64  `json:"page_size,optional"` // 每页数量
	TopicId  int64  `json:"topic_id,optional"`  // 主题id
	ParentId int64  `json:"parent_id,optional"` // 父评论id
	Type     int64  `json:"type,optional"`      // 评论类型 1.文章 2.友链 3.说说
	OrderBy  string `json:"order_by,optional"`  // 排序字段 create_at|like_count
}

// CommentDTO 评论
type CommentDTO struct {
	Id             int64       `json:"id"`              // 评论id
	UserId         int64       `json:"user_id"`         // 用户id
	Nickname       string      `json:"nickname"`        // 用户昵称
	Avatar         string      `json:"avatar"`          // 用户头像
	Website        string      `json:"website"`         // 个人网站
	CommentContent string      `json:"comment_content"` // 评论内容
	LikeCount      int64       `json:"like_count"`      // 点赞数
	CreatedAt      time.Time   `json:"created_at"`      // 评论时间
	ReplyCount     int64       `json:"reply_count"`     // 回复量
	ReplyDTOList   []*ReplyDTO `json:"reply_dto_list"`  // 回复列表
}

// ReplyDTO 回复
type ReplyDTO struct {
	Id             int64     `json:"id"`              // 评论id
	ParentId       int64     `json:"parent_id"`       // 父评论id
	UserId         int64     `json:"user_id"`         // 用户id
	Nickname       string    `json:"nickname"`        // 用户昵称
	Avatar         string    `json:"avatar"`          // 用户头像
	Website        string    `json:"website"`         // 个人网站
	ReplyUserId    int64     `json:"reply_user_id"`   // 被回复用户id
	ReplyNickname  string    `json:"reply_nickname"`  // 被回复用户昵称
	ReplyWebsite   string    `json:"reply_website"`   // 被回复个人网站
	CommentContent string    `json:"comment_content"` // 评论内容
	LikeCount      int64     `json:"like_count"`      // 点赞数
	CreatedAt      time.Time `json:"created_at"`      // 评论时间
}

type CommentBackDTO struct {
	Id             int64     `json:"id"`
	Avatar         string    `json:"avatar"`
	Nickname       string    `json:"nickname"`
	ReplyNickname  string    `json:"reply_nickname"`
	TopicTitle     string    `json:"topic_title"`
	CommentContent string    `json:"comment_content"`
	Type           int64     `json:"type"`
	IsReview       int64     `json:"is_review"`
	CreatedAt      time.Time `json:"created_at"`
}
