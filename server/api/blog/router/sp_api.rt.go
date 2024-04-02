package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type ApiRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewApiRouter(svcCtx *svctx.ServiceContext) *ApiRouter {
	return &ApiRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Api 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ApiRouter) InitApiRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewApiController(s.svcCtx)
	{
		loginRouter.POST("/api/create_api", handler.CreateApi)            // 新建Api
		loginRouter.PUT("/api/update_api", handler.UpdateApi)             // 更新Api
		loginRouter.DELETE("/api/delete_api", handler.DeleteApi)          // 删除Api
		loginRouter.DELETE("/api/delete_api_list", handler.DeleteApiList) // 批量删除Api列表

		loginRouter.POST("/api/find_api", handler.FindApi)                         // 查询Api
		loginRouter.POST("/api/find_api_list", handler.FindApiList)                // 分页查询Api列表
		loginRouter.POST("/api/sync_api_list", handler.SyncApiList)                // 同步Api列表
		loginRouter.POST("/api/clean_api_list", handler.CleanApiList)              // 清空Api列表
		loginRouter.POST("/api/find_api_details_list", handler.FindApiDetailsList) // 获取Api列表
	}
}
