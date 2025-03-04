package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type WebsiteService struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsiteService(svcCtx *svctx.ServiceContext) *WebsiteService {
	return &WebsiteService{
		svcCtx: svcCtx,
	}
}

// 获取博客前台首页信息
func (s *WebsiteService) GetBlogHomeInfo(reqCtx *request.Context, in *dto.GetBlogHomeInfoReq) (out *dto.GetBlogHomeInfoResp, err error) {
	// todo

	return
}

// 获取关于我的信息
func (s *WebsiteService) GetAboutMe(reqCtx *request.Context, in *dto.GetAboutMeReq) (out *dto.GetAboutMeResp, err error) {
	// todo

	return
}
