package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AlbumRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewAlbumRouter(svcCtx *svctx.ServiceContext) *AlbumRouter {
	return &AlbumRouter{
		svcCtx: svcCtx,
	}
}

func (s *AlbumRouter) Register(r *gin.RouterGroup) {
	// Album
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewAlbumController(s.svcCtx)
		// 创建相册
		group.POST("/album/add_album", handler.AddAlbum)
		// 删除相册
		group.DELETE("/album/deletes_album", handler.DeletesAlbum)
		// 分页获取相册列表
		group.POST("/album/find_album_list", handler.FindAlbumList)
		// 查询相册
		group.POST("/album/get_album", handler.GetAlbum)
		// 预删除相册
		group.POST("/album/pre_delete_album", handler.PreDeleteAlbum)
		// 更新相册
		group.PUT("/album/update_album", handler.UpdateAlbum)
	}
}
