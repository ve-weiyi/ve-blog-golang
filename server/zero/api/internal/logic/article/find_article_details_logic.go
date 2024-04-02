package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindArticleDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleDetailsLogic {
	return &FindArticleDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleDetailsLogic) FindArticleDetails(req *types.IdReq) (resp *types.ArticlePageDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
