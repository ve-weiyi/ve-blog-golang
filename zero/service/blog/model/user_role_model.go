package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserRoleModel = (*customUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(db *gorm.DB, cache *redis.Client) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(db, cache),
	}
}
