package discussionservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/model"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/discussionrpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建评论
func (l *CreateCommentLogic) CreateComment(in *discussionrpc.CreateCommentRequest) (*discussionrpc.CreateCommentResponse, error) {
	entity := &model.TComment{
		UserId:         in.UserId,
		DeviceId:       in.DeviceId,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		ReplyId:        in.ReplyId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
	}

	_, err := l.svcCtx.TCommentModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &discussionrpc.CreateCommentResponse{Id: entity.Id}, nil
}
