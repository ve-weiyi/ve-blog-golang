package logic

import (
	"github.com/gin-gonic/gin"
)

func (s *AuthRouter) InitRbacRouter(Router *gin.RouterGroup) {
	//authRouter := Router.Group("rbac")
	//authOperationRouter := Router.Group("rbac")
	//authOperationRouter.Use(middleware.JwtToken())
	//authOperationRouter.Use(middleware.OperationRecord())
	//
	//var self = s.svcCtx.AppController.RbacController
	//{
	//	authOperationRouter.POST("updateCasbin", self.UpdateCasbin)                             // 注册
	//	authOperationRouter.POST("getPolicyPathByAuthorityId", self.GetPolicyPathByAuthorityId) // 登录
	//}
}
