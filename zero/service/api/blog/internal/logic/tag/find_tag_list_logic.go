package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取标签列表
func NewFindTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagListLogic {
	return &FindTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTagListLogic) FindTagList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.TagRpc.FindTagList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.TagRpc.FindTagCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.Tag
	for _, v := range out.List {
		row, _ := l.svcCtx.TagRpc.FindTagArticleCount(l.ctx, &blogrpc.FindTagArticleCountReq{
			TagId: v.Id,
		})

		m := convert.ConvertTagTypes(v)
		m.ArticleCount = row.Count
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
