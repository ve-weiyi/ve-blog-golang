package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type ApiRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewApiRouter(svcCtx *svctx.ServiceContext) *ApiRouter {
	return &ApiRouter{
		svcCtx: svcCtx,
	}
}

func (s *ApiRouter) Register(r *gin.RouterGroup) {
	// Api
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewApiController(s.svcCtx)
		// 创建api路由
		group.POST("/api/add_api", handler.AddApi)
		// 清空接口列表
		group.POST("/api/clean_api_list", handler.CleanApiList)
		// 删除api路由
		group.DELETE("/api/deletes_api", handler.DeletesApi)
		// 分页获取api路由列表
		group.POST("/api/find_api_list", handler.FindApiList)
		// 同步api列表
		group.POST("/api/sync_api_list", handler.SyncApiList)
		// 更新api路由
		group.PUT("/api/update_api", handler.UpdateApi)
	}
}
