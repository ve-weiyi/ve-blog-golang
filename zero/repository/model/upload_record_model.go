package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UploadRecordModel = (*customUploadRecordModel)(nil)

type (
	// UploadRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUploadRecordModel.
	UploadRecordModel interface {
		uploadRecordModel
	}

	customUploadRecordModel struct {
		*defaultUploadRecordModel
	}
)

// NewUploadRecordModel returns a model for the database table.
func NewUploadRecordModel(db *gorm.DB, cache *redis.Client) UploadRecordModel {
	return &customUploadRecordModel{
		defaultUploadRecordModel: newUploadRecordModel(db, cache),
	}
}
