package me

import (
	"context"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/infra/biz/bizheader"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/permissionservice"
)

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户角色
func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles(req *types.EmptyReq) (resp *types.GetUserRolesResp, err error) {
	userId := cast.ToString(l.ctx.Value(bizheader.HeaderUid))

	out, err := l.svcCtx.PermissionService.GetUserRoles(l.ctx, &permissionservice.GetUserRolesRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.UserRole
	for _, v := range out.List {
		list = append(list, &types.UserRole{
			Id:          v.Id,
			ParentId:    v.ParentId,
			RoleKey:     v.RoleKey,
			RoleLabel:   v.RoleLabel,
			RoleComment: v.RoleComment,
		})
	}

	return &types.GetUserRolesResp{List: list}, nil
}
