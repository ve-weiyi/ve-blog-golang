package response

import "time"

type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"` // 分类名
	ArticleCount int64     `json:"article_count"`
	CreatedAt    time.Time `json:"created_at"` // 创建时间
	UpdatedAt    time.Time `json:"updated_at"` // 更新时间
}
