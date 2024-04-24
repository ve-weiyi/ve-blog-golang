package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserInformationModel = (*customUserInformationModel)(nil)

type (
	// UserInformationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInformationModel.
	UserInformationModel interface {
		userInformationModel
	}

	customUserInformationModel struct {
		*defaultUserInformationModel
	}
)

// NewUserInformationModel returns a model for the database table.
func NewUserInformationModel(db *gorm.DB, cache *redis.Client) UserInformationModel {
	return &customUserInformationModel{
		defaultUserInformationModel: newUserInformationModel(db, cache),
	}
}
