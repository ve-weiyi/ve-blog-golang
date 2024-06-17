package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取首页推荐文章列表
func NewFindArticleRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleRecommendLogic {
	return &FindArticleRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleRecommendLogic) FindArticleRecommend(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.ArticleRpc.FindArticleCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var aids []int64
	var cids []int64
	for _, v := range out.List {
		aids = append(aids, v.Id)
		cids = append(cids, v.CategoryId)
	}

	// 查询分类
	categories, err := l.svcCtx.CategoryRpc.FindCategoryListByIds(l.ctx, &blogrpc.IdsReq{Ids: cids})
	if err != nil {
		return nil, err
	}
	// 查询标签
	tms, err := l.svcCtx.TagRpc.FindTagMapByArticleIds(l.ctx, &blogrpc.IdsReq{Ids: aids})
	if err != nil {
		return nil, err
	}
	// 转换数据
	var list []*types.ArticleHome
	for _, v := range out.List {
		var category string
		for _, c := range categories.List {
			if v.CategoryId == c.Id {
				category = c.CategoryName
			}
		}

		var tags []string
		ts := tms.TagMapList[v.Id].List
		if ts != nil {
			for _, t := range ts {
				tags = append(tags, t.TagName)
			}
		}

		m := convert.ConvertArticleHomeTypes(v)
		m.CategoryName = category
		m.TagNameList = tags
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = total.Count
	resp.List = list
	return
}
