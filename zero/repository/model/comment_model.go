package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(db *gorm.DB, cache *redis.Client) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(db, cache),
	}
}
