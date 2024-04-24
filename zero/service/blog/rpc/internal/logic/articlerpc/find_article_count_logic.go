package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticleCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleCountLogic {
	return &FindArticleCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindArticleCountLogic) FindArticleCount(in *blog.PageQuery) (*blog.CountResp, error) {
	_, _, _, conditions, params := convert.ParsePageQuery(in)

	count, err := l.svcCtx.ArticleModel.FindCount(l.ctx, conditions, params)
	if err != nil {
		return nil, err
	}

	return &blog.CountResp{
		Count: count,
	}, nil
}
