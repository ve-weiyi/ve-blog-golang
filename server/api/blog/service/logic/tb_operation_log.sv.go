package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type OperationLogService struct {
	svcCtx *svc.ServiceContext
}

func NewOperationLogService(svcCtx *svc.ServiceContext) *OperationLogService {
	return &OperationLogService{
		svcCtx: svcCtx,
	}
}

// 创建OperationLog记录
func (s *OperationLogService) CreateOperationLog(reqCtx *request.Context, operationLog *entity.OperationLog) (data *entity.OperationLog, err error) {
	return s.svcCtx.OperationLogRepository.CreateOperationLog(reqCtx, operationLog)
}

// 删除OperationLog记录
func (s *OperationLogService) DeleteOperationLog(reqCtx *request.Context, operationLog *entity.OperationLog) (rows int64, err error) {
	return s.svcCtx.OperationLogRepository.DeleteOperationLog(reqCtx, operationLog)
}

// 更新OperationLog记录
func (s *OperationLogService) UpdateOperationLog(reqCtx *request.Context, operationLog *entity.OperationLog) (data *entity.OperationLog, err error) {
	return s.svcCtx.OperationLogRepository.UpdateOperationLog(reqCtx, operationLog)
}

// 查询OperationLog记录
func (s *OperationLogService) FindOperationLog(reqCtx *request.Context, operationLog *entity.OperationLog) (data *entity.OperationLog, err error) {
	return s.svcCtx.OperationLogRepository.GetOperationLog(reqCtx, operationLog.ID)
}

// 批量删除OperationLog记录
func (s *OperationLogService) DeleteOperationLogByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.OperationLogRepository.DeleteOperationLogByIds(reqCtx, ids)
}

// 分页获取OperationLog记录
func (s *OperationLogService) FindOperationLogList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.OperationLog, total int64, err error) {
	return s.svcCtx.OperationLogRepository.FindOperationLogList(reqCtx, page)
}
