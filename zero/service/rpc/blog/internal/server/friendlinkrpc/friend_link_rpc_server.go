// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/friendlinkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

type FriendLinkRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedFriendLinkRpcServer
}

func NewFriendLinkRpcServer(svcCtx *svc.ServiceContext) *FriendLinkRpcServer {
	return &FriendLinkRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建友链
func (s *FriendLinkRpcServer) AddFriendLink(ctx context.Context, in *blog.FriendLink) (*blog.FriendLink, error) {
	l := friendlinkrpclogic.NewAddFriendLinkLogic(ctx, s.svcCtx)
	return l.AddFriendLink(in)
}

// 更新友链
func (s *FriendLinkRpcServer) UpdateFriendLink(ctx context.Context, in *blog.FriendLink) (*blog.FriendLink, error) {
	l := friendlinkrpclogic.NewUpdateFriendLinkLogic(ctx, s.svcCtx)
	return l.UpdateFriendLink(in)
}

// 删除友链
func (s *FriendLinkRpcServer) DeleteFriendLink(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := friendlinkrpclogic.NewDeleteFriendLinkLogic(ctx, s.svcCtx)
	return l.DeleteFriendLink(in)
}

// 批量删除友链
func (s *FriendLinkRpcServer) DeleteFriendLinkList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := friendlinkrpclogic.NewDeleteFriendLinkListLogic(ctx, s.svcCtx)
	return l.DeleteFriendLinkList(in)
}

// 查询友链
func (s *FriendLinkRpcServer) FindFriendLink(ctx context.Context, in *blog.IdReq) (*blog.FriendLink, error) {
	l := friendlinkrpclogic.NewFindFriendLinkLogic(ctx, s.svcCtx)
	return l.FindFriendLink(in)
}

// 查询友链列表
func (s *FriendLinkRpcServer) FindFriendLinkList(ctx context.Context, in *blog.PageQuery) (*blog.FriendLinkPageResp, error) {
	l := friendlinkrpclogic.NewFindFriendLinkListLogic(ctx, s.svcCtx)
	return l.FindFriendLinkList(in)
}

// 查询友链数量
func (s *FriendLinkRpcServer) FindFriendLinkCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := friendlinkrpclogic.NewFindFriendLinkCountLogic(ctx, s.svcCtx)
	return l.FindFriendLinkCount(in)
}