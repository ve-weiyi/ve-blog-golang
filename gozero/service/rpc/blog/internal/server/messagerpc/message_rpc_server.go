// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/logic/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
)

type MessageRpcServer struct {
	svcCtx *svc.ServiceContext
	messagerpc.UnimplementedMessageRpcServer
}

func NewMessageRpcServer(svcCtx *svc.ServiceContext) *MessageRpcServer {
	return &MessageRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建聊天记录
func (s *MessageRpcServer) AddChatMessage(ctx context.Context, in *messagerpc.ChatMessageNewReq) (*messagerpc.ChatMessageDetails, error) {
	l := messagerpclogic.NewAddChatMessageLogic(ctx, s.svcCtx)
	return l.AddChatMessage(in)
}

// 更新聊天记录
func (s *MessageRpcServer) UpdateChatMessage(ctx context.Context, in *messagerpc.ChatMessageNewReq) (*messagerpc.ChatMessageDetails, error) {
	l := messagerpclogic.NewUpdateChatMessageLogic(ctx, s.svcCtx)
	return l.UpdateChatMessage(in)
}

// 删除聊天记录
func (s *MessageRpcServer) DeletesChatMessage(ctx context.Context, in *messagerpc.IdsReq) (*messagerpc.BatchResp, error) {
	l := messagerpclogic.NewDeletesChatMessageLogic(ctx, s.svcCtx)
	return l.DeletesChatMessage(in)
}

// 查询聊天记录
func (s *MessageRpcServer) GetChatMessage(ctx context.Context, in *messagerpc.IdReq) (*messagerpc.ChatMessageDetails, error) {
	l := messagerpclogic.NewGetChatMessageLogic(ctx, s.svcCtx)
	return l.GetChatMessage(in)
}

// 查询聊天记录列表
func (s *MessageRpcServer) FindChatMessageList(ctx context.Context, in *messagerpc.FindChatMessageListReq) (*messagerpc.FindChatMessageListResp, error) {
	l := messagerpclogic.NewFindChatMessageListLogic(ctx, s.svcCtx)
	return l.FindChatMessageList(in)
}

// 创建留言
func (s *MessageRpcServer) AddRemark(ctx context.Context, in *messagerpc.RemarkNewReq) (*messagerpc.RemarkDetails, error) {
	l := messagerpclogic.NewAddRemarkLogic(ctx, s.svcCtx)
	return l.AddRemark(in)
}

// 更新留言
func (s *MessageRpcServer) UpdateRemark(ctx context.Context, in *messagerpc.RemarkUpdateReq) (*messagerpc.RemarkDetails, error) {
	l := messagerpclogic.NewUpdateRemarkLogic(ctx, s.svcCtx)
	return l.UpdateRemark(in)
}

// 删除留言
func (s *MessageRpcServer) DeletesRemark(ctx context.Context, in *messagerpc.IdsReq) (*messagerpc.BatchResp, error) {
	l := messagerpclogic.NewDeletesRemarkLogic(ctx, s.svcCtx)
	return l.DeletesRemark(in)
}

// 查询留言
func (s *MessageRpcServer) GetRemark(ctx context.Context, in *messagerpc.IdReq) (*messagerpc.RemarkDetails, error) {
	l := messagerpclogic.NewGetRemarkLogic(ctx, s.svcCtx)
	return l.GetRemark(in)
}

// 查询留言列表
func (s *MessageRpcServer) FindRemarkList(ctx context.Context, in *messagerpc.FindRemarkListReq) (*messagerpc.FindRemarkListResp, error) {
	l := messagerpclogic.NewFindRemarkListLogic(ctx, s.svcCtx)
	return l.FindRemarkList(in)
}

// 创建评论
func (s *MessageRpcServer) AddComment(ctx context.Context, in *messagerpc.CommentNewReq) (*messagerpc.CommentDetails, error) {
	l := messagerpclogic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

// 删除评论
func (s *MessageRpcServer) DeleteComment(ctx context.Context, in *messagerpc.IdsReq) (*messagerpc.BatchResp, error) {
	l := messagerpclogic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

// 查询评论
func (s *MessageRpcServer) GetComment(ctx context.Context, in *messagerpc.IdReq) (*messagerpc.CommentDetails, error) {
	l := messagerpclogic.NewGetCommentLogic(ctx, s.svcCtx)
	return l.GetComment(in)
}

// 查询评论列表
func (s *MessageRpcServer) FindCommentList(ctx context.Context, in *messagerpc.FindCommentListReq) (*messagerpc.FindCommentListResp, error) {
	l := messagerpclogic.NewFindCommentListLogic(ctx, s.svcCtx)
	return l.FindCommentList(in)
}

// 查询评论回复列表
func (s *MessageRpcServer) FindCommentReplyList(ctx context.Context, in *messagerpc.FindCommentReplyListReq) (*messagerpc.FindCommentReplyListResp, error) {
	l := messagerpclogic.NewFindCommentReplyListLogic(ctx, s.svcCtx)
	return l.FindCommentReplyList(in)
}

// 查询评论回复数量
func (s *MessageRpcServer) FindTopicCommentCounts(ctx context.Context, in *messagerpc.IdsReq) (*messagerpc.FindTopicCommentCountsResp, error) {
	l := messagerpclogic.NewFindTopicCommentCountsLogic(ctx, s.svcCtx)
	return l.FindTopicCommentCounts(in)
}

// 更新评论审核状态
func (s *MessageRpcServer) UpdateCommentReview(ctx context.Context, in *messagerpc.UpdateCommentReviewReq) (*messagerpc.BatchResp, error) {
	l := messagerpclogic.NewUpdateCommentReviewLogic(ctx, s.svcCtx)
	return l.UpdateCommentReview(in)
}

// 更新评论
func (s *MessageRpcServer) UpdateCommentContent(ctx context.Context, in *messagerpc.UpdateCommentContentReq) (*messagerpc.CommentDetails, error) {
	l := messagerpclogic.NewUpdateCommentContentLogic(ctx, s.svcCtx)
	return l.UpdateCommentContent(in)
}

// 点赞评论
func (s *MessageRpcServer) LikeComment(ctx context.Context, in *messagerpc.IdReq) (*messagerpc.EmptyResp, error) {
	l := messagerpclogic.NewLikeCommentLogic(ctx, s.svcCtx)
	return l.LikeComment(in)
}

// 用户点赞的评论
func (s *MessageRpcServer) FindUserLikeComment(ctx context.Context, in *messagerpc.UserIdReq) (*messagerpc.FindLikeCommentResp, error) {
	l := messagerpclogic.NewFindUserLikeCommentLogic(ctx, s.svcCtx)
	return l.FindUserLikeComment(in)
}
