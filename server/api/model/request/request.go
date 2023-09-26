package request

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

// 请求上下文,一般存放请求头参数
type Context struct {
	context.Context `json:"-" header:"-"`
	Token           string `json:"token" header:"token" example:""`
	UID             int    `json:"uid" header:"-" example:""`
	Username        string `json:"username" header:"-" example:""`
	IpAddress       string `json:"ip_address" header:"-" example:""`
	IpSource        string `json:"ip_source" header:"-" example:""`
	Agent           string `json:"agent" header:"-" example:""`
}

func (s *Context) GetContext() context.Context {
	return s.Context
}

// PageQuery Paging common input parameter structure
type PageQuery struct {
	Page       int          `json:"page" form:"page"`             // 页码
	PageSize   int          `json:"page_size" form:"page_size"`   // 每页大小
	Sorts      []*Sort      `json:"sorts" form:"sorts"`           // 排序
	Conditions []*Condition `json:"conditions" form:"conditions"` // 使用条件语句查询
}

type Sort struct {
	Field string `json:"field"`                  // 表字段
	Order string `json:"order" enums:"asc,desc"` // 排序规则 asc|desc
}

// 字段，关键字，匹配规则
type Condition struct {
	Field string      `json:"field"`                      // 表字段
	Value interface{} `json:"value"`                      // 值
	Rule  string      `json:"rule" enums:"=,like,in,<,>"` // 比较运算符（Comparison Operators）。规则 =,like,in,<,>
	Flag  string      `json:"flag" enums:"and,or"`        // 逻辑运算符（Logical Operators）。标识 and、or,默认and
}

func (page *PageQuery) Limit() int {
	if page.PageSize == 0 {
		page.PageSize = 10
	}
	return page.PageSize
}

func (page *PageQuery) Offset() int {

	return (page.Page - 1) * page.Limit()
}

func (page *PageQuery) ResetPage() int {
	page.Page = 0
	page.PageSize = 0
	return -1
}

// 排序语句
func (page *PageQuery) OrderClause() string {
	if len(page.Sorts) == 0 {
		return ""
	}

	var query string
	var flag string
	for i, order := range page.Sorts {
		if i == 0 {
			flag = ""
		} else {
			flag = ","
		}
		query += fmt.Sprintf("%s `%s` %s", flag, jsonconv.Camel2Case(order.Field), order.Order)
	}

	return query
}

// 条件语句
func (page *PageQuery) WhereClause() (string, []interface{}) {
	if len(page.Conditions) == 0 {
		return "", nil
	}

	var query string
	var args []interface{}
	var flag string
	for i, condition := range page.Conditions {
		if i == 0 {
			flag = ""
		} else {
			flag = condition.Flag
			if flag == "" {
				flag = "and"
			}
		}

		switch condition.Rule {
		case "like":
			query += fmt.Sprintf("%s %s %s ? ", flag, condition.Field, condition.Rule)
			args = append(args, "%"+condition.Value.(string)+"%")
		case "in":
			query += fmt.Sprintf("%s %s %s (?) ", flag, condition.Field, condition.Rule)
			args = append(args, condition.Value)
		default:
			query += fmt.Sprintf("%s %s %s ? ", flag, condition.Field, condition.Rule)
			args = append(args, condition.Value)
		}
	}

	return query, args
}

func (page *PageQuery) FindCondition(name string) *Condition {
	for _, condition := range page.Conditions {
		if condition.Field == name {
			return condition
		}
	}
	return nil
}
