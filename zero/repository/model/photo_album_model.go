package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ PhotoAlbumModel = (*customPhotoAlbumModel)(nil)

type (
	// PhotoAlbumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhotoAlbumModel.
	PhotoAlbumModel interface {
		photoAlbumModel
	}

	customPhotoAlbumModel struct {
		*defaultPhotoAlbumModel
	}
)

// NewPhotoAlbumModel returns a model for the database table.
func NewPhotoAlbumModel(db *gorm.DB, cache *redis.Client) PhotoAlbumModel {
	return &customPhotoAlbumModel{
		defaultPhotoAlbumModel: newPhotoAlbumModel(db, cache),
	}
}
