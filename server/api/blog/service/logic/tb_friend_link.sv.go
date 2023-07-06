package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
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
	return s.svcCtx.FriendLinkRepository.CreateFriendLink(friendLink)
}

// 删除FriendLink记录
func (s *FriendLinkService) DeleteFriendLink(reqCtx *request.Context, friendLink *entity.FriendLink) (rows int64, err error) {
	return s.svcCtx.FriendLinkRepository.DeleteFriendLink(friendLink)
}

// 更新FriendLink记录
func (s *FriendLinkService) UpdateFriendLink(reqCtx *request.Context, friendLink *entity.FriendLink) (data *entity.FriendLink, err error) {
	return s.svcCtx.FriendLinkRepository.UpdateFriendLink(friendLink)
}

// 查询FriendLink记录
func (s *FriendLinkService) GetFriendLink(reqCtx *request.Context, friendLink *entity.FriendLink) (data *entity.FriendLink, err error) {
	return s.svcCtx.FriendLinkRepository.GetFriendLink(friendLink.ID)
}

// 批量删除FriendLink记录
func (s *FriendLinkService) DeleteFriendLinkByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.FriendLinkRepository.DeleteFriendLinkByIds(ids)
}

// 分页获取FriendLink记录
func (s *FriendLinkService) FindFriendLinkList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.FriendLink, total int64, err error) {
	return s.svcCtx.FriendLinkRepository.FindFriendLinkList(page)
}
