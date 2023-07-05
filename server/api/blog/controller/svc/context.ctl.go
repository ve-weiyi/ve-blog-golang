package svc

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
	"github.com/ve-weiyi/ve-admin-store/server/config"
)

// 注册需要用到的rpc
type ControllerContext struct {
	*service.AppService
}

func NewControllerContext(cfg *config.Config) *ControllerContext {
	ctx := svc.NewServiceContext(cfg)
	sv := service.NewService(ctx)
	if sv == nil {
		panic("sv cannot be null")
	}

	return &ControllerContext{
		AppService: sv,
	}
}
