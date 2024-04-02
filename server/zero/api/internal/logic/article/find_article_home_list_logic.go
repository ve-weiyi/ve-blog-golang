package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleHomeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindArticleHomeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleHomeListLogic {
	return &FindArticleHomeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleHomeListLogic) FindArticleHomeList(req *types.PageQuery) (resp []types.ArticleHome, err error) {
	// todo: add your logic here and delete this line

	return
}
