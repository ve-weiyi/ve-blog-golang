package article

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文章相关推荐
func NewGetArticleRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleRecommendLogic {
	return &GetArticleRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleRecommendLogic) GetArticleRecommend(req *types.IdReq) (resp *types.ArticleRecommendResp, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.ArticleRpc.FindArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查询上一篇文章
	last, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, &blog.PageQuery{
		Page:       1,
		PageSize:   1,
		Sorts:      "id desc",
		Conditions: "id < ?",
		Args:       []string{cast.ToString(out.Id)},
	})

	// 查询下一篇文章
	next, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, &blog.PageQuery{
		Page:       1,
		PageSize:   1,
		Sorts:      "id asc",
		Conditions: "id > ?",
		Args:       []string{cast.ToString(out.Id)},
	})

	// 查询推荐文章
	recommend, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, &blog.PageQuery{
		Page:       1,
		PageSize:   5,
		Sorts:      "id desc",
		Conditions: "category_id = ?",
		Args:       []string{cast.ToString(out.CategoryId)},
	})

	// 查询最新文章
	newest, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, &blog.PageQuery{
		Page:     1,
		PageSize: 5,
		Sorts:    "id desc",
	})

	resp = &types.ArticleRecommendResp{}
	resp.ArticleHomeDTO = *convert.ConvertArticleHomeTypes(out)

	for _, v := range last.List {
		resp.LastArticle = convert.ConvertArticlePreviewTypes(v)
	}

	for _, v := range next.List {
		resp.NextArticle = convert.ConvertArticlePreviewTypes(v)
	}

	for _, v := range recommend.List {
		resp.RecommendArticleList = append(resp.RecommendArticleList, convert.ConvertArticlePreviewTypes(v))
	}

	for _, v := range newest.List {
		resp.NewestArticleList = append(resp.NewestArticleList, convert.ConvertArticlePreviewTypes(v))
	}

	return
}
