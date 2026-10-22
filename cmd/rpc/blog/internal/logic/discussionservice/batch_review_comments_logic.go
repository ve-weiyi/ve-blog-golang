package discussionservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/discussionrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

type BatchReviewCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchReviewCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchReviewCommentsLogic {
	return &BatchReviewCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchReviewCommentsLogic) BatchReviewComments(in *discussionrpc.BatchReviewCommentsRequest) (*discussionrpc.BatchReviewCommentsResponse, error) {
	if len(in.Ids) == 0 {
		return &discussionrpc.BatchReviewCommentsResponse{}, nil
	}
	if in.IsReview != enums.CommentIsReviewNo && in.IsReview != enums.CommentIsReviewYes {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "审核状态不合法")
	}

	rows, err := l.svcCtx.TCommentModel.UpdateFields(l.ctx,
		map[string]interface{}{"is_review": in.IsReview}, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &discussionrpc.BatchReviewCommentsResponse{SuccessCount: rows}, nil
}
