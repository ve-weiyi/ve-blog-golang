package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTagCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagCountLogic {
	return &FindTagCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章标签数量
func (l *FindTagCountLogic) FindTagCount(in *blog.PageQuery) (*blog.CountResp, error) {
	_, _, _, conditions, params := convert.ParsePageQuery(in)

	count, err := l.svcCtx.TagModel.FindCount(l.ctx, conditions, params)
	if err != nil {
		return nil, err
	}

	return &blog.CountResp{
		Count: count,
	}, nil
}
