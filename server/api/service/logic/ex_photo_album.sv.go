package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 分页获取PhotoAlbum记录
func (s *PhotoAlbumService) FindPhotoAlbumDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.PhotoAlbumDetails, total int64, err error) {
	albumList, total, err := s.svcCtx.PhotoAlbumRepository.FindPhotoAlbumList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询相册下的照片数量
	for _, in := range albumList {
		query := &request.PageQuery{Conditions: []*request.Condition{{Field: "album_id", Rule: "=", Value: in.ID}}}
		_, count, err := s.svcCtx.PhotoRepository.FindPhotoList(reqCtx, query)
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
	album, err := s.svcCtx.PhotoAlbumRepository.FindPhotoAlbum(reqCtx, id)
	if err != nil {
		return nil, err
	}

	query := &request.PageQuery{Conditions: []*request.Condition{{Field: "album_id", Rule: "=", Value: id}}}
	_, count, err := s.svcCtx.PhotoRepository.FindPhotoList(reqCtx, query)
	if err != nil {
		return nil, err
	}

	out := &response.PhotoAlbumDetails{
		PhotoAlbum: album,
		PhotoCount: count,
	}

	return out, nil
}
