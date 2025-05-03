package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type VisitLogService struct {
	svcCtx *svctx.ServiceContext
}

func NewVisitLogService(svcCtx *svctx.ServiceContext) *VisitLogService {
	return &VisitLogService{
		svcCtx: svcCtx,
	}
}

// 删除操作记录
func (s *VisitLogService) DeletesVisitLog(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取操作记录列表
func (s *VisitLogService) FindVisitLogList(reqCtx *request.Context, in *dto.VisitLogQuery) (out *dto.PageResp, err error) {
	// todo

	return
}
