package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AlbumService struct {
	svcCtx *svctx.ServiceContext
}

func NewAlbumService(svcCtx *svctx.ServiceContext) *AlbumService {
	return &AlbumService{
		svcCtx: svcCtx,
	}
}

// 创建相册
func (s *AlbumService) AddAlbum(reqCtx *request.Context, in *dto.AlbumNewReq) (out *dto.AlbumBackVO, err error) {
	// todo

	return
}

// 删除相册
func (s *AlbumService) DeleteAlbum(reqCtx *request.Context, in *dto.IdReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取相册列表
func (s *AlbumService) FindAlbumList(reqCtx *request.Context, in *dto.AlbumQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 查询相册
func (s *AlbumService) GetAlbum(reqCtx *request.Context, in *dto.IdReq) (out *dto.AlbumBackVO, err error) {
	// todo

	return
}

// 更新相册
func (s *AlbumService) UpdateAlbum(reqCtx *request.Context, in *dto.AlbumNewReq) (out *dto.AlbumBackVO, err error) {
	// todo

	return
}
