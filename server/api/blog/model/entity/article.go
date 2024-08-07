package entity

import "time"

// TableNameArticle return the table name of <article>
const TableNameArticle = "article"

// Article mapped from table <article>
type Article struct {
	Id             int64     `gorm:"column:id" json:"id" `                           // id
	UserId         int64     `gorm:"column:user_id" json:"user_id" `                 // 作者
	CategoryId     int64     `gorm:"column:category_id" json:"category_id" `         // 文章分类
	ArticleCover   string    `gorm:"column:article_cover" json:"article_cover" `     // 文章缩略图
	ArticleTitle   string    `gorm:"column:article_title" json:"article_title" `     // 标题
	ArticleContent string    `gorm:"column:article_content" json:"article_content" ` // 内容
	Type           int64     `gorm:"column:type" json:"type" `                       // 文章类型 1原创 2转载 3翻译
	OriginalUrl    string    `gorm:"column:original_url" json:"original_url" `       // 原文链接
	IsTop          int64     `gorm:"column:is_top" json:"is_top" `                   // 是否置顶 0否 1是
	IsDelete       int64     `gorm:"column:is_delete" json:"is_delete" `             // 是否删除  0否 1是
	Status         int64     `gorm:"column:status" json:"status" `                   // 状态值 1 公开 2 私密 3 草稿 4 已删除
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at" `           // 发表时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at" `           // 更新时间
}

// TableName Article 's table name
func (*Article) TableName() string {
	return TableNameArticle
}
