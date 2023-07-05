package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type PhotoService struct {
	svcCtx *svc.ServiceContext
}

func NewPhotoService(svcCtx *svc.ServiceContext) *PhotoService {
	return &PhotoService{
		svcCtx: svcCtx,
	}
}

// 创建Photo记录
func (s *PhotoService) CreatePhoto(reqCtx *request.Context, photo *entity.Photo) (data *entity.Photo, err error) {
	return s.svcCtx.PhotoRepository.CreatePhoto(reqCtx, photo)
}

// 更新Photo记录
func (s *PhotoService) UpdatePhoto(reqCtx *request.Context, photo *entity.Photo) (data *entity.Photo, err error) {
	return s.svcCtx.PhotoRepository.UpdatePhoto(reqCtx, photo)
}

// 删除Photo记录
func (s *PhotoService) DeletePhoto(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.PhotoRepository.DeletePhoto(reqCtx, id)
}

// 查询Photo记录
func (s *PhotoService) FindPhoto(reqCtx *request.Context, id int) (data *entity.Photo, err error) {
	return s.svcCtx.PhotoRepository.FindPhoto(reqCtx, id)
}

// 批量删除Photo记录
func (s *PhotoService) DeletePhotoByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.PhotoRepository.DeletePhotoByIds(reqCtx, ids)
}

// 分页获取Photo记录
func (s *PhotoService) FindPhotoList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Photo, total int64, err error) {
	return s.svcCtx.PhotoRepository.FindPhotoList(reqCtx, page)
}
