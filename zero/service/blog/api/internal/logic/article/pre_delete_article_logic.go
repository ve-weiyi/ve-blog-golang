package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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

func (l *PreDeleteArticleLogic) PreDeleteArticle(req *types.ArticlePreDeleteReq) (resp *types.EmptyResp, err error) {
	article, err := l.svcCtx.ArticleRpc.FindArticle(l.ctx, &blog.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	article.IsDelete = req.IsDelete
	_, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, article)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
