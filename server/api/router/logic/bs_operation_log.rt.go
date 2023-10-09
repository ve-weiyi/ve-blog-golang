package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type OperationLogRouter struct {
	svcCtx *svc.RouterContext
}

func NewOperationLogRouter(svcCtx *svc.RouterContext) *OperationLogRouter {
	return &OperationLogRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 OperationLog 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *OperationLogRouter) InitOperationLogBasicRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.OperationLogController
	{
		publicRouter.POST("operation_log", handler.CreateOperationLog)       // 新建OperationLog
		publicRouter.PUT("operation_log", handler.UpdateOperationLog)        // 更新OperationLog
		publicRouter.DELETE("operation_log/:id", handler.DeleteOperationLog) // 删除OperationLog
		publicRouter.GET("operation_log/:id", handler.FindOperationLog)      // 查询OperationLog

		publicRouter.DELETE("operation_log/batch_delete", handler.DeleteOperationLogByIds) // 批量删除OperationLog列表
		publicRouter.POST("operation_log/list", handler.FindOperationLogList)              // 分页查询OperationLog列表
	}
}
