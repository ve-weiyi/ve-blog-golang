// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type ChatRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedChatRpcServer
}

func NewChatRpcServer(svcCtx *svc.ServiceContext) *ChatRpcServer {
	return &ChatRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建聊天记录
func (s *ChatRpcServer) AddChatRecord(ctx context.Context, in *blog.ChatRecord) (*blog.ChatRecord, error) {
	l := chatrpclogic.NewAddChatRecordLogic(ctx, s.svcCtx)
	return l.AddChatRecord(in)
}

// 更新聊天记录
func (s *ChatRpcServer) UpdateChatRecord(ctx context.Context, in *blog.ChatRecord) (*blog.ChatRecord, error) {
	l := chatrpclogic.NewUpdateChatRecordLogic(ctx, s.svcCtx)
	return l.UpdateChatRecord(in)
}

// 删除聊天记录
func (s *ChatRpcServer) DeleteChatRecord(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := chatrpclogic.NewDeleteChatRecordLogic(ctx, s.svcCtx)
	return l.DeleteChatRecord(in)
}

// 批量删除聊天记录
func (s *ChatRpcServer) DeleteChatRecordList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := chatrpclogic.NewDeleteChatRecordListLogic(ctx, s.svcCtx)
	return l.DeleteChatRecordList(in)
}

// 查询聊天记录
func (s *ChatRpcServer) FindChatRecord(ctx context.Context, in *blog.IdReq) (*blog.ChatRecord, error) {
	l := chatrpclogic.NewFindChatRecordLogic(ctx, s.svcCtx)
	return l.FindChatRecord(in)
}

// 查询聊天记录列表
func (s *ChatRpcServer) FindChatRecordList(ctx context.Context, in *blog.PageQuery) (*blog.ChatRecordPageResp, error) {
	l := chatrpclogic.NewFindChatRecordListLogic(ctx, s.svcCtx)
	return l.FindChatRecordList(in)
}

// 查询聊天记录数量
func (s *ChatRpcServer) FindChatRecordCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := chatrpclogic.NewFindChatRecordCountLogic(ctx, s.svcCtx)
	return l.FindChatRecordCount(in)
}
