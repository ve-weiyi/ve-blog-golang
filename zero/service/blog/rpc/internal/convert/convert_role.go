package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertRoleModelToPb(in *model.Role) (out *blog.Role) {
	out = &blog.Role{
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

func ConvertRolePbToModel(in *blog.Role) (out *model.Role) {
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

func ConvertRoleModelToDetailPb(in *model.Role) (out *blog.RoleDetails) {
	out = &blog.RoleDetails{
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
