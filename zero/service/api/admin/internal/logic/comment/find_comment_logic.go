package comment

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询评论
func NewFindCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentLogic {
	return &FindCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentLogic) FindComment(req *types.IdReq) (resp *types.CommentNewReq, err error) {
	// todo: add your logic here and delete this line

	return
}
