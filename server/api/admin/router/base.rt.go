package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type BaseRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewBaseRouter(svcCtx *svctx.ServiceContext) *BaseRouter {
	return &BaseRouter{
		svcCtx: svcCtx,
	}
}

func (s *BaseRouter) Register(r *gin.RouterGroup) {
	// Base
	// []
	{
		group := r.Group("/admin_api/v1")

		handler := controller.NewBaseController(s.svcCtx)
		// ping
		group.GET("/ping", handler.Ping)
	}
}
