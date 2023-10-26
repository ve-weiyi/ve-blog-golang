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
func (s *ChatRecordService) CreateChatRecord(reqCtx *request.Context, chatRecord *entity.ChatRecord) (data *entity.ChatRecord, err error) {
	return s.svcCtx.ChatRecordRepository.CreateChatRecord(reqCtx, chatRecord)
}

// 更新ChatRecord记录
func (s *ChatRecordService) UpdateChatRecord(reqCtx *request.Context, chatRecord *entity.ChatRecord) (data *entity.ChatRecord, err error) {
	return s.svcCtx.ChatRecordRepository.UpdateChatRecord(reqCtx, chatRecord)
}

// 删除ChatRecord记录
func (s *ChatRecordService) DeleteChatRecord(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.ChatRecordRepository.DeleteChatRecordById(reqCtx, id)
}

// 查询ChatRecord记录
func (s *ChatRecordService) FindChatRecord(reqCtx *request.Context, id int) (data *entity.ChatRecord, err error) {
	return s.svcCtx.ChatRecordRepository.FindChatRecordById(reqCtx, id)
}

// 批量删除ChatRecord记录
func (s *ChatRecordService) DeleteChatRecordByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.ChatRecordRepository.DeleteChatRecordByIds(reqCtx, ids)
}

// 分页获取ChatRecord记录
func (s *ChatRecordService) FindChatRecordList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.ChatRecord, total int64, err error) {
	list, err = s.svcCtx.ChatRecordRepository.FindChatRecordList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.ChatRecordRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
