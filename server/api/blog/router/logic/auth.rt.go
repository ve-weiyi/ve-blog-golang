package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type AuthRouter struct {
	svcCtx *svc.RouterContext
}

func NewAuthRouter(ctx *svc.RouterContext) *AuthRouter {
	return &AuthRouter{
		svcCtx: ctx,
	}
}

// InitAuthRouter 初始化 Auth 路由信息
func (s *AuthRouter) InitAuthRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	authPublicRouter := publicRouter.Group("")
	//authOperationRouter := loginRouter.Group("")

	var self = s.svcCtx.AppController.AuthController
	{
		authPublicRouter.POST("user/login", self.Login)                         // 登录
		authPublicRouter.GET("user/logout", self.Logout)                        // 登出
		authPublicRouter.POST("user/register", self.Register)                   // 注册
		authPublicRouter.POST("user/register/email", self.RegisterEmail)        // 注册邮件
		authPublicRouter.POST("user/password/reset", self.ResetPassword)        // 重置密码
		authPublicRouter.POST("user/password/forget", self.ForgetPasswordEmail) // 忘记密码-发送邮件
		authPublicRouter.POST("user/oauth/login", self.OauthLogin)              // 使用第三方登录
		authPublicRouter.POST("user/oauth/url", self.GetAuthorizeUrl)           // 获取授权地址
	}
}
