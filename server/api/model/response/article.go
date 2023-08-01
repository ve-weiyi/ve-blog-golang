package response

import "time"

type ArticleDTO struct {
	ID             int       `json:"id"`               // 文章ID
	ArticleCover   string    `json:"article_cover"`    // 文章缩略图
	ArticleTitle   string    `json:"article_title"`    // 标题
	ArticleContent string    `json:"article_content"`  // 内容
	LikeCount      int       `json:"like_count"`       // 点赞量
	ViewsCount     int       `json:"views_count"`      // 浏览量
	Type           int       `json:"type"`             // 文章类型
	OriginalURL    string    `json:"original_url"`     // 原文链接
	CreatedAt      time.Time `json:"created_at"`       // 发表时间
	UpdatedAt      time.Time `json:"updated_at"`       // 更新时间
	CategoryID     int       `json:"category_id"`      // 文章分类ID
	CategoryName   string    `json:"category_name"`    // 文章分类名
	ArticleTagList []*TagDTO `json:"article_tag_list"` // 文章标签列表
}

type ArticleConditionDTO struct {
	ArticleDTOList []*ArticleDTO `json:"article_dto_list"` // 文章列表
	ConditionName  string        `json:"condition_name"`   // 条件名
}

// ArticleDetails represents an article
type ArticleDetails struct {
	ID                   int                    `json:"id"`                     // 文章ID
	ArticleCover         string                 `json:"article_cover"`          // 文章缩略图
	ArticleTitle         string                 `json:"article_title"`          // 标题
	ArticleContent       string                 `json:"article_content"`        // 内容
	LikeCount            int                    `json:"like_count"`             // 点赞量
	ViewsCount           int                    `json:"views_count"`            // 浏览量
	Type                 int                    `json:"type"`                   // 文章类型
	OriginalURL          string                 `json:"original_url"`           // 原文链接
	CreatedAt            time.Time              `json:"created_at"`             // 发表时间
	UpdatedAt            time.Time              `json:"updated_at"`             // 更新时间
	CategoryID           int                    `json:"category_id"`            // 文章分类ID
	CategoryName         string                 `json:"category_name"`          // 文章分类名
	ArticleTagList       []*TagDTO              `json:"article_tag_list"`       // 文章标签列表
	LastArticle          *ArticlePaginationDTO  `json:"last_article"`           // 上一篇文章
	NextArticle          *ArticlePaginationDTO  `json:"next_article"`           // 下一篇文章
	RecommendArticleList []*ArticleRecommendDTO `json:"recommend_article_list"` // 推荐文章列表
	NewestArticleList    []*ArticleRecommendDTO `json:"newest_article_list"`    // 最新文章列表
}

// TagDTO represents a tag
type TagDTO struct {
	// tag fields...
	ID      int    `json:"id"`       // 标签ID
	TagName string `json:"tag_name"` // 标签名
}

// ArticlePaginationDTO represents pagination information for an article
type ArticlePaginationDTO struct {
	// pagination fields...
	ID           int    `json:"id"`            // 文章ID
	ArticleCover string `json:"article_cover"` // 文章缩略图
	ArticleTitle string `json:"article_title"` // 标题
}

// ArticleRecommendDTO represents a recommended article
type ArticleRecommendDTO struct {
	// recommended article fields...
	ID           int       `json:"id"`            // 文章ID
	ArticleCover string    `json:"article_cover"` // 文章缩略图
	ArticleTitle string    `json:"article_title"` // 标题
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
}

// ArticleArchivesDTO represents a recommended article
type ArticleArchivesDTO struct {
	ID           int       `json:"id"`            // 文章ID
	ArticleCover string    `json:"article_cover"` // 文章缩略图
	ArticleTitle string    `json:"article_title"` // 标题
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
}
