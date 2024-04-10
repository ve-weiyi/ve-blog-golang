package role

import (
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

func convertIdReq(in *types.IdReq) (out *rolerpc.IdReq) {
	out = &rolerpc.IdReq{
		Id: in.ID,
	}
	return
}

func convertPageQuery(in *types.PageQuery) (out *rolerpc.PageQuery) {
	out = &rolerpc.PageQuery{
		Limit: &rolerpc.PageLimit{
			Page:     in.Limit.Page,
			PageSize: in.Limit.PageSize,
		},
		Sorts:      make([]*account.PageSort, 0),
		Conditions: make([]*account.PageCondition, 0),
	}

	for _, sort := range in.Sorts {
		out.Sorts = append(out.Sorts, &rolerpc.PageSort{
			Field: sort.Field,
			Order: sort.Order,
		})
	}

	for _, condition := range in.Conditions {
		out.Conditions = append(out.Conditions, &account.PageCondition{
			Field:    condition.Field,
			Operator: condition.Operator,
			Value:    cast.ToString(condition.Value),
		})
	}

	return
}

func convertRoleDetailsTypes(in *rolerpc.RoleDetailsDTO) (out *types.RoleDetailsDTO) {
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

func convertRoleTypes(in *rolerpc.Role) (out *types.Role) {
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

func convertRolePb(in *types.Role) (out *rolerpc.Role) {
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
