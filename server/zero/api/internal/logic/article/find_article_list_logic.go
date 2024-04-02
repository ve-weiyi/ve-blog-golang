package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleListLogic {
	return &FindArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleListLogic) FindArticleList(req *types.PageQuery) (resp []types.ArticleBack, err error) {
	// todo: add your logic here and delete this line

	return
}
