package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveArticleLogic {
	return &SaveArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveArticleLogic) SaveArticle(req *types.ArticleDetailsDTOReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
