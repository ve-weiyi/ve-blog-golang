package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UniqueViewModel = (*customUniqueViewModel)(nil)

type (
	// UniqueViewModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUniqueViewModel.
	UniqueViewModel interface {
		uniqueViewModel
	}

	customUniqueViewModel struct {
		*defaultUniqueViewModel
	}
)

// NewUniqueViewModel returns a model for the database table.
func NewUniqueViewModel(db *gorm.DB, cache *redis.Client) UniqueViewModel {
	return &customUniqueViewModel{
		defaultUniqueViewModel: newUniqueViewModel(db, cache),
	}
}
