package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
)

// 分页获取PhotoAlbum记录
func (l *PhotoAlbumService) FindPhotoAlbumDetailsList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.PhotoAlbumDetailsDTO, total int64, err error) {
	albumList, total, err := l.FindPhotoAlbumList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询相册下的照片数量
	for _, in := range albumList {
		count, err := l.svcCtx.PhotoRepository.Count(reqCtx, "album_id = ?", in.Id)
		if err != nil {
			return nil, 0, err
		}

		out := &dto.PhotoAlbumDetailsDTO{
			PhotoAlbum: in,
			PhotoCount: count,
		}

		list = append(list, out)
	}

	return list, total, err
}

// 查询PhotoAlbum记录
func (l *PhotoAlbumService) FindPhotoAlbumDetails(reqCtx *request.Context, req *request.IdReq) (data *dto.PhotoAlbumDetailsDTO, err error) {
	album, err := l.svcCtx.PhotoAlbumRepository.First(reqCtx, "id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.PhotoRepository.Count(reqCtx, "album_id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	out := &dto.PhotoAlbumDetailsDTO{
		PhotoAlbum: album,
		PhotoCount: count,
	}

	return out, nil
}
