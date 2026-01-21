package newsrpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/enums"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rpcutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
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
func (l *AddCommentLogic) AddComment(in *newsrpc.AddCommentReq) (*newsrpc.AddCommentResp, error) {
	uid, _ := rpcutils.GetUserIdFromCtx(l.ctx)
	tid, _ := rpcutils.GetTerminalIdFromCtx(l.ctx)

	entity := &model.TComment{
		Id:             0,
		UserId:         uid,
		TerminalId:     tid,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		ReplyId:        in.ReplyId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         enums.CommentStatusNormal,
	}

	_, err := l.svcCtx.TCommentModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &newsrpc.AddCommentResp{
		Comment: convertCommentOut(entity),
	}, nil
}
