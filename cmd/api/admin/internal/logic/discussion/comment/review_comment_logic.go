package comment

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/discussionservice"
)

type ReviewCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewCommentLogic {
	return &ReviewCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewCommentLogic) ReviewComment(req *types.ReviewCommentReq) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.DiscussionService.ReviewComment(l.ctx, &discussionservice.ReviewCommentRequest{
		Id:       req.CommentId,
		IsReview: req.IsReview,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
