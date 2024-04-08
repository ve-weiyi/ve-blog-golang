package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type ChatRecordService struct {
	svcCtx *svc.ServiceContext
}

func NewChatRecordService(svcCtx *svc.ServiceContext) *ChatRecordService {
	return &ChatRecordService{
		svcCtx: svcCtx,
	}
}

// 创建ChatRecord记录
func (l *ChatRecordService) CreateChatRecord(reqCtx *request.Context, chatRecord *entity.ChatRecord) (data *entity.ChatRecord, err error) {
	return l.svcCtx.ChatRecordRepository.Create(reqCtx, chatRecord)
}

// 更新ChatRecord记录
func (l *ChatRecordService) UpdateChatRecord(reqCtx *request.Context, chatRecord *entity.ChatRecord) (data *entity.ChatRecord, err error) {
	return l.svcCtx.ChatRecordRepository.Update(reqCtx, chatRecord)
}

// 删除ChatRecord记录
func (l *ChatRecordService) DeleteChatRecord(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.ChatRecordRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询ChatRecord记录
func (l *ChatRecordService) FindChatRecord(reqCtx *request.Context, req *request.IdReq) (data *entity.ChatRecord, err error) {
	return l.svcCtx.ChatRecordRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除ChatRecord记录
func (l *ChatRecordService) DeleteChatRecordList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.ChatRecordRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取ChatRecord记录
func (l *ChatRecordService) FindChatRecordList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.ChatRecord, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.ChatRecordRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.ChatRecordRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
