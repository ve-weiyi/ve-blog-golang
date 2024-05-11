package commentrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentListLogic {
	return &FindCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取评论列表
func (l *FindCommentListLogic) FindCommentList(in *blog.PageQuery) (*blog.CommentPageResp, error) {
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.CommentModel.FindList(l.ctx, limit, offset, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.Comment
	for _, v := range result {
		list = append(list, convert.ConvertCommentModelToPb(v))
	}

	return &blog.CommentPageResp{
		List: list,
	}, nil
}
