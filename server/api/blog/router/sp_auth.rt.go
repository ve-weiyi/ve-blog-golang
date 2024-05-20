package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type AuthRouter struct {
	svcCtx *svc.ServiceContext
}

func NewAuthRouter(svcCtx *svc.ServiceContext) *AuthRouter {
	return &AuthRouter{
		svcCtx: svcCtx,
	}
}

// InitAuthRouter 初始化 Auth 路由信息
func (s *AuthRouter) InitAuthRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	authPublicRouter := publicRouter.Group("/")

	var handler = controller.NewAuthController(s.svcCtx)
	{
		authPublicRouter.POST("/login", handler.Login)                    // 登录
		authPublicRouter.POST("/logout", handler.Logout)                  // 登出
		authPublicRouter.POST("/register", handler.Register)              // 注册
		authPublicRouter.POST("/register/email", handler.RegisterEmail)   // 注册邮件
		authPublicRouter.POST("/oauth/login", handler.OauthLogin)         // 使用第三方登录
		authPublicRouter.POST("/oauth/url", handler.GetOauthAuthorizeUrl) // 获取授权地址
	}
}
