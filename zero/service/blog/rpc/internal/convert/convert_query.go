package convert

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ParsePageQuery(in *blog.PageQuery) (limit int, offset int, sorts string, conditions string, params []interface{}) {
	limit, offset = LimitClause(in)
	//sorts = OrderClause(in.Sorts)
	//conditions, params = ConditionClause(in.Conditions)

	sorts = in.Sorts
	conditions = in.Conditions

	for _, v := range in.Args {
		params = append(params, v)
	}

	return limit, offset, sorts, conditions, params
}

// 分页语句
func LimitClause(page *blog.PageQuery) (limit int, offset int) {
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
func OrderClause(sorts []*blog.PageSort) string {
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
func ConditionClause(conditions []*blog.PageCondition) (string, []interface{}) {
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
