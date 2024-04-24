package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserOauthModel = (*customUserOauthModel)(nil)

type (
	// UserOauthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserOauthModel.
	UserOauthModel interface {
		userOauthModel
	}

	customUserOauthModel struct {
		*defaultUserOauthModel
	}
)

// NewUserOauthModel returns a model for the database table.
func NewUserOauthModel(db *gorm.DB, cache *redis.Client) UserOauthModel {
	return &customUserOauthModel{
		defaultUserOauthModel: newUserOauthModel(db, cache),
	}
}
