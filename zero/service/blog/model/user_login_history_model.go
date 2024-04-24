package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserLoginHistoryModel = (*customUserLoginHistoryModel)(nil)

type (
	// UserLoginHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLoginHistoryModel.
	UserLoginHistoryModel interface {
		userLoginHistoryModel
	}

	customUserLoginHistoryModel struct {
		*defaultUserLoginHistoryModel
	}
)

// NewUserLoginHistoryModel returns a model for the database table.
func NewUserLoginHistoryModel(db *gorm.DB, cache *redis.Client) UserLoginHistoryModel {
	return &customUserLoginHistoryModel{
		defaultUserLoginHistoryModel: newUserLoginHistoryModel(db, cache),
	}
}
