package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type PageService struct {
	svcCtx *svctx.ServiceContext
}

func NewPageService(svcCtx *svctx.ServiceContext) *PageService {
	return &PageService{
		svcCtx: svcCtx,
	}
}

// 创建页面
func (s *PageService) AddPage(reqCtx *request.Context, in *dto.PageNewReq) (out *dto.PageBackVO, err error) {
	// todo

	return
}

// 删除页面
func (s *PageService) DeletePage(reqCtx *request.Context, in *dto.IdReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取页面列表
func (s *PageService) FindPageList(reqCtx *request.Context, in *dto.PageQueryReq) (out *dto.PageResp, err error) {
	// todo

	return
}

// 更新页面
func (s *PageService) UpdatePage(reqCtx *request.Context, in *dto.PageNewReq) (out *dto.PageBackVO, err error) {
	// todo

	return
}
