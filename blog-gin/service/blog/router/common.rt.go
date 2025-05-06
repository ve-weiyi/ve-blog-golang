package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CommonRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewCommonRouter(svcCtx *svctx.ServiceContext) *CommonRouter {
	return &CommonRouter{
		svcCtx: svcCtx,
	}
}

func (s *CommonRouter) Register(r *gin.RouterGroup) {
	// Common
	// []
	{
		group := r.Group("/blog-api/v1")

		handler := controller.NewCommonController(s.svcCtx)
		// ping
		group.GET("/ping", handler.Ping)
	}
}
