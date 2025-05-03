package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type RemarkService struct {
	svcCtx *svctx.ServiceContext
}

func NewRemarkService(svcCtx *svctx.ServiceContext) *RemarkService {
	return &RemarkService{
		svcCtx: svcCtx,
	}
}

// 批量删除留言
func (s *RemarkService) BatchDeleteRemark(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 删除留言
func (s *RemarkService) DeleteRemark(reqCtx *request.Context, in *dto.IdReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取留言列表
func (s *RemarkService) FindRemarkList(reqCtx *request.Context, in *dto.RemarkQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 更新留言
func (s *RemarkService) UpdateRemarkReview(reqCtx *request.Context, in *dto.RemarkReviewReq) (out *dto.BatchResp, err error) {
	// todo

	return
}
