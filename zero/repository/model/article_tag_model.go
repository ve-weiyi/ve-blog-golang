package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ArticleTagModel = (*customArticleTagModel)(nil)

type (
	// ArticleTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleTagModel.
	ArticleTagModel interface {
		articleTagModel
	}

	customArticleTagModel struct {
		*defaultArticleTagModel
	}
)

// NewArticleTagModel returns a model for the database table.
func NewArticleTagModel(db *gorm.DB, cache *redis.Client) ArticleTagModel {
	return &customArticleTagModel{
		defaultArticleTagModel: newArticleTagModel(db, cache),
	}
}
