package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleLogic {
	return &FindArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleLogic) FindArticle(req *types.IdReq) (resp *types.ArticleBack, err error) {
	// todo: add your logic here and delete this line

	return
}
