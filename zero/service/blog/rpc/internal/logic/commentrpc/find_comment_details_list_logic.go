package commentrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentDetailsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentDetailsListLogic {
	return &FindCommentDetailsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取评论列表
func (l *FindCommentDetailsListLogic) FindCommentDetailsList(in *blog.PageQuery) (*blog.CommentDetailsPageResp, error) {
	// todo: add your logic here and delete this line

	return &blog.CommentDetailsPageResp{}, nil
}
