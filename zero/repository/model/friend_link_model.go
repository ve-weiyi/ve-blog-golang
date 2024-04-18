package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ FriendLinkModel = (*customFriendLinkModel)(nil)

type (
	// FriendLinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendLinkModel.
	FriendLinkModel interface {
		friendLinkModel
	}

	customFriendLinkModel struct {
		*defaultFriendLinkModel
	}
)

// NewFriendLinkModel returns a model for the database table.
func NewFriendLinkModel(db *gorm.DB, cache *redis.Client) FriendLinkModel {
	return &customFriendLinkModel{
		defaultFriendLinkModel: newFriendLinkModel(db, cache),
	}
}
