package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type PageRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewPageRouter(svcCtx *svctx.ServiceContext) *PageRouter {
	return &PageRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Page 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PageRouter) InitPageRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewPageController(s.svcCtx)
	{
		publicRouter.POST("/page/create_page", handler.CreatePage)            // 新建Page
		publicRouter.PUT("/page/update_page", handler.UpdatePage)             // 更新Page
		publicRouter.DELETE("/page/delete_page", handler.DeletePage)          // 删除Page
		publicRouter.DELETE("/page/delete_page_list", handler.DeletePageList) // 批量删除Page列表

		publicRouter.POST("/page/find_page", handler.FindPage)          // 查询Page
		publicRouter.POST("/page/find_page_list", handler.FindPageList) // 分页查询Page列表
	}
}
