package request

type ArticleCondition struct {
	TagID      int `json:"tag_id"`      // 文章标签ID
	CategoryID int `json:"category_id"` // 文章分类ID
}
