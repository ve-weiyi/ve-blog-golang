package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleSeriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindArticleSeriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleSeriesLogic {
	return &FindArticleSeriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleSeriesLogic) FindArticleSeries(req *types.ArticleConditionReq) (resp *types.ArticleConditionDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
