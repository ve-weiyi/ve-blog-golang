package mediaservicelogic

import (
	"context"

	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/model"
)

func findPhotoCountGroupAlbum(ctx context.Context, db *gorm.DB, list []*model.TAlbum) map[int64]int {
	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	if len(ids) == 0 {
		return nil
	}

	var results []struct {
		AlbumID    int64 `gorm:"column:album_id"`
		PhotoCount int   `gorm:"column:photo_count"`
	}
	err := db.WithContext(ctx).Model(&model.TPhoto{}).
		Select("album_id, COUNT(*) as photo_count").
		Where("album_id IN ?", ids).
		Group("album_id").
		Scan(&results).Error
	if err != nil {
		return nil
	}

	cm := make(map[int64]int)
	for _, r := range results {
		cm[r.AlbumID] = r.PhotoCount
	}
	return cm
}

func convertAlbumOut(in *model.TAlbum, photoCountMap map[int64]int) *mediarpc.Album {
	var count int
	if photoCountMap != nil {
		count = photoCountMap[in.Id]
	}
	return &mediarpc.Album{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		IsDelete:   in.IsDelete,
		Status:     in.Status,
		CreatedAt:  in.CreatedAt.UnixMilli(),
		UpdatedAt:  in.UpdatedAt.UnixMilli(),
		PhotoCount: int64(count),
	}
}

func convertPhotoOut(in *model.TPhoto) *mediarpc.Photo {
	return &mediarpc.Photo{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
		CreatedAt: in.CreatedAt.UnixMilli(),
		UpdatedAt: in.UpdatedAt.UnixMilli(),
	}
}
