package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/repository/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

func ConvertRoleModelToPb(in *model.Role) (out *account.Role) {
	out = &account.Role{
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

func ConvertRolePbToModel(in *account.Role) (out *model.Role) {
	out = &model.Role{
		Id:          in.Id,
		ParentId:    in.ParentId,
		RoleDomain:  in.RoleDomain,
		RoleName:    in.RoleName,
		RoleComment: in.RoleComment,
		IsDisable:   in.IsDisable,
		IsDefault:   in.IsDefault,
		CreatedAt:   time.Unix(in.CreatedAt, 0),
		UpdatedAt:   time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertRoleModelToDetailPb(in *model.Role) (out *account.RoleDetailsDTO) {
	out = &account.RoleDetailsDTO{
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
