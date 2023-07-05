package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type TalkService struct {
	svcCtx *svc.ServiceContext
}

func NewTalkService(svcCtx *svc.ServiceContext) *TalkService {
	return &TalkService{
		svcCtx: svcCtx,
	}
}

// 创建Talk记录
func (s *TalkService) CreateTalk(reqCtx *request.Context, talk *entity.Talk) (data *entity.Talk, err error) {
	return s.svcCtx.TalkRepository.CreateTalk(talk)
}

// 删除Talk记录
func (s *TalkService) DeleteTalk(reqCtx *request.Context, talk *entity.Talk) (rows int64, err error) {
	return s.svcCtx.TalkRepository.DeleteTalk(talk)
}

// 更新Talk记录
func (s *TalkService) UpdateTalk(reqCtx *request.Context, talk *entity.Talk) (data *entity.Talk, err error) {
	return s.svcCtx.TalkRepository.UpdateTalk(talk)
}

// 查询Talk记录
func (s *TalkService) GetTalk(reqCtx *request.Context, talk *entity.Talk) (data *entity.Talk, err error) {
	return s.svcCtx.TalkRepository.GetTalk(talk.ID)
}

// 批量删除Talk记录
func (s *TalkService) DeleteTalkByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.TalkRepository.DeleteTalkByIds(ids)
}

// 分页获取Talk记录
func (s *TalkService) FindTalkList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Talk, total int64, err error) {
	return s.svcCtx.TalkRepository.FindTalkList(page)
}
