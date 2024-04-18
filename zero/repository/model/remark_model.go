package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ RemarkModel = (*customRemarkModel)(nil)

type (
	// RemarkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRemarkModel.
	RemarkModel interface {
		remarkModel
	}

	customRemarkModel struct {
		*defaultRemarkModel
	}
)

// NewRemarkModel returns a model for the database table.
func NewRemarkModel(db *gorm.DB, cache *redis.Client) RemarkModel {
	return &customRemarkModel{
		defaultRemarkModel: newRemarkModel(db, cache),
	}
}
