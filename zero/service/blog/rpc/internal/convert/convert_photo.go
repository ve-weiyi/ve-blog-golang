package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertPhotoPbToModel(in *blog.Photo) (out *model.Photo) {
	out = &model.Photo{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
		// CreatedAt: time.Unix(in.CreatedAt, 0),
		// UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertPhotoModelToPb(in *model.Photo) (out *blog.Photo) {
	out = &blog.Photo{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertPhotoAlbumPbToModel(in *blog.PhotoAlbum) (out *model.PhotoAlbum) {
	out = &model.PhotoAlbum{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		IsDelete:   in.IsDelete,
		Status:     in.Status,
		CreatedAt:  time.Unix(in.CreatedAt, 0),
		UpdatedAt:  time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertPhotoAlbumModelToPb(in *model.PhotoAlbum) (out *blog.PhotoAlbum) {
	out = &blog.PhotoAlbum{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		IsDelete:   in.IsDelete,
		Status:     in.Status,
		CreatedAt:  in.CreatedAt.Unix(),
		UpdatedAt:  in.UpdatedAt.Unix(),
	}

	return out
}
