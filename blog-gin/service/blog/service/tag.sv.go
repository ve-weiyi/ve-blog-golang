package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
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
