package entity

// TableNameArticleTag return the table name of <article_tag>
const TableNameArticleTag = "article_tag"

// ArticleTag mapped from table <article_tag>
type ArticleTag struct {
	Id        int64 `gorm:"column:id" json:"id" `                 // id
	ArticleId int64 `gorm:"column:article_id" json:"article_id" ` // 文章id
	TagId     int64 `gorm:"column:tag_id" json:"tag_id" `         // 标签id
}

// TableName ArticleTag 's table name
func (*ArticleTag) TableName() string {
	return TableNameArticleTag
}
