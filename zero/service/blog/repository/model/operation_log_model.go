package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ OperationLogModel = (*customOperationLogModel)(nil)

type (
	// OperationLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOperationLogModel.
	OperationLogModel interface {
		operationLogModel
	}

	customOperationLogModel struct {
		*defaultOperationLogModel
	}
)

// NewOperationLogModel returns a model for the database table.
func NewOperationLogModel(db *gorm.DB, cache *redis.Client) OperationLogModel {
	return &customOperationLogModel{
		defaultOperationLogModel: newOperationLogModel(db, cache),
	}
}
