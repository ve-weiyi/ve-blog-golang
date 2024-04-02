package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type PhotoRouter struct {
	svcCtx *svc.RouterContext
}

func NewPhotoRouter(svcCtx *svc.RouterContext) *PhotoRouter {
	return &PhotoRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Photo 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PhotoRouter) InitPhotoRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.PhotoController
	{
		publicRouter.POST("/photo/create_photo", handler.CreatePhoto)   // 新建Photo
		publicRouter.PUT("/photo/update_photo", handler.UpdatePhoto)    // 更新Photo
		publicRouter.DELETE("/photo/delete_photo", handler.DeletePhoto) // 删除Photo
		publicRouter.POST("/photo/find_photo", handler.FindPhoto)       // 查询Photo

		publicRouter.DELETE("/photo/delete_photo_list", handler.DeletePhotoList) // 批量删除Photo列表
		publicRouter.POST("/photo/find_photo_list", handler.FindPhotoList)       // 分页查询Photo列表
	}
}
