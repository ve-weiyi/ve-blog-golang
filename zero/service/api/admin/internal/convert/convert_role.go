package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
)

func ConvertRolePb(in *types.Role) (out *blogrpc.Role) {
	out = &blogrpc.Role{
		Id:          in.Id,
		ParentId:    in.RolePid,
		RoleDomain:  in.RoleDomain,
		RoleName:    in.RoleName,
		RoleComment: in.RoleComment,
		IsDisable:   in.IsDisable,
		IsDefault:   in.IsDefault,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return
}

func ConvertRoleTypes(in *blogrpc.Role) (out *types.Role) {
	out = &types.Role{
		Id:          in.Id,
		RolePid:     in.ParentId,
		RoleDomain:  in.RoleDomain,
		RoleName:    in.RoleName,
		RoleComment: in.RoleComment,
		IsDisable:   in.IsDisable,
		IsDefault:   in.IsDefault,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return
}

func ConvertRoleDetailsTypes(in *blogrpc.RoleDetails) (out *types.RoleDetails) {
	out = &types.RoleDetails{
		Id:          in.Id,
		RolePid:     in.ParentId,
		RoleDomain:  in.RoleDomain,
		RoleName:    in.RoleName,
		RoleComment: in.RoleComment,
		IsDisable:   in.IsDisable,
		IsDefault:   in.IsDefault,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return
}
