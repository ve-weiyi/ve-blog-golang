package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ RoleApiModel = (*customRoleApiModel)(nil)

type (
	// RoleApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleApiModel.
	RoleApiModel interface {
		roleApiModel
	}

	customRoleApiModel struct {
		*defaultRoleApiModel
	}
)

// NewRoleApiModel returns a model for the database table.
func NewRoleApiModel(db *gorm.DB, cache *redis.Client) RoleApiModel {
	return &customRoleApiModel{
		defaultRoleApiModel: newRoleApiModel(db, cache),
	}
}
