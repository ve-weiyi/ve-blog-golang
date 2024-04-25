package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleResourcesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindRoleResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleResourcesLogic {
	return &FindRoleResourcesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleResourcesLogic) FindRoleResources(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.RoleResourcesResp, err error) {
	in := convert.ConvertIdReq(req)
	out, err := l.svcCtx.RoleRpc.FindRoleResources(l.ctx, in)
	if err != nil {
		return
	}

	resp = &types.RoleResourcesResp{}
	resp.RoleId = out.RoleId
	resp.ApiIds = out.ApiIds
	resp.MenuIds = out.MenuIds

	return
}
