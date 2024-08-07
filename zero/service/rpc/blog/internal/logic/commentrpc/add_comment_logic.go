package commentrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/commentrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建评论
func (l *AddCommentLogic) AddComment(in *commentrpc.CommentNew) (*commentrpc.CommentDetails, error) {
	entity := ConvertCommentIn(in)

	_, err := l.svcCtx.CommentModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return ConvertCommentOut(entity), nil
}
