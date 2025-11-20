package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type PhotoService struct {
	svcCtx *svctx.ServiceContext
}

func NewPhotoService(svcCtx *svctx.ServiceContext) *PhotoService {
	return &PhotoService{
		svcCtx: svcCtx,
	}
}

// 创建照片
func (s *PhotoService) AddPhoto(reqCtx *request.Context, in *dto.PhotoNewReq) (out *dto.PhotoBackVO, err error) {
	// todo

	return
}

// 删除照片
func (s *PhotoService) DeletesPhoto(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取照片列表
func (s *PhotoService) FindPhotoList(reqCtx *request.Context, in *dto.PhotoQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 预删除照片
func (s *PhotoService) PreDeletePhoto(reqCtx *request.Context, in *dto.PreDeletePhotoReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 更新照片
func (s *PhotoService) UpdatePhoto(reqCtx *request.Context, in *dto.PhotoNewReq) (out *dto.PhotoBackVO, err error) {
	// todo

	return
}
