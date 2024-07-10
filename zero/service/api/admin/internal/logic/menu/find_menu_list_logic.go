package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuListLogic {
	return &FindMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMenuListLogic) FindMenuList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.MenuRpc.FindMenuList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.MenuDetails
	for _, role := range out.List {
		list = append(list, convert.ConvertMenuDetailsTypes(role))
	}

	resp = &types.PageResp{}
	resp.Page = 1
	resp.PageSize = int64(len(list))
	resp.Total = int64(len(list))
	resp.List = list
	return resp, nil
}
