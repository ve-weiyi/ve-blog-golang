package dto

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
	Status         int64     `json:"status"`          // 状态值 1 公开 2 私密 3 草稿 4 已删除
	CreatedAt      time.Time `json:"created_at"`      // 发表时间
	UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
	CategoryName   string    `json:"category_name"`   // 文章分类名
	TagNameList    []string  `json:"tag_name_list"`   // 文章标签列表
}

// 首页文章列表
type ArticleHome struct {
	ArticleDTO
	ArticleCategory *CategoryDTO `json:"article_category"` // 文章分类
	ArticleTagList  []*TagDTO    `json:"article_tag_list"` // 文章标签列表
}

// 后台文章列表
type ArticleBack struct {
	ArticleDTO
	CategoryName string   `json:"category_name"` // 文章分类名
	TagNameList  []string `json:"tag_name_list"` // 文章标签列表
}

type ArticleDTO struct {
	Id             int64     `json:"id"`              // 文章ID
	ArticleCover   string    `json:"article_cover"`   // 文章缩略图
	ArticleTitle   string    `json:"article_title"`   // 标题
	ArticleContent string    `json:"article_content"` // 内容
	LikeCount      int64     `json:"like_count"`      // 点赞量
	ViewsCount     int64     `json:"views_count"`     // 浏览量
	Type           int64     `json:"type"`            // 文章类型
	OriginalUrl    string    `json:"original_url"`    // 原文链接
	IsTop          int64     `json:"is_top"`          // 是否置顶
	IsDelete       int64     `json:"is_delete"`       // 是否删除
	Status         int64     `json:"status"`          // 状态值 1 公开 2 私密 3 草稿 4 已删除
	CreatedAt      time.Time `json:"created_at"`      // 发表时间
	UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
}

// 文章系别列表
type ArticleClassifyResp struct {
	ArticleList   []*ArticleHome `json:"article_list"`   // 文章列表
	ConditionName string         `json:"condition_name"` // 条件名
}

// 文章详情，包含文章内容，上一篇文章，下一篇文章，推荐文章列表，最新文章列表
type ArticlePageDetailsDTO struct {
	ArticleHome
	LastArticle          *ArticlePreviewDTO   `json:"last_article"`           // 上一篇文章
	NextArticle          *ArticlePreviewDTO   `json:"next_article"`           // 下一篇文章
	RecommendArticleList []*ArticlePreviewDTO `json:"recommend_article_list"` // 推荐文章列表
	NewestArticleList    []*ArticlePreviewDTO `json:"newest_article_list"`    // 最新文章列表
}

// 文章预览
type ArticlePreviewDTO struct {
	Id           int64     `json:"id"`            // 文章ID
	ArticleCover string    `json:"article_cover"` // 文章缩略图
	ArticleTitle string    `json:"article_title"` // 标题
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
}
