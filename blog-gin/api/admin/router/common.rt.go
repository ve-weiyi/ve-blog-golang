package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/admin/handler"
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
		group := r.Group("/admin-api/v1")

		h := handler.NewCommonController(s.svcCtx)
		// ping
		group.GET("/ping", h.Ping)
	}
}
