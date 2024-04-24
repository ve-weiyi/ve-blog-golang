package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ CategoryModel = (*customCategoryModel)(nil)

type (
	// CategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoryModel.
	CategoryModel interface {
		categoryModel
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

// NewCategoryModel returns a model for the database table.
func NewCategoryModel(db *gorm.DB, cache *redis.Client) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(db, cache),
	}
}
