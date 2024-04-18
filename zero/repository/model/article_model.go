package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(db *gorm.DB, cache *redis.Client) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(db, cache),
	}
}
