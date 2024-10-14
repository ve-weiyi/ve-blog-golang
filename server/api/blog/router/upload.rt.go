package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type UploadRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadRouter(svcCtx *svctx.ServiceContext) *UploadRouter {
	return &UploadRouter{
		svcCtx: svcCtx,
	}
}

func (s *UploadRouter) Register(r *gin.RouterGroup) {
	// Upload
	// [SignToken]
	{
		group := r.Group("/api/v1")
		group.Use(s.svcCtx.MiddlewareSignToken)

		handler := controller.NewUploadController(s.svcCtx)
		// 上传文件
		group.POST("/upload/upload_file", handler.UploadFile)
	}
}
