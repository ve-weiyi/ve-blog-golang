package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type OperationLogRouter struct {
	svcCtx *svc.ServiceContext
}

func NewOperationLogRouter(svcCtx *svc.ServiceContext) *OperationLogRouter {
	return &OperationLogRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 OperationLog 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *OperationLogRouter) InitOperationLogRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewOperationLogController(s.svcCtx)
	{
		loginRouter.POST("/operation_log/create_operation_log", handler.CreateOperationLog)            // 新建OperationLog
		loginRouter.PUT("/operation_log/update_operation_log", handler.UpdateOperationLog)             // 更新OperationLog
		loginRouter.DELETE("/operation_log/delete_operation_log", handler.DeleteOperationLog)          // 删除OperationLog
		loginRouter.DELETE("/operation_log/delete_operation_log_list", handler.DeleteOperationLogList) // 批量删除OperationLog列表

		publicRouter.POST("/operation_log/find_operation_log", handler.FindOperationLog)         // 查询OperationLog
		loginRouter.POST("/operation_log/find_operation_log_list", handler.FindOperationLogList) // 分页查询OperationLog列表
	}
}
