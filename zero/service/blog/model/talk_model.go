package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ TalkModel = (*customTalkModel)(nil)

type (
	// TalkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTalkModel.
	TalkModel interface {
		talkModel
	}

	customTalkModel struct {
		*defaultTalkModel
	}
)

// NewTalkModel returns a model for the database table.
func NewTalkModel(db *gorm.DB, cache *redis.Client) TalkModel {
	return &customTalkModel{
		defaultTalkModel: newTalkModel(db, cache),
	}
}
