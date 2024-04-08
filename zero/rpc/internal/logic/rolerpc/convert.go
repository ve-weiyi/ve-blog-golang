package rolerpclogic

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/zero/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

func convertRoleModelToPb(in *model.Role) (out *account.Role) {
	out = &account.Role{
		Id:          in.ID,
		RolePid:     in.RolePID,
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

func convertRolePbToModel(in *account.Role) (out *model.Role) {
	out = &model.Role{
		ID:          in.Id,
		RolePID:     in.RolePid,
		RoleDomain:  in.RoleDomain,
		RoleName:    in.RoleName,
		RoleComment: in.RoleComment,
		IsDisable:   in.IsDisable,
		IsDefault:   in.IsDefault,
	}

	return out
}

func convertRoleModelToDetailPb(in *model.Role) (out *account.RoleDetailsDTO) {
	out = &account.RoleDetailsDTO{
		Id:          in.ID,
		RolePid:     in.RolePID,
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

func parsePageQuery(in *account.PageQuery) (limit int, offset int, sorts string, conditions string, params []interface{}) {
	limit, offset = LimitClause(in.Limit)
	sorts = OrderClause(in.Sorts)
	conditions, params = ConditionClause(in.Conditions)
	return limit, offset, sorts, conditions, params
}

// 分页语句
func LimitClause(page *account.PageLimit) (limit int, offset int) {
	var p, s int64

	p = page.Page
	s = page.PageSize

	if p <= 0 {
		p = 1
	}

	if s <= 0 {
		s = 10
	}

	limit = int(s)
	offset = (int(p) - 1) * limit

	return limit, offset
}

// 排序语句
func OrderClause(sorts []*account.PageSort) string {
	if len(sorts) == 0 {
		return ""
	}

	var query string
	var flag string
	for i, order := range sorts {
		if i == 0 {
			flag = ""
		} else {
			flag = ","
		}
		query += fmt.Sprintf("%s `%s` %s", flag, order.Field, order.Order)
	}

	return query
}

// 转换条件语句
func ConditionClause(conditions []*account.PageCondition) (string, []interface{}) {
	if len(conditions) == 0 {
		return "", nil
	}

	var query string
	var args []interface{}
	var flag string
	for i, condition := range conditions {
		if i == 0 {
			flag = ""
		} else {
			flag = condition.Logic
			if flag == "" {
				flag = "and"
			}
		}

		switch condition.Operator {
		case "in":
			query += fmt.Sprintf("%s %s %s (?) ", flag, condition.Field, condition.Operator)
			args = append(args, condition.Value)
		case "like":
			query += fmt.Sprintf("%s %s %s ? ", flag, condition.Field, condition.Operator)
			args = append(args, "%"+condition.Value+"%")
		case "not like":
			query += fmt.Sprintf("%s %s %s ? ", flag, condition.Field, condition.Operator)
			args = append(args, condition.Value+"%")
		default:
			query += fmt.Sprintf("%s %s %s ? ", flag, condition.Field, condition.Operator)
			args = append(args, condition.Value)
		}
	}

	return query, args
}
