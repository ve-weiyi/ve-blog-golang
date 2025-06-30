package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
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
	// [TerminalToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.TerminalToken)

		handler := controller.NewUploadController(s.svcCtx)
		// 删除文件列表
		group.DELETE("/upload/deletes_upload_file", handler.DeletesUploadFile)
		// 获取文件列表
		group.POST("/upload/list_upload_file", handler.ListUploadFile)
		// 上传文件列表
		group.POST("/upload/multi_upload_file", handler.MultiUploadFile)
		// 上传文件
		group.POST("/upload/upload_file", handler.UploadFile)
	}
}
