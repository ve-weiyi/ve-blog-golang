package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"
)

func ConvertRoleTypes(in *rolerpc.Role) (out *types.Role) {
	out = &types.Role{
		ID:          in.Id,
		RolePID:     in.ParentId,
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
		Id:          in.ID,
		ParentId:    in.RolePID,
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

func ConvertRoleDetailsTypes(in *rolerpc.RoleDetailsDTO) (out *types.RoleDetailsDTO) {
	out = &types.RoleDetailsDTO{
		ID:          in.Id,
		RolePID:     in.ParentId,
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
