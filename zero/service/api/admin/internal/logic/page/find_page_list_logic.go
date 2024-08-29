package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取页面列表
func NewFindPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageListLogic {
	return &FindPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPageListLogic) FindPageList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := ConvertPageQuery(req)
	out, err := l.svcCtx.PageRpc.FindPageList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.PageRpc.FindPageCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.Page
	for _, v := range out.List {
		list = append(list, ConvertPageTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
