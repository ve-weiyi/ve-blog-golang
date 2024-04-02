package comment

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentDetailsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindCommentDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentDetailsListLogic {
	return &FindCommentDetailsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentDetailsListLogic) FindCommentDetailsList(req *types.PageQuery) (resp []types.CommentDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
