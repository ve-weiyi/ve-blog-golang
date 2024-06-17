package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticleByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleByCategoryLogic {
	return &FindArticleByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章列表
func (l *FindArticleByCategoryLogic) FindArticleByCategory(in *blog.FindArticleByCategoryReq) (*blog.ArticlePageResp, error) {
	result, err := l.svcCtx.ArticleModel.FindALL(l.ctx, "category_id in (?)", in.CategoryIds)
	if err != nil {
		return nil, err
	}

	var list []*blog.Article
	for _, v := range result {
		list = append(list, convert.ConvertArticleModelToPb(v))
	}

	return &blog.ArticlePageResp{
		List: list,
	}, nil
}
