package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleListLogic {
	return &FindRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleListLogic) FindRoleList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.RoleRpc.FindRoleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.RoleDetails
	for _, role := range out.List {
		list = append(list, convert.ConvertRoleDetailsTypes(role))
	}

	resp = &types.PageResp{}
	resp.Page = in.Limit.Page
	resp.PageSize = in.Limit.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
