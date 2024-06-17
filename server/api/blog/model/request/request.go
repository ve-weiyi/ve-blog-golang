package request

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
)

// 请求上下文,一般存放请求头参数
type Context struct {
	context.Context `json:"-" header:"-"`
	Token           string `json:"token" header:"token" example:""`
	Uid             int64  `json:"uid" header:"-" example:""`
	IpAddress       string `json:"ip_address" header:"-" example:""`
	UserAgent       string `json:"user_agent" header:"-" example:""`
}

func (s *Context) GetContext() context.Context {
	return s.Context
}

func (s *Context) GetIpSource() string {
	ip := s.IpAddress
	location, err := ipx.GetIpInfoByBaidu(ip)
	if err != nil {
		return "未知ip"
	}
	return location.Location
}

type IdReq struct {
	Id int64 `json:"id" form:"id" binding:"required"`
}

type IdsReq struct {
	Ids []int64 `json:"ids" form:"ids" binding:"required"`
}

type EmptyReq struct{}
