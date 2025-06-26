package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type UserRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewUserRouter(svcCtx *svctx.ServiceContext) *UserRouter {
	return &UserRouter{
		svcCtx: svcCtx,
	}
}

func (s *UserRouter) Register(r *gin.RouterGroup) {
	// User
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		handler := controller.NewUserController(s.svcCtx)
		// 删除用户绑定第三方平台账号
		group.POST("/user/delete_user_bind_third_party", handler.DeleteUserBindThirdParty)
		// 获取用户接口权限
		group.GET("/user/get_user_apis", handler.GetUserApis)
		// 获取用户信息
		group.GET("/user/get_user_info", handler.GetUserInfo)
		// 查询用户登录历史
		group.POST("/user/get_user_login_history_list", handler.GetUserLoginHistoryList)
		// 获取用户菜单权限
		group.GET("/user/get_user_menus", handler.GetUserMenus)
		// 获取用户角色
		group.GET("/user/get_user_roles", handler.GetUserRoles)
		// 修改用户头像
		group.POST("/user/update_user_avatar", handler.UpdateUserAvatar)
		// 修改用户绑定邮箱
		group.POST("/user/update_user_bind_email", handler.UpdateUserBindEmail)
		// 修改用户绑定手机号
		group.POST("/user/update_user_bind_phone", handler.UpdateUserBindPhone)
		// 修改用户绑定第三方平台账号
		group.POST("/user/update_user_bind_third_party", handler.UpdateUserBindThirdParty)
		// 修改用户信息
		group.POST("/user/update_user_info", handler.UpdateUserInfo)
		// 修改用户密码
		group.POST("/user/update_user_password", handler.UpdateUserPassword)
	}
}
