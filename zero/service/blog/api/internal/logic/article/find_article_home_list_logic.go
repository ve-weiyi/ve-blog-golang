package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleHomeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取文章列表
func NewFindArticleHomeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleHomeListLogic {
	return &FindArticleHomeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleHomeListLogic) FindArticleHomeList(reqCtx *types.RestHeader, req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.ArticleRpc.FindArticleCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ArticleHome
	for _, v := range out.List {
		var category string
		ct, err := l.svcCtx.CategoryRpc.FindCategory(l.ctx, &blog.IdReq{Id: v.CategoryId})
		if ct != nil {
			category = ct.CategoryName
		}

		ts, err := l.svcCtx.TagRpc.FindTagListByArticleId(l.ctx, &blog.IdReq{Id: v.Id})
		if err != nil {
			return nil, err
		}

		var tags []string
		for _, tag := range ts.List {
			tags = append(tags, tag.TagName)
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
