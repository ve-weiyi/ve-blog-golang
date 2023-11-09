package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
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
	return s.svcCtx.TalkRepository.CreateTalk(reqCtx, talk)
}

// 更新Talk记录
func (s *TalkService) UpdateTalk(reqCtx *request.Context, talk *entity.Talk) (data *entity.Talk, err error) {
	return s.svcCtx.TalkRepository.UpdateTalk(reqCtx, talk)
}

// 删除Talk记录
func (s *TalkService) DeleteTalk(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.TalkRepository.DeleteTalkById(reqCtx, id)
}

// 查询Talk记录
func (s *TalkService) FindTalk(reqCtx *request.Context, id int) (data *entity.Talk, err error) {
	return s.svcCtx.TalkRepository.FindTalkById(reqCtx, id)
}

// 批量删除Talk记录
func (s *TalkService) DeleteTalkByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.TalkRepository.DeleteTalkByIds(reqCtx, ids)
}

// 分页获取Talk记录
func (s *TalkService) FindTalkList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Talk, total int64, err error) {
	list, err = s.svcCtx.TalkRepository.FindTalkList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.TalkRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
