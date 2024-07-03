package article

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleClassifyCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过标签或者id获取文章列表
func NewFindArticleClassifyCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleClassifyCategoryLogic {
	return &FindArticleClassifyCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleClassifyCategoryLogic) FindArticleClassifyCategory(req *types.ArticleClassifyReq) (resp *types.ArticleClassifyResp, err error) {
	cs, err := l.svcCtx.CategoryRpc.FindCategoryList(l.ctx, &blog.PageQuery{
		Page:       1,
		PageSize:   1,
		Sorts:      "id desc",
		Conditions: "category_name = ?",
		Args:       []string{cast.ToString(req.ClassifyName)},
	})
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, v := range cs.List {
		ids = append(ids, v.Id)
	}

	as, err := l.svcCtx.ArticleRpc.FindArticleByCategory(l.ctx, &blog.FindArticleByCategoryReq{
		CategoryIds: ids,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.ArticlePreviewDTO
	for _, v := range as.List {
		m := convert.ConvertArticlePreviewTypes(v)
		list = append(list, m)
	}

	resp = &types.ArticleClassifyResp{}
	resp.ConditionName = req.ClassifyName
	resp.ArticleList = list
	return
}
