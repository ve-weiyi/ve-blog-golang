package comment

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/discussionservice"
)

type BatchReviewCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchReviewCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchReviewCommentsLogic {
	return &BatchReviewCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchReviewCommentsLogic) BatchReviewComments(req *types.BatchReviewCommentsReq) (resp *types.BatchResp, err error) {
	out, err := l.svcCtx.DiscussionService.BatchReviewComments(l.ctx, &discussionservice.BatchReviewCommentsRequest{
		Ids:      req.Ids,
		IsReview: req.IsReview,
	})
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
