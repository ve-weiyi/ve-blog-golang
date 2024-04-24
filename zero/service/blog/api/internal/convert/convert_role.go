package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/rolerpc"
)

func ConvertRoleTypes(in *rolerpc.Role) (out *types.Role) {
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

func ConvertRolePb(in *types.Role) (out *rolerpc.Role) {
	out = &rolerpc.Role{
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

func ConvertRoleDetailsTypes(in *rolerpc.RoleDetails) (out *types.RoleDetails) {
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
