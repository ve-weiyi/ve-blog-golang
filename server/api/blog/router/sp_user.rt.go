package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type UserRouter struct {
	svcCtx *svc.ServiceContext
}

func NewUserRouter(svcCtx *svc.ServiceContext) *UserRouter {
	return &UserRouter{
		svcCtx: svcCtx,
	}
}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewUserController(s.svcCtx)
	{
		loginRouter.GET("/user/get_user_info", handler.GetUserInfo)   // 用户信息
		loginRouter.GET("/user/get_user_menus", handler.GetUserMenus) // 用户菜单
		loginRouter.GET("/user/get_user_apis", handler.GetUserApis)   // 用户资源

		loginRouter.POST("/user/update_user_info", handler.UpdateUserInfo)     // 更新用户信息
		loginRouter.POST("/user/update_user_avatar", handler.UpdateUserAvatar) // 更新用户头像
		loginRouter.POST("/user/update_user_status", handler.UpdateUserStatus) // 更新用户状态
		loginRouter.POST("/user/update_user_roles", handler.UpdateUserRoles)   // 更新用户角色

		// 管理员操作
		loginRouter.POST("/user/find_user_list", handler.FindUserList)              // 获取用户列表
		loginRouter.POST("/user/find_online_user_list", handler.FindOnlineUserList) // 获取在线用户列表
		loginRouter.POST("/user/find_user_areas", handler.FindUserAreas)            // 获取用户地区列表

		loginRouter.POST("/user/login_history", handler.FindUserLoginHistoryList)                 // 用户登录历史
		loginRouter.DELETE("/user/delete_login_history_list", handler.DeleteUserLoginHistoryList) // 批量删除用户登录历史
	}
}
