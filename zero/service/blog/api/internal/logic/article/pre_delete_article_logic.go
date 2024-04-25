package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreDeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章-逻辑删除
func NewPreDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreDeleteArticleLogic {
	return &PreDeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PreDeleteArticleLogic) PreDeleteArticle(reqCtx *types.RestHeader, req *types.ArticleDeleteReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
