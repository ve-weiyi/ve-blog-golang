package dto

import "time"

type CategoryDetailsDTO struct {
	Id           int64     `json:"id"`
	CategoryName string    `json:"category_name"` // 分类名
	ArticleCount int64     `json:"article_count"`
	CreatedAt    time.Time `json:"created_at"` // 创建时间
	UpdatedAt    time.Time `json:"updated_at"` // 更新时间
}

type TagDetailsDTO struct {
	Id           int64     `json:"id"`            // 标签ID
	TagName      string    `json:"tag_name"`      // 标签名
	ArticleCount int64     `json:"article_count"` // 文章数量
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
}

// 分类
type CategoryDTO struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"` // 分类名
}

// 标签
type TagDTO struct {
	Id      int64  `json:"id"`       // 标签ID
	TagName string `json:"tag_name"` // 标签名
}
