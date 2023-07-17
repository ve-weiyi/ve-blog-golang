package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
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

	var handler = s.svcCtx.AppController.PhotoController
	{
		publicRouter.POST("photo/create", handler.CreatePhoto)   // 新建Photo
		publicRouter.PUT("photo/update", handler.UpdatePhoto)    // 更新Photo
		publicRouter.DELETE("photo/delete", handler.DeletePhoto) // 删除Photo
		publicRouter.POST("photo/find", handler.FindPhoto)       // 查询Photo

		publicRouter.DELETE("photo/deleteByIds", handler.DeletePhotoByIds) // 批量删除Photo列表
		publicRouter.POST("photo/list", handler.FindPhotoList)             // 分页查询Photo列表
	}
}
