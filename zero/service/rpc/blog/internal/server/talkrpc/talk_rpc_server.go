// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type TalkRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedTalkRpcServer
}

func NewTalkRpcServer(svcCtx *svc.ServiceContext) *TalkRpcServer {
	return &TalkRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建说说
func (s *TalkRpcServer) AddTalk(ctx context.Context, in *blog.Talk) (*blog.Talk, error) {
	l := talkrpclogic.NewAddTalkLogic(ctx, s.svcCtx)
	return l.AddTalk(in)
}

// 更新说说
func (s *TalkRpcServer) UpdateTalk(ctx context.Context, in *blog.Talk) (*blog.Talk, error) {
	l := talkrpclogic.NewUpdateTalkLogic(ctx, s.svcCtx)
	return l.UpdateTalk(in)
}

// 删除说说
func (s *TalkRpcServer) DeleteTalk(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := talkrpclogic.NewDeleteTalkLogic(ctx, s.svcCtx)
	return l.DeleteTalk(in)
}

// 批量删除说说
func (s *TalkRpcServer) DeleteTalkList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := talkrpclogic.NewDeleteTalkListLogic(ctx, s.svcCtx)
	return l.DeleteTalkList(in)
}

// 查询说说
func (s *TalkRpcServer) FindTalk(ctx context.Context, in *blog.IdReq) (*blog.Talk, error) {
	l := talkrpclogic.NewFindTalkLogic(ctx, s.svcCtx)
	return l.FindTalk(in)
}

// 查询说说列表
func (s *TalkRpcServer) FindTalkList(ctx context.Context, in *blog.PageQuery) (*blog.FindTalkListResp, error) {
	l := talkrpclogic.NewFindTalkListLogic(ctx, s.svcCtx)
	return l.FindTalkList(in)
}

// 查询说说数量
func (s *TalkRpcServer) FindTalkCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := talkrpclogic.NewFindTalkCountLogic(ctx, s.svcCtx)
	return l.FindTalkCount(in)
}

// 点赞说说
func (s *TalkRpcServer) LikeTalk(ctx context.Context, in *blog.IdReq) (*blog.EmptyResp, error) {
	l := talkrpclogic.NewLikeTalkLogic(ctx, s.svcCtx)
	return l.LikeTalk(in)
}

// 用户点赞的说说
func (s *TalkRpcServer) FindUserLikeTalk(ctx context.Context, in *blog.UserIdReq) (*blog.FindLikeTalkResp, error) {
	l := talkrpclogic.NewFindUserLikeTalkLogic(ctx, s.svcCtx)
	return l.FindUserLikeTalk(in)
}
