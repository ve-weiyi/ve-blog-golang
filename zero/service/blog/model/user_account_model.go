package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserAccountModel = (*customUserAccountModel)(nil)

type (
	// UserAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAccountModel.
	UserAccountModel interface {
		userAccountModel
	}

	customUserAccountModel struct {
		*defaultUserAccountModel
	}
)

// NewUserAccountModel returns a model for the database table.
func NewUserAccountModel(db *gorm.DB, cache *redis.Client) UserAccountModel {
	return &customUserAccountModel{
		defaultUserAccountModel: newUserAccountModel(db, cache),
	}
}
