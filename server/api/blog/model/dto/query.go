package dto

import (
	"fmt"
)

// PageQuery Paging common input parameter structure
type PageQuery struct {
	Limit      PageLimit
	Sorts      []*PageSort      `json:"sorts" form:"sorts"`           // 排序
	Conditions []*PageCondition `json:"conditions" form:"conditions"` // 使用条件语句查询
}

func (s *PageQuery) PageClause() (int, int) {
	return LimitClause(s.Limit)
}

func (s *PageQuery) OrderClause() string {
	return OrderClause(s.Sorts)
}

func (s *PageQuery) ConditionClause() (string, []interface{}) {
	return ConditionClause(s.Conditions)
}

// 分页
type PageLimit struct {
	Page     int64 `json:"page" form:"page"`           // 页码
	PageSize int64 `json:"page_size" form:"page_size"` // 每页大小
}

// 排序语句
type PageSort struct {
	Field string `json:"field"`                  // 表字段
	Order string `json:"order" enums:"asc,desc"` // 排序规则 asc|desc
}

// 查询条件
type PageCondition struct {
	Field    string      `json:"field"`                          // 表字段
	Value    interface{} `json:"value"`                          // 值
	Operator string      `json:"operator" enums:"=,like,in,<,>"` // 比较运算符（Comparison Operators）。规则 =,like,in,<,>
	Logic    string      `json:"logic" enums:"and,or"`           // 逻辑运算符（Logical Operators）。标识 and、or,默认and
}

// 分页语句
func LimitClause(page PageLimit) (limit int, offset int) {

	var p, s int

	p = int(page.Page)
	s = int(page.PageSize)

	if p <= 0 {
		p = 1
	}

	if s <= 0 {
		s = 10
	}

	limit = s
	offset = (p - 1) * s

	return limit, offset
}

// 排序语句
func OrderClause(sorts []*PageSort) string {
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
func ConditionClause(conditions []*PageCondition) (string, []interface{}) {
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
			args = append(args, "%"+condition.Value.(string)+"%")
		case "not like":
			query += fmt.Sprintf("%s %s %s ? ", flag, condition.Field, condition.Operator)
			args = append(args, condition.Value.(string)+"%")
		default:
			query += fmt.Sprintf("%s %s %s ? ", flag, condition.Field, condition.Operator)
			args = append(args, condition.Value)
		}
	}

	return query, args
}
