package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ RoleMenuModel = (*customRoleMenuModel)(nil)

type (
	// RoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleMenuModel.
	RoleMenuModel interface {
		roleMenuModel
	}

	customRoleMenuModel struct {
		*defaultRoleMenuModel
	}
)

// NewRoleMenuModel returns a model for the database table.
func NewRoleMenuModel(db *gorm.DB, cache *redis.Client) RoleMenuModel {
	return &customRoleMenuModel{
		defaultRoleMenuModel: newRoleMenuModel(db, cache),
	}
}
