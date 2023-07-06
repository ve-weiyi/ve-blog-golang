package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
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
	s.svcCtx.Log.Info("创建ChatRecord记录")
	return s.svcCtx.ChatRecordRepository.CreateChatRecord(reqCtx, chatRecord)
}

// 删除ChatRecord记录
func (s *ChatRecordService) DeleteChatRecord(reqCtx *request.Context, chatRecord *entity.ChatRecord) (rows int64, err error) {
	return s.svcCtx.ChatRecordRepository.DeleteChatRecord(reqCtx, chatRecord)
}

// 更新ChatRecord记录
func (s *ChatRecordService) UpdateChatRecord(reqCtx *request.Context, chatRecord *entity.ChatRecord) (data *entity.ChatRecord, err error) {
	return s.svcCtx.ChatRecordRepository.UpdateChatRecord(reqCtx, chatRecord)
}

// 查询ChatRecord记录
func (s *ChatRecordService) GetChatRecord(reqCtx *request.Context, chatRecord *entity.ChatRecord) (data *entity.ChatRecord, err error) {
	return s.svcCtx.ChatRecordRepository.GetChatRecord(reqCtx, chatRecord.ID)
}

// 批量删除ChatRecord记录
func (s *ChatRecordService) DeleteChatRecordByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.ChatRecordRepository.DeleteChatRecordByIds(reqCtx, ids)
}

// 分页获取ChatRecord记录
func (s *ChatRecordService) FindChatRecordList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.ChatRecord, total int64, err error) {
	return s.svcCtx.ChatRecordRepository.FindChatRecordList(reqCtx, page)
}
