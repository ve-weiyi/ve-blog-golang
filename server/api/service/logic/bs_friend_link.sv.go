package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type FriendLinkService struct {
	svcCtx *svc.ServiceContext
}

func NewFriendLinkService(svcCtx *svc.ServiceContext) *FriendLinkService {
	return &FriendLinkService{
		svcCtx: svcCtx,
	}
}

// 创建FriendLink记录
func (s *FriendLinkService) CreateFriendLink(reqCtx *request.Context, friendLink *entity.FriendLink) (data *entity.FriendLink, err error) {
	return s.svcCtx.FriendLinkRepository.CreateFriendLink(reqCtx, friendLink)
}

// 更新FriendLink记录
func (s *FriendLinkService) UpdateFriendLink(reqCtx *request.Context, friendLink *entity.FriendLink) (data *entity.FriendLink, err error) {
	return s.svcCtx.FriendLinkRepository.UpdateFriendLink(reqCtx, friendLink)
}

// 删除FriendLink记录
func (s *FriendLinkService) DeleteFriendLink(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.FriendLinkRepository.DeleteFriendLink(reqCtx, id)
}

// 查询FriendLink记录
func (s *FriendLinkService) FindFriendLink(reqCtx *request.Context, id int) (data *entity.FriendLink, err error) {
	return s.svcCtx.FriendLinkRepository.FindFriendLink(reqCtx, id)
}

// 批量删除FriendLink记录
func (s *FriendLinkService) DeleteFriendLinkByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.FriendLinkRepository.DeleteFriendLinkByIds(reqCtx, ids)
}

// 分页获取FriendLink记录
func (s *FriendLinkService) FindFriendLinkList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.FriendLink, total int64, err error) {
	list, err = s.svcCtx.FriendLinkRepository.FindFriendLinkList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.FriendLinkRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
