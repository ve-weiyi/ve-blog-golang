package svc

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/config"
)

// 注册需要用到的api
type RouterContext struct {
	*controller.AppController
}

func NewRouterContext(cfg *config.Config) *RouterContext {
	ctx := svc.NewControllerContext(cfg)
	ctl := controller.NewController(ctx)
	if ctl == nil {
		panic("ctl cannot be null")
	}

	return &RouterContext{
		AppController: ctl,
	}
}
