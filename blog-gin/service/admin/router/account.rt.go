package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AccountRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewAccountRouter(svcCtx *svctx.ServiceContext) *AccountRouter {
	return &AccountRouter{
		svcCtx: svcCtx,
	}
}

func (s *AccountRouter) Register(r *gin.RouterGroup) {
	// Account
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		handler := controller.NewAccountController(s.svcCtx)
		// 查询用户列表
		group.POST("/account/find_account_list", handler.FindAccountList)
		// 查询在线用户列表
		group.POST("/account/find_account_online_list", handler.FindAccountOnlineList)
		// 修改用户密码
		group.POST("/account/update_account_password", handler.UpdateAccountPassword)
		// 修改用户角色
		group.POST("/account/update_account_roles", handler.UpdateAccountRoles)
		// 修改用户状态
		group.POST("/account/update_account_status", handler.UpdateAccountStatus)
	}
}
