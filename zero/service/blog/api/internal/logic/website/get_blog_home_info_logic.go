package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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

func (l *GetBlogHomeInfoLogic) GetBlogHomeInfo(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.BlogHomeInfo, err error) {
	ac, err := l.svcCtx.ArticleRpc.FindArticleCount(l.ctx, &blog.PageQuery{})
	if err != nil {
		return nil, err
	}

	cc, err := l.svcCtx.CategoryRpc.FindCategoryCount(l.ctx, &blog.PageQuery{})
	if err != nil {
		return nil, err
	}

	tc, err := l.svcCtx.TagRpc.FindTagCount(l.ctx, &blog.PageQuery{})
	if err != nil {
		return nil, err
	}

	in := &blog.FindConfigReq{
		ConfigKey: "website_config",
	}

	out, err := l.svcCtx.ConfigRpc.FindConfig(l.ctx, in)
	if err != nil {
		return nil, err
	}

	config := &types.WebsiteConfig{}
	jsonconv.JsonToObject(out.ConfigValue, &config)

	resp = &types.BlogHomeInfo{
		ArticleCount:  ac.Count,
		CategoryCount: cc.Count,
		TagCount:      tc.Count,
		ViewsCount:    "",
		WebsiteConfig: *config,
		PageList:      make([]*types.PageDTO, 0),
	}

	return resp, nil
}
