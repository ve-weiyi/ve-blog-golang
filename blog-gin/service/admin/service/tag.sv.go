package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
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

// 创建标签
func (s *TagService) AddTag(reqCtx *request.Context, in *dto.TagNewReq) (out *dto.TagBackVO, err error) {
	// todo

	return
}

// 批量删除标签
func (s *TagService) BatchDeleteTag(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 删除标签
func (s *TagService) DeleteTag(reqCtx *request.Context, in *dto.IdReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取标签列表
func (s *TagService) FindTagList(reqCtx *request.Context, in *dto.TagQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 更新标签
func (s *TagService) UpdateTag(reqCtx *request.Context, in *dto.TagNewReq) (out *dto.TagBackVO, err error) {
	// todo

	return
}
