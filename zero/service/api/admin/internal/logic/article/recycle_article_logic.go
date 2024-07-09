package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecycleArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 回收文章
func NewRecycleArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecycleArticleLogic {
	return &RecycleArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecycleArticleLogic) RecycleArticle(req *types.ArticleRecycleReq) (resp *types.EmptyResp, err error) {
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
