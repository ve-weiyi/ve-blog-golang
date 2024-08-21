package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/global"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticlesByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticlesByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticlesByCategoryLogic {
	return &FindArticlesByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章列表
func (l *FindArticlesByCategoryLogic) FindArticlesByCategory(in *blog.FindArticlesByCategoryReq) (*blog.FindArticleListResp, error) {
	category, err := l.svcCtx.CategoryModel.FindOneByCategoryName(l.ctx, in.CategoryName)
	if err != nil {
		return nil, err
	}

	// 查询文章信息
	records, err := l.svcCtx.ArticleModel.FindALL(l.ctx, "category_id = ? and status = ?", category.Id, global.ArticleStatusPublic)
	if err != nil {
		return nil, err
	}

	acm, err := findCategoryGroupArticle(l.ctx, l.svcCtx, records)
	if err != nil {
		return nil, err

	}

	atm, err := findTagGroupArticle(l.ctx, l.svcCtx, records)
	if err != nil {
		return nil, err
	}

	var list []*blog.ArticleDetails
	for _, v := range records {
		list = append(list, convertArticleOut(v, acm, atm))
	}

	return &blog.FindArticleListResp{
		List: list,
	}, nil
}
