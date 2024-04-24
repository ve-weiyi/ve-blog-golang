package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ TagModel = (*customTagModel)(nil)

type (
	// TagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagModel.
	TagModel interface {
		tagModel
	}

	customTagModel struct {
		*defaultTagModel
	}
)

// NewTagModel returns a model for the database table.
func NewTagModel(db *gorm.DB, cache *redis.Client) TagModel {
	return &customTagModel{
		defaultTagModel: newTagModel(db, cache),
	}
}
