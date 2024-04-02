package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleArchivesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindArticleArchivesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleArchivesLogic {
	return &FindArticleArchivesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleArchivesLogic) FindArticleArchives(req *types.PageQuery) (resp []types.ArticlePreviewDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
