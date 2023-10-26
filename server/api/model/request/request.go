package request

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
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
	sqlx.PageLimit
	Sorts      []*sqlx.Sort      `json:"sorts" form:"sorts"`           // 排序
	Conditions []*sqlx.Condition `json:"conditions" form:"conditions"` // 使用条件语句查询
}
