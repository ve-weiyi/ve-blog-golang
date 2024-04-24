package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ WebsiteConfigModel = (*customWebsiteConfigModel)(nil)

type (
	// WebsiteConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWebsiteConfigModel.
	WebsiteConfigModel interface {
		websiteConfigModel
	}

	customWebsiteConfigModel struct {
		*defaultWebsiteConfigModel
	}
)

// NewWebsiteConfigModel returns a model for the database table.
func NewWebsiteConfigModel(db *gorm.DB, cache *redis.Client) WebsiteConfigModel {
	return &customWebsiteConfigModel{
		defaultWebsiteConfigModel: newWebsiteConfigModel(db, cache),
	}
}
