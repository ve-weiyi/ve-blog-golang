package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type CategoryRouter struct {
	svcCtx *svc.ServiceContext
}

func NewCategoryRouter(svcCtx *svc.ServiceContext) *CategoryRouter {
	return &CategoryRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Category 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CategoryRouter) InitCategoryRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewCategoryController(s.svcCtx)
	{
		loginRouter.POST("/category/create_category", handler.CreateCategory)            // 新建Category
		loginRouter.PUT("/category/update_category", handler.UpdateCategory)             // 更新Category
		loginRouter.DELETE("/category/delete_category", handler.DeleteCategory)          // 删除Category
		loginRouter.DELETE("/category/delete_category_list", handler.DeleteCategoryList) // 批量删除Category列表

		publicRouter.POST("/category/find_category", handler.FindCategory)          // 查询Category
		publicRouter.POST("/category/find_category_list", handler.FindCategoryList) // 分页查询Category列表
	}
}
