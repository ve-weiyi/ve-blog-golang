package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AuthRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewAuthRouter(svcCtx *svctx.ServiceContext) *AuthRouter {
	return &AuthRouter{
		svcCtx: svcCtx,
	}
}

func (s *AuthRouter) Register(r *gin.RouterGroup) {
	// Auth
	// []
	{
		group := r.Group("/admin-api/v1")

		handler := controller.NewAuthController(s.svcCtx)
		// 获取游客身份信息
		group.GET("/get_tourist_info", handler.GetTouristInfo)
	}
	// Auth
	// []
	{
		group := r.Group("/admin-api/v1")

		handler := controller.NewAuthController(s.svcCtx)
		// 邮箱登录
		group.POST("/email_login", handler.EmailLogin)
		// 获取验证码
		group.POST("/get_captcha_code", handler.GetCaptchaCode)
		// 第三方登录授权地址
		group.POST("/get_oauth_authorize_url", handler.GetOauthAuthorizeUrl)
		// 登录
		group.POST("/login", handler.Login)
		// 手机登录
		group.POST("/phone_login", handler.PhoneLogin)
		// 注册
		group.POST("/register", handler.Register)
		// 重置密码
		group.POST("/reset_password", handler.ResetPassword)
		// 发送邮件验证码
		group.POST("/send_email_verify_code", handler.SendEmailVerifyCode)
		// 发送手机验证码
		group.POST("/send_phone_verify_code", handler.SendPhoneVerifyCode)
		// 第三方登录
		group.POST("/third_login", handler.ThirdLogin)
	}
	// Auth
	// [JwtToken]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)

		handler := controller.NewAuthController(s.svcCtx)
		// 注销
		group.POST("/logoff", handler.Logoff)
		// 登出
		group.POST("/logout", handler.Logout)
	}
}
