package comment

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentListLogic {
	return &DeleteCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentListLogic) DeleteCommentList(req *types.IdsReq) (resp *types.BatchResult, err error) {
	// todo: add your logic here and delete this line

	return
}
