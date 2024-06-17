package request

import (
	"time"
)

type ArticleClassifyReq struct {
	//TagId      int64 `json:"tag_id"`      // 文章标签ID
	//CategoryId int64 `json:"category_id"` // 文章分类ID
	ClassifyName string `json:"classify_name"` // 分类名
}

type ArticleTopReq struct {
	Id    int64 `json:"id"`     // 文章ID
	IsTop int64 `json:"is_top"` // 是否置顶
}

type ArticleDeleteReq struct {
	Id       int64 `json:"id"`        // 文章ID
	IsDelete int64 `json:"is_delete"` // 是否删除
}

type ArticleDetailsDTOReq struct {
	Id             int64     `json:"id"`              // 文章ID
	ArticleCover   string    `json:"article_cover"`   // 文章缩略图
	ArticleTitle   string    `json:"article_title"`   // 标题
	ArticleContent string    `json:"article_content"` // 内容
	LikeCount      int64     `json:"like_count"`      // 点赞量
	ViewsCount     int64     `json:"views_count"`     // 浏览量
	Type           int64     `json:"type"`            // 文章类型
	OriginalUrl    string    `json:"original_url"`    // 原文链接
	IsTop          int64     `json:"is_top"`          // 是否置顶
	Status         int64     `json:"status"`          // 状态值 1 公开 2 私密 3 评论可见
	CreatedAt      time.Time `json:"created_at"`      // 发表时间
	UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
	CategoryName   string    `json:"category_name"`   // 文章分类名
	TagNameList    []string  `json:"tag_name_list"`   // 文章标签列表
}
