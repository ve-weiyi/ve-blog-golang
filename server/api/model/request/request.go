package request

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/iputil"
)

// 请求上下文,一般存放请求头参数
type Context struct {
	context.Context `json:"-" header:"-"`
	Token           string `json:"token" header:"token" example:""`
	UID             int    `json:"uid" header:"-" example:""`
	Username        string `json:"username" header:"-" example:""`
	IpAddress       string `json:"ip_address" header:"-" example:""`
	Agent           string `json:"agent" header:"-" example:""`
}

func (s *Context) GetContext() context.Context {
	return s.Context
}

func (s *Context) GetIpSource() string {
	ip := s.IpAddress
	location, err := iputil.GetIpInfoByBaidu(ip)
	if err != nil {
		global.LOG.Println("GetIpInfoByBaidu:", err)
		return "未知ip"
	}
	return location.Location
}

type IdReq struct {
	Id int `json:"id" form:"id" binding:"required"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

type EmptyReq struct{}
