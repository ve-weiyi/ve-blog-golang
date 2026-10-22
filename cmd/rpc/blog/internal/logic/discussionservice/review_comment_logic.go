package discussionservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/discussionrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

type ReviewCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReviewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewCommentLogic {
	return &ReviewCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReviewCommentLogic) ReviewComment(in *discussionrpc.ReviewCommentRequest) (*discussionrpc.ReviewCommentResponse, error) {
	if in.IsReview != enums.CommentIsReviewNo && in.IsReview != enums.CommentIsReviewYes {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "审核状态不合法")
	}

	if _, err := l.svcCtx.TCommentModel.FindById(l.ctx, in.Id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.NewBizError(bizcode.CodeResourceNotFound, "评论不存在")
		}
		return nil, err
	}

	if _, err := l.svcCtx.TCommentModel.UpdateFields(l.ctx,
		map[string]interface{}{"is_review": in.IsReview}, "id = ?", in.Id); err != nil {
		return nil, err
	}

	return &discussionrpc.ReviewCommentResponse{Success: true}, nil
}
