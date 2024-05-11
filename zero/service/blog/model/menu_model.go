package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
	}

	customMenuModel struct {
		*defaultMenuModel
	}
)

// NewMenuModel returns a model for the database table.
func NewMenuModel(db *gorm.DB, cache *redis.Client) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(db, cache),
	}
}
