package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type PhotoAlbumRouter struct {
	svcCtx *svc.ServiceContext
}

func NewPhotoAlbumRouter(svcCtx *svc.ServiceContext) *PhotoAlbumRouter {
	return &PhotoAlbumRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 PhotoAlbum 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PhotoAlbumRouter) InitPhotoAlbumRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewPhotoAlbumController(s.svcCtx)
	{
		publicRouter.POST("/photo_album/create_photo_album", handler.CreatePhotoAlbum)            // 新建PhotoAlbum
		publicRouter.PUT("/photo_album/update_photo_album", handler.UpdatePhotoAlbum)             // 更新PhotoAlbum
		publicRouter.DELETE("/photo_album/delete_photo_album", handler.DeletePhotoAlbum)          // 删除PhotoAlbum
		publicRouter.DELETE("/photo_album/delete_photo_album_list", handler.DeletePhotoAlbumList) // 批量删除PhotoAlbum列表

		publicRouter.POST("/photo_album/find_photo_album", handler.FindPhotoAlbum)          // 查询PhotoAlbum
		publicRouter.POST("/photo_album/find_photo_album_list", handler.FindPhotoAlbumList) // 分页查询PhotoAlbum列表

		publicRouter.POST("/photo_album/find_photo_album_details", handler.FindPhotoAlbumDetails)          // 获取PhotoAlbum详情
		publicRouter.POST("/photo_album/find_photo_album_details_list", handler.FindPhotoAlbumDetailsList) // 获取PhotoAlbum详情列表
	}
}
