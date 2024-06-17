package dto

import "github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"

type PhotoAlbumDetailsDTO struct {
	*entity.PhotoAlbum
	PhotoCount int64 `json:"photo_count"`
}
