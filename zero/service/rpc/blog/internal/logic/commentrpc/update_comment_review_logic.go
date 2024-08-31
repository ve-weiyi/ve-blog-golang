package commentrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/commentrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentReviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentReviewLogic {
	return &UpdateCommentReviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新评论审核状态
func (l *UpdateCommentReviewLogic) UpdateCommentReview(in *commentrpc.UpdateCommentReviewReq) (*commentrpc.CommentDetails, error) {
	entity, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	entity.IsReview = in.IsReview

	_, err = l.svcCtx.CommentModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertCommentOut(entity), nil
}
