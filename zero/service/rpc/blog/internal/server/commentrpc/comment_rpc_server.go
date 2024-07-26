// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/commentrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

type CommentRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedCommentRpcServer
}

func NewCommentRpcServer(svcCtx *svc.ServiceContext) *CommentRpcServer {
	return &CommentRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建评论
func (s *CommentRpcServer) AddComment(ctx context.Context, in *blog.Comment) (*blog.Comment, error) {
	l := commentrpclogic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

// 更新评论
func (s *CommentRpcServer) UpdateComment(ctx context.Context, in *blog.Comment) (*blog.Comment, error) {
	l := commentrpclogic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

// 删除评论
func (s *CommentRpcServer) DeleteComment(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := commentrpclogic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

// 批量删除评论
func (s *CommentRpcServer) DeleteCommentList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := commentrpclogic.NewDeleteCommentListLogic(ctx, s.svcCtx)
	return l.DeleteCommentList(in)
}

// 查询评论
func (s *CommentRpcServer) FindComment(ctx context.Context, in *blog.IdReq) (*blog.Comment, error) {
	l := commentrpclogic.NewFindCommentLogic(ctx, s.svcCtx)
	return l.FindComment(in)
}

// 查询评论列表
func (s *CommentRpcServer) FindCommentList(ctx context.Context, in *blog.PageQuery) (*blog.CommentPageResp, error) {
	l := commentrpclogic.NewFindCommentListLogic(ctx, s.svcCtx)
	return l.FindCommentList(in)
}

// 查询评论数量
func (s *CommentRpcServer) FindCommentCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := commentrpclogic.NewFindCommentCountLogic(ctx, s.svcCtx)
	return l.FindCommentCount(in)
}

// 点赞评论
func (s *CommentRpcServer) LikeComment(ctx context.Context, in *blog.IdReq) (*blog.EmptyResp, error) {
	l := commentrpclogic.NewLikeCommentLogic(ctx, s.svcCtx)
	return l.LikeComment(in)
}