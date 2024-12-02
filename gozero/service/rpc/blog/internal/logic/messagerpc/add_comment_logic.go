package messagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建评论
func (l *AddCommentLogic) AddComment(in *messagerpc.CommentNewReq) (*messagerpc.CommentDetails, error) {
	entity := convertCommentIn(in)

	_, err := l.svcCtx.TCommentModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertCommentOut(entity), nil
}

func convertCommentIn(in *messagerpc.CommentNewReq) (out *model.TComment) {
	out = &model.TComment{
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		ReplyMsgId:     in.ReplyMsgId,
		UserId:         in.UserId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
		IsReview:       in.IsReview,
		// CreatedAt:      time.Unix(in.CreatedAt, 0),
		// UpdatedAt:      time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func convertCommentOut(in *model.TComment) (out *messagerpc.CommentDetails) {
	out = &messagerpc.CommentDetails{
		Id:             in.Id,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		ReplyMsgId:     in.ReplyMsgId,
		UserId:         in.UserId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
		LikeCount:      in.LikeCount,
	}

	return out
}
