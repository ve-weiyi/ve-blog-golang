package permissionrpclogic

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
)

func convertRoleIn(in *permissionrpc.RoleNewReq) (out *model.Role) {
	out = &model.Role{
		Id:          in.Id,
		ParentId:    in.ParentId,
		RoleDomain:  in.RoleDomain,
		RoleName:    in.RoleName,
		RoleComment: in.RoleComment,
		IsDisable:   in.IsDisable,
		IsDefault:   in.IsDefault,
	}

	return out
}

func convertRoleOut(in *model.Role) (out *permissionrpc.RoleDetails) {
	out = &permissionrpc.RoleDetails{
		Id:          in.Id,
		ParentId:    in.ParentId,
		RoleDomain:  in.RoleDomain,
		RoleName:    in.RoleName,
		RoleComment: in.RoleComment,
		IsDisable:   in.IsDisable,
		IsDefault:   in.IsDefault,
		CreatedAt:   in.CreatedAt.Unix(),
		UpdatedAt:   in.UpdatedAt.Unix(),
	}

	return out
}
