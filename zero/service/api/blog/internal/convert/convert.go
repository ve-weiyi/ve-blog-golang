package convert

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func EmptyReq() (out *blog.EmptyReq) {
	out = &blog.EmptyReq{}
	return
}

func ConvertIdReq(in *types.IdReq) (out *blog.IdReq) {
	out = &blog.IdReq{
		Id: in.Id,
	}
	return
}

func ConvertIdsReq(in *types.IdsReq) (out *blog.IdsReq) {
	out = &blog.IdsReq{
		Ids: in.Ids,
	}
	return
}

func ConvertPageQuery(in *types.PageQuery) (out *blog.PageQuery) {
	out = &blog.PageQuery{}
	out.Page = in.Page
	out.PageSize = in.PageSize
	out.Sorts = OrderClause(in.Sorts)
	out.Conditions, out.Args = ConditionClause(in.Conditions)

	//for _, sort := range in.Sorts {
	//	out.Sorts = append(out.Sorts, &blog.PageSort{
	//		Field: sort.Field,
	//		Order: sort.Order,
	//	})
	//}
	//
	//for _, condition := range in.Conditions {
	//	out.Conditions = append(out.Conditions, &blog.PageCondition{
	//		Field:    condition.Field,
	//		Operator: condition.Operator,
	//		Value:    cast.ToString(condition.Value),
	//	})
	//}

	return
}

// 分页语句
func LimitClause(page *types.PageQuery) (limit int, offset int) {
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
func OrderClause(sorts []*types.PageSort) string {
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
func ConditionClause(conditions []*types.PageCondition) (query string, args []string) {
	if len(conditions) == 0 {
		return "", nil
	}

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
