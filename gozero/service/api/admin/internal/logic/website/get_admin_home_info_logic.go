package website

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/websiterpc"

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

	// 查询用户总浏览量
	view, err := l.svcCtx.WebsiteRpc.GetUserTotalVisit(l.ctx, &websiterpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	// 查询用户数量
	users, err := l.svcCtx.AccountRpc.AnalysisUser(l.ctx, &accountrpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	// 查询文章
	articles, err := l.svcCtx.ArticleRpc.AnalysisArticle(l.ctx, &articlerpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	// 查询消息数量
	messages, err := l.svcCtx.MessageRpc.AnalysisMessage(l.ctx, &messagerpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	ars := make([]*types.ArticleViewDTO, 0)
	for _, v := range articles.ArticleRankList {
		m := &types.ArticleViewDTO{
			Id:           v.Id,
			ArticleTitle: v.ArticleTitle,
			ViewCount:    v.ViewCount,
		}

		ars = append(ars, m)
	}

	tvs := make([]*types.TagDTO, 0)
	for _, v := range articles.TagList {
		m := &types.TagDTO{
			Id:           v.Id,
			TagName:      v.TagName,
			ArticleCount: v.ArticleCount,
		}

		tvs = append(tvs, m)
	}

	cvs := make([]*types.CategoryDTO, 0)
	for _, v := range articles.CategoryList {
		m := &types.CategoryDTO{
			Id:           v.Id,
			CategoryName: v.CategoryName,
			ArticleCount: v.ArticleCount,
		}

		cvs = append(cvs, m)
	}

	archives, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, &articlerpc.FindArticleListReq{})
	if err != nil {
		return nil, err
	}

	asm := make(map[string]int64)
	for _, v := range archives.List {
		date := time.Unix(v.CreatedAt, 0).Format(time.DateOnly)
		if _, ok := asm[date]; ok {
			asm[date]++
		} else {
			asm[date] = 1
		}
	}

	ass := make([]*types.ArticleStatisticsDTO, 0)
	for k, v := range asm {
		m := &types.ArticleStatisticsDTO{
			Date:  k,
			Count: v,
		}

		ass = append(ass, m)
	}

	// 查询用户浏览量
	daily, err := l.svcCtx.WebsiteRpc.GetUserDailyVisit(l.ctx, &websiterpc.EmptyReq{})
	if err != nil {
		return nil, err
	}
	uvs := make([]*types.UserVisitDTO, 0)
	for _, v := range daily.List {
		m := &types.UserVisitDTO{
			Date:  v.Date,
			Count: v.ViewCount,
		}

		uvs = append(uvs, m)
	}

	resp = &types.AdminHomeInfo{
		ViewCount:         view.Count,
		UserCount:         users.UserCount,
		ArticleCount:      articles.ArticleCount,
		RemarkCount:       messages.RemarkCount,
		CategoryList:      cvs,
		TagList:           tvs,
		ArticleViewRanks:  ars,
		ArticleStatistics: ass,
		UserVisitDaliy:    uvs,
	}

	return
}
