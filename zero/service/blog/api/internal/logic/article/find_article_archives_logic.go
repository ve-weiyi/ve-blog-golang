package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleArchivesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文章归档(时间轴)
func NewFindArticleArchivesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleArchivesLogic {
	return &FindArticleArchivesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleArchivesLogic) FindArticleArchives(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	in.Sorts = "id desc"
	in.Conditions = "status = ?"
	in.Args = []string{"1"}
	out, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.ArticleRpc.FindArticleCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ArticlePreviewDTO
	for _, v := range out.List {
		m := convert.ConvertArticlePreviewTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = total.Count
	resp.List = list
	return
}
