package article

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章详情
func NewGetArticleDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleDetailsLogic {
	return &GetArticleDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleDetailsLogic) GetArticleDetails(req *types.IdReq) (resp *types.ArticleDeatils, err error) {
	in := &blogrpc.IdReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.ArticleRpc.GetArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查询上一篇文章
	last, err := l.svcCtx.ArticleRpc.FindArticlePublicList(l.ctx, &blogrpc.FindArticleListReq{
		Page:       1,
		PageSize:   1,
		Sorts:      "id desc",
		Conditions: "id < ?",
		Args:       []string{cast.ToString(out.Id)},
	})

	// 查询下一篇文章
	next, err := l.svcCtx.ArticleRpc.FindArticlePublicList(l.ctx, &blogrpc.FindArticleListReq{
		Page:       1,
		PageSize:   1,
		Sorts:      "id asc",
		Conditions: "id > ?",
		Args:       []string{cast.ToString(out.Id)},
	})

	// 查询推荐文章
	recommend, err := l.svcCtx.ArticleRpc.FindArticlePublicList(l.ctx, &blogrpc.FindArticleListReq{
		Page:       1,
		PageSize:   5,
		Sorts:      "id desc",
		Conditions: "category_id = ?",
		Args:       []string{cast.ToString(out.CategoryId)},
	})

	// 查询最新文章
	newest, err := l.svcCtx.ArticleRpc.FindArticlePublicList(l.ctx, &blogrpc.FindArticleListReq{
		Page:     1,
		PageSize: 5,
		Sorts:    "id desc",
	})

	resp = &types.ArticleDeatils{}
	resp.ArticleHome = *convert.ConvertArticleHomeTypes(out)

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
