package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(db *gorm.DB, cache *redis.Client) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(db, cache),
	}
}
