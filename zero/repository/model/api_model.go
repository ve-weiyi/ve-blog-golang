package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ApiModel = (*customApiModel)(nil)

type (
	// ApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApiModel.
	ApiModel interface {
		apiModel
	}

	customApiModel struct {
		*defaultApiModel
	}
)

// NewApiModel returns a model for the database table.
func NewApiModel(db *gorm.DB, cache *redis.Client) ApiModel {
	return &customApiModel{
		defaultApiModel: newApiModel(db, cache),
	}
}
