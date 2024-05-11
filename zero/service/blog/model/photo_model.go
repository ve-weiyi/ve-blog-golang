package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ PhotoModel = (*customPhotoModel)(nil)

type (
	// PhotoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhotoModel.
	PhotoModel interface {
		photoModel
	}

	customPhotoModel struct {
		*defaultPhotoModel
	}
)

// NewPhotoModel returns a model for the database table.
func NewPhotoModel(db *gorm.DB, cache *redis.Client) PhotoModel {
	return &customPhotoModel{
		defaultPhotoModel: newPhotoModel(db, cache),
	}
}
