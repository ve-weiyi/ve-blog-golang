package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type WebsiteRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsiteRouter(svcCtx *svctx.ServiceContext) *WebsiteRouter {
	return &WebsiteRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Blog 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *WebsiteRouter) InitWebsiteRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewWebsiteController(s.svcCtx)
	{
		publicRouter.GET("/blog", handler.GetBlogHomeInfo)  // 获取博客首页信息
		loginRouter.GET("/admin", handler.GetAdminHomeInfo) // 获取后台首页信息

		publicRouter.GET("/blog/about_me", handler.GetAboutMe)     // 查询关于我
		loginRouter.POST("/admin/about_me", handler.UpdateAboutMe) // 更新关于我

		publicRouter.GET("/blog/get_website_config", handler.GetWebsiteConfig)       // 获取网站配置
		loginRouter.PUT("/admin/update_website_config", handler.UpdateWebsiteConfig) // 更新网站配置

		loginRouter.POST("/admin/config", handler.GetConfig)   // 获取网站配置
		loginRouter.PUT("/admin/config", handler.UpdateConfig) // 更新网站配置

		loginRouter.GET("/admin/system/state", handler.GetSystemState) // 获取系统信息
		publicRouter.POST("/chat/records", handler.FindChatRecords)    // 查询前台聊天记录
	}
}
