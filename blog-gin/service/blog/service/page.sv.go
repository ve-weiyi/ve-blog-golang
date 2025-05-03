package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
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

// 分页获取页面列表
func (s *PageService) FindPageList(reqCtx *request.Context, in *dto.PageQueryReq) (out *dto.PageResp, err error) {
	// todo

	return
}
