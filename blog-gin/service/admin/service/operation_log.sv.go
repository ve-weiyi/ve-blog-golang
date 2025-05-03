package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type OperationLogService struct {
	svcCtx *svctx.ServiceContext
}

func NewOperationLogService(svcCtx *svctx.ServiceContext) *OperationLogService {
	return &OperationLogService{
		svcCtx: svcCtx,
	}
}

// 删除操作记录
func (s *OperationLogService) DeletesOperationLog(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取操作记录列表
func (s *OperationLogService) FindOperationLogList(reqCtx *request.Context, in *dto.OperationLogQuery) (out *dto.PageResp, err error) {
	// todo

	return
}
