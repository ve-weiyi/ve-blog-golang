package response

import "time"

type ArticleDetails struct {
	ID             int       `json:"id"`              // 文章ID
	ArticleCover   string    `json:"article_cover"`   // 文章缩略图
	ArticleTitle   string    `json:"article_title"`   // 标题
	ArticleContent string    `json:"article_content"` // 内容
	LikeCount      int       `json:"like_count"`      // 点赞量
	ViewsCount     int       `json:"views_count"`     // 浏览量
	Type           int       `json:"type"`            // 文章类型
	OriginalURL    string    `json:"original_url"`    // 原文链接
	IsTop          int       `json:"is_top"`          // 是否置顶
	IsDelete       int       `json:"is_delete"`       // 是否删除
	Status         int       `json:"status"`          // 状态值 1 公开 2 私密 3 评论可见
	CreatedAt      time.Time `json:"created_at"`      // 发表时间
	UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
	CategoryName   string    `json:"category_name"`   // 文章分类名
	TagNameList    []string  `json:"tag_name_list"`   // 文章标签列表
}

type ArticleConditionDTO struct {
	ArticleDTOList []*ArticleDetails `json:"article_dto_list"` // 文章列表
	ConditionName  string            `json:"condition_name"`   // 条件名
}

// ArticleRecommendDetails represents an article
type ArticleRecommendDetails struct {
	ArticleDetails
	LastArticle          *ArticlePreviewDTO   `json:"last_article"`           // 上一篇文章
	NextArticle          *ArticlePreviewDTO   `json:"next_article"`           // 下一篇文章
	RecommendArticleList []*ArticlePreviewDTO `json:"recommend_article_list"` // 推荐文章列表
	NewestArticleList    []*ArticlePreviewDTO `json:"newest_article_list"`    // 最新文章列表
}

// TagDTO 标签
type TagDTO struct {
	ID      int    `json:"id"`       // 标签ID
	TagName string `json:"tag_name"` // 标签名
}

// CategoryDTO 分类
type CategoryDTO struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"` // 分类名
	ArticleCount int64  `json:"article_count"`
}

// ArticleArchivesDTO 文章预览
type ArticlePreviewDTO struct {
	ID           int       `json:"id"`            // 文章ID
	ArticleCover string    `json:"article_cover"` // 文章缩略图
	ArticleTitle string    `json:"article_title"` // 标题
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
}
