package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/pagerpc"

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

func (l *FindPageListLogic) FindPageList(req *types.PageQueryReq) (resp *types.PageResp, err error) {
	in := &pagerpc.FindPageListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
	}

	out, err := l.svcCtx.PageRpc.FindPageList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.PageBackDTO
	for _, v := range out.List {
		m := ConvertPageTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
