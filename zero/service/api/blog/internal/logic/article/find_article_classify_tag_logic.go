package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleClassifyTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过标签获取文章列表
func NewFindArticleClassifyTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleClassifyTagLogic {
	return &FindArticleClassifyTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleClassifyTagLogic) FindArticleClassifyTag(req *types.ArticleClassifyQueryReq) (resp *types.PageResp, err error) {
	in := &articlerpc.FindArticlesByTagReq{
		TagName: req.ClassifyName,
	}
	out, err := l.svcCtx.ArticleRpc.FindArticlesByTag(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ArticleHome
	// 转换数据
	for _, v := range out.List {
		m := convert.ConvertArticleHomeTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = out.Total
	resp.List = list
	return
}
