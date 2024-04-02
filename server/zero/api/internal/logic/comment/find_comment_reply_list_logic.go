package comment

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentReplyListLogic {
	return &FindCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentReplyListLogic) FindCommentReplyList(req *types.PageQuery) (resp []types.ReplyDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
