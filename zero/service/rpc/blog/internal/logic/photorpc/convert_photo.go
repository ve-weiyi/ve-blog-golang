package photorpclogic

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
)

func ConvertPhotoIn(in *photorpc.PhotoNew) (out *model.Photo) {
	out = &model.Photo{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
	}

	return out
}

func ConvertPhotoOut(in *model.Photo) (out *photorpc.PhotoDetails) {
	out = &photorpc.PhotoDetails{
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

func ConvertAlbumIn(in *photorpc.AlbumNew) (out *model.PhotoAlbum) {
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

func ConvertAlbumOut(in *model.PhotoAlbum) (out *photorpc.AlbumDetails) {
	out = &photorpc.AlbumDetails{
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
