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
		authPublicRouter.POST("login", self.Login)                  // 登录
		authPublicRouter.GET("logout", self.Logout)                 // 登出
		authPublicRouter.POST("register", self.Register)            // 注册
		authPublicRouter.POST("register/email", self.RegisterEmail) // 注册邮件
		authPublicRouter.POST("oauth/login", self.OauthLogin)       // 使用第三方登录
		authPublicRouter.POST("oauth/url", self.GetAuthorizeUrl)    // 获取授权地址
	}
}
