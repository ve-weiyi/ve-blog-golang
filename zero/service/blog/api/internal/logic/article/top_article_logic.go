package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 置顶文章
func NewTopArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopArticleLogic {
	return &TopArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopArticleLogic) TopArticle(req *types.ArticleTopReq) (resp *types.EmptyResp, err error) {
	article, err := l.svcCtx.ArticleRpc.FindArticle(l.ctx, &blog.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	article.IsTop = req.IsTop
	_, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, article)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
