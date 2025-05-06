package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type FileRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewFileRouter(svcCtx *svctx.ServiceContext) *FileRouter {
	return &FileRouter{
		svcCtx: svcCtx,
	}
}

func (s *FileRouter) Register(r *gin.RouterGroup) {
	// File
	// [TimeToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.MiddlewareTimeToken)

		handler := controller.NewFileController(s.svcCtx)
		// 上传文件列表
		group.POST("/file/multi_upload_file", handler.MultiUploadFile)
		// 上传文件
		group.POST("/file/upload_file", handler.UploadFile)
	}
}
