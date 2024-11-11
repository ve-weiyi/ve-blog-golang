package service

import (
	"github.com/ve-weiyi/ve-blog-golang/gin/api/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
)

type TagService struct {
	svcCtx *svctx.ServiceContext
}

func NewTagService(svcCtx *svctx.ServiceContext) *TagService {
	return &TagService{
		svcCtx: svcCtx,
	}
}

// 分页获取标签列表
func (s *TagService) FindTagList(reqCtx *request.Context, in *dto.TagQueryReq) (out *dto.PageResp, err error) {
	// todo

	return
}
