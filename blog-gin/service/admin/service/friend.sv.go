package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type FriendService struct {
	svcCtx *svctx.ServiceContext
}

func NewFriendService(svcCtx *svctx.ServiceContext) *FriendService {
	return &FriendService{
		svcCtx: svcCtx,
	}
}

// 创建友链
func (s *FriendService) AddFriend(reqCtx *request.Context, in *dto.FriendNewReq) (out *dto.FriendBackVO, err error) {
	// todo

	return
}

// 删除友链
func (s *FriendService) DeletesFriend(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取友链列表
func (s *FriendService) FindFriendList(reqCtx *request.Context, in *dto.FriendQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 更新友链
func (s *FriendService) UpdateFriend(reqCtx *request.Context, in *dto.FriendNewReq) (out *dto.FriendBackVO, err error) {
	// todo

	return
}
