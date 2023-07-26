package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type PhotoAlbumService struct {
	svcCtx *svc.ServiceContext
}

func NewPhotoAlbumService(svcCtx *svc.ServiceContext) *PhotoAlbumService {
	return &PhotoAlbumService{
		svcCtx: svcCtx,
	}
}

// 创建PhotoAlbum记录
func (s *PhotoAlbumService) CreatePhotoAlbum(reqCtx *request.Context, photoAlbum *entity.PhotoAlbum) (data *entity.PhotoAlbum, err error) {
	return s.svcCtx.PhotoAlbumRepository.CreatePhotoAlbum(reqCtx, photoAlbum)
}

// 更新PhotoAlbum记录
func (s *PhotoAlbumService) UpdatePhotoAlbum(reqCtx *request.Context, photoAlbum *entity.PhotoAlbum) (data *entity.PhotoAlbum, err error) {
	return s.svcCtx.PhotoAlbumRepository.UpdatePhotoAlbum(reqCtx, photoAlbum)
}

// 删除PhotoAlbum记录
func (s *PhotoAlbumService) DeletePhotoAlbum(reqCtx *request.Context, id int) (rows int64, err error) {
	return s.svcCtx.PhotoAlbumRepository.DeletePhotoAlbum(reqCtx, id)
}

// 查询PhotoAlbum记录
func (s *PhotoAlbumService) FindPhotoAlbum(reqCtx *request.Context, id int) (data *entity.PhotoAlbum, err error) {
	return s.svcCtx.PhotoAlbumRepository.FindPhotoAlbum(reqCtx, id)
}

// 批量删除PhotoAlbum记录
func (s *PhotoAlbumService) DeletePhotoAlbumByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.PhotoAlbumRepository.DeletePhotoAlbumByIds(reqCtx, ids)
}

// 分页获取PhotoAlbum记录
func (s *PhotoAlbumService) FindPhotoAlbumList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.PhotoAlbum, total int64, err error) {
	return s.svcCtx.PhotoAlbumRepository.FindPhotoAlbumList(reqCtx, page)
}
