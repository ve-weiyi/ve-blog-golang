package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type ApiRouter struct {
	svcCtx *svc.RouterContext
}

func NewApiRouter(svcCtx *svc.RouterContext) *ApiRouter {
	return &ApiRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Api 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ApiRouter) InitApiRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.ApiController
	{
		publicRouter.POST("api/create", handler.CreateApi)   // 新建Api
		publicRouter.PUT("api/update", handler.UpdateApi)    // 更新Api
		publicRouter.DELETE("api/delete", handler.DeleteApi) // 删除Api
		publicRouter.POST("api/find", handler.FindApi)       // 查询Api

		publicRouter.DELETE("api/deleteByIds", handler.DeleteApiByIds) // 批量删除Api列表
		publicRouter.POST("api/list", handler.FindApiList)             // 分页查询Api列表

		loginRouter.POST("admin/apis", handler.GetApis) // 获取Api列表
	}
}
