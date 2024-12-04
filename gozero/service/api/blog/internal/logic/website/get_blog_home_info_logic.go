package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlogHomeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取博客前台首页信息
func NewGetBlogHomeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogHomeInfoLogic {
	return &GetBlogHomeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBlogHomeInfoLogic) GetBlogHomeInfo(req *types.GetBlogHomeInfoReq) (resp *types.GetBlogHomeInfoResp, err error) {
	analysis, err := l.svcCtx.ArticleRpc.AnalysisArticle(l.ctx, &articlerpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	visit, err := l.svcCtx.WebsiteRpc.GetUserTotalVisit(l.ctx, &websiterpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	pages, err := l.svcCtx.WebsiteRpc.FindPageList(l.ctx, &websiterpc.FindPageListReq{})
	if err != nil {
		return nil, err
	}

	ps := make([]*types.PageDTO, 0)
	for _, v := range pages.List {
		p := &types.PageDTO{
			Id:         v.Id,
			PageName:   v.PageName,
			PageLabel:  v.PageLabel,
			PageCover:  v.PageCover,
			IsCarousel: v.IsCarousel,
		}
		ps = append(ps, p)
	}

	conf, err := l.svcCtx.ConfigRpc.FindConfig(l.ctx, &configrpc.FindConfigReq{
		ConfigKey: "website_config",
	})
	if err != nil {
		return nil, err
	}

	config := types.WebsiteConfigDTO{}
	jsonconv.JsonToAny(conf.ConfigValue, &config)

	resp = &types.GetBlogHomeInfoResp{
		ArticleCount:  analysis.ArticleCount,
		CategoryCount: analysis.CategoryCount,
		TagCount:      analysis.TagCount,
		ViewsCount:    visit.Count,
		WebsiteConfig: config,
		PageList:      ps,
	}

	return resp, nil
}
