package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ChatRecordModel = (*customChatRecordModel)(nil)

type (
	// ChatRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRecordModel.
	ChatRecordModel interface {
		chatRecordModel
	}

	customChatRecordModel struct {
		*defaultChatRecordModel
	}
)

// NewChatRecordModel returns a model for the database table.
func NewChatRecordModel(db *gorm.DB, cache *redis.Client) ChatRecordModel {
	return &customChatRecordModel{
		defaultChatRecordModel: newChatRecordModel(db, cache),
	}
}
