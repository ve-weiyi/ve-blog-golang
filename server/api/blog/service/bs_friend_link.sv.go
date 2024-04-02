package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
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
func (l *FriendLinkService) CreateFriendLink(reqCtx *request.Context, friendLink *entity.FriendLink) (data *entity.FriendLink, err error) {
	return l.svcCtx.FriendLinkRepository.Create(reqCtx, friendLink)
}

// 更新FriendLink记录
func (l *FriendLinkService) UpdateFriendLink(reqCtx *request.Context, friendLink *entity.FriendLink) (data *entity.FriendLink, err error) {
	return l.svcCtx.FriendLinkRepository.Update(reqCtx, friendLink)
}

// 删除FriendLink记录
func (l *FriendLinkService) DeleteFriendLink(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.FriendLinkRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询FriendLink记录
func (l *FriendLinkService) FindFriendLink(reqCtx *request.Context, req *request.IdReq) (data *entity.FriendLink, err error) {
	return l.svcCtx.FriendLinkRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除FriendLink记录
func (l *FriendLinkService) DeleteFriendLinkList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.FriendLinkRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取FriendLink记录
func (l *FriendLinkService) FindFriendLinkList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.FriendLink, total int64, err error) {
	limit := page.Limit
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.FriendLinkRepository.FindList(reqCtx, limit.Page, limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.FriendLinkRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
