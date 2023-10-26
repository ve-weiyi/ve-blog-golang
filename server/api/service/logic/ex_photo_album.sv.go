package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
)

// 分页获取PhotoAlbum记录
func (s *PhotoAlbumService) FindPhotoAlbumDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.PhotoAlbumDetails, total int64, err error) {
	albumList, total, err := s.FindPhotoAlbumList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询相册下的照片数量
	for _, in := range albumList {
		count, err := s.svcCtx.PhotoRepository.Count(reqCtx, &sqlx.Condition{Field: "album_id", Rule: "=", Value: in.ID})
		if err != nil {
			return nil, 0, err
		}

		out := &response.PhotoAlbumDetails{
			PhotoAlbum: in,
			PhotoCount: count,
		}

		list = append(list, out)
	}

	return list, total, err
}

// 查询PhotoAlbum记录
func (s *PhotoAlbumService) FindPhotoAlbumDetails(reqCtx *request.Context, id int) (data *response.PhotoAlbumDetails, err error) {
	album, err := s.svcCtx.PhotoAlbumRepository.FindPhotoAlbumById(reqCtx, id)
	if err != nil {
		return nil, err
	}

	count, err := s.svcCtx.PhotoRepository.Count(reqCtx, &sqlx.Condition{Field: "album_id", Rule: "=", Value: id})
	if err != nil {
		return nil, err
	}

	out := &response.PhotoAlbumDetails{
		PhotoAlbum: album,
		PhotoCount: count,
	}

	return out, nil
}
