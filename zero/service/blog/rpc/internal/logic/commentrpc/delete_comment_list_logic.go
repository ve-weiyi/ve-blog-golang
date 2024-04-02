package commentrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentListLogic {
	return &DeleteCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除评论
func (l *DeleteCommentListLogic) DeleteCommentList(in *blog.IdsReq) (*blog.BatchResp, error) {
	result, err := l.svcCtx.CommentModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: result,
	}, nil
}
