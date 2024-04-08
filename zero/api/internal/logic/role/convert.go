package role

import (
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
			Value:    condition.Value,
		})
	}

	return
}

func convertRoleApi(in *rolerpc.Role) (out *types.Role) {
	out = &types.Role{
		ID:          in.Id,
		RolePID:     in.RolePid,
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

func convertRoleRpc(in *types.Role) (out *rolerpc.Role) {
	out = &rolerpc.Role{
		Id:          in.ID,
		RolePid:     in.RolePID,
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
