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
func (l *OperationLogService) CreateOperationLog(reqCtx *request.Context, operationLog *entity.OperationLog) (data *entity.OperationLog, err error) {
	return l.svcCtx.OperationLogRepository.Create(reqCtx, operationLog)
}

// 更新OperationLog记录
func (l *OperationLogService) UpdateOperationLog(reqCtx *request.Context, operationLog *entity.OperationLog) (data *entity.OperationLog, err error) {
	return l.svcCtx.OperationLogRepository.Update(reqCtx, operationLog)
}

// 删除OperationLog记录
func (l *OperationLogService) DeleteOperationLog(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.OperationLogRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询OperationLog记录
func (l *OperationLogService) FindOperationLog(reqCtx *request.Context, req *request.IdReq) (data *entity.OperationLog, err error) {
	return l.svcCtx.OperationLogRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除OperationLog记录
func (l *OperationLogService) DeleteOperationLogList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.OperationLogRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取OperationLog记录
func (l *OperationLogService) FindOperationLogList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.OperationLog, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.OperationLogRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.OperationLogRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
