package website

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminHomeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取后台首页信息
func NewGetAdminHomeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminHomeInfoLogic {
	return &GetAdminHomeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminHomeInfoLogic) GetAdminHomeInfo(req *types.EmptyReq) (resp *types.AdminHomeInfo, err error) {

	in := &blogrpc.PageQuery{}

	// 查询文章
	articles, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查询分类
	categories, err := l.svcCtx.CategoryRpc.FindCategoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查询标签
	tags, err := l.svcCtx.TagRpc.FindTagList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查询消息
	msgCount, err := l.svcCtx.RemarkRpc.FindRemarkCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查询用户数量
	userCount, err := l.svcCtx.UserRpc.FindUserList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	views, err := l.svcCtx.BlogRpc.GetUserVisitList(l.ctx, &blogrpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	var cs []*types.CategoryDTO
	var ts []*types.TagDTO
	var ars []*types.ArticleViewRankDTO
	var ass []*types.ArticleStatisticsDTO
	var uvs []*types.UniqueViewDTO

	mad := make(map[string]int64)

	for _, v := range categories.List {
		cs = append(cs, convert.ConvertHomeCategoryTypes(v))
	}

	for _, v := range tags.List {
		ts = append(ts, convert.ConvertHomeTagTypes(v))
	}

	for _, v := range articles.List {
		ars = append(ars, convert.ConvertHomeArticleRankTypes(v))
		mad[time.Unix(v.CreatedAt, 0).Format(time.DateOnly)]++
	}

	for _, v := range views.List {
		uvs = append(uvs, convert.ConvertHomeViewTypes(v))
	}

	for k, v := range mad {
		as := &types.ArticleStatisticsDTO{
			Date:  k,
			Count: v,
		}

		ass = append(ass, as)
	}

	resp = &types.AdminHomeInfo{
		ViewsCount:            0,
		MessageCount:          msgCount.Count,
		UserCount:             userCount.Total,
		ArticleCount:          int64(len(articles.List)),
		CategoryList:          cs,
		TagList:               ts,
		ArticleViewRankList:   ars,
		ArticleStatisticsList: ass,
		UniqueViewList:        uvs,
	}

	return
}
