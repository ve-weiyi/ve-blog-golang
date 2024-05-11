package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ PageModel = (*customPageModel)(nil)

type (
	// PageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPageModel.
	PageModel interface {
		pageModel
	}

	customPageModel struct {
		*defaultPageModel
	}
)

// NewPageModel returns a model for the database table.
func NewPageModel(db *gorm.DB, cache *redis.Client) PageModel {
	return &customPageModel{
		defaultPageModel: newPageModel(db, cache),
	}
}
