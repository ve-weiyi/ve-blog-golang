package router

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/router/logic"
	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type AppRouter struct {
	svcCtx             *svc.RouterContext        //持有的controller引用
	BlogRouter         *logic.BlogRouter         //博客
	AuthRouter         *logic.AuthRouter         //权限认证
	AdminRouter        *logic.AdminRouter        //管理员
	UserRouter         *logic.UserRouter         //用户
	ApiRouter          *logic.ApiRouter          //api
	MenuRouter         *logic.MenuRouter         //菜单
	RoleRouter         *logic.RoleRouter         //角色
	ArticleRouter      *logic.ArticleRouter      //文章
	CategoryRouter     *logic.CategoryRouter     //文章分类
	FriendLinkRouter   *logic.FriendLinkRouter   //文章分类
	TagRouter          *logic.TagRouter          //文章标签
	PageRouter         *logic.PageRouter         //页面
	CommentRouter      *logic.CommentRouter      //评论
	PhotoRouter        *logic.PhotoRouter        //照片
	PhotoAlbumRouter   *logic.PhotoAlbumRouter   //相册
	TalkRouter         *logic.TalkRouter         //说说
	CaptchaRouter      *logic.CaptchaRouter      //验证码
	UploadRouter       *logic.UploadRouter       //文件上传
	OperationLogRouter *logic.OperationLogRouter //操作记录
	RemarkRouter       *logic.RemarkRouter       //留言
}

func NewRouter(svcCtx *svc.RouterContext) *AppRouter {
	return &AppRouter{
		svcCtx:             svcCtx,
		BlogRouter:         logic.NewBlogRouter(svcCtx),
		AuthRouter:         logic.NewAuthRouter(svcCtx),
		AdminRouter:        logic.NewAdminRouter(svcCtx),
		UserRouter:         logic.NewUserRouter(svcCtx),
		ApiRouter:          logic.NewApiRouter(svcCtx),
		MenuRouter:         logic.NewMenuRouter(svcCtx),
		RoleRouter:         logic.NewRoleRouter(svcCtx),
		ArticleRouter:      logic.NewArticleRouter(svcCtx),
		CategoryRouter:     logic.NewCategoryRouter(svcCtx),
		FriendLinkRouter:   logic.NewFriendLinkRouter(svcCtx),
		TagRouter:          logic.NewTagRouter(svcCtx),
		PageRouter:         logic.NewPageRouter(svcCtx),
		CommentRouter:      logic.NewCommentRouter(svcCtx),
		PhotoRouter:        logic.NewPhotoRouter(svcCtx),
		PhotoAlbumRouter:   logic.NewPhotoAlbumRouter(svcCtx),
		TalkRouter:         logic.NewTalkRouter(svcCtx),
		CaptchaRouter:      logic.NewCaptchaRouter(svcCtx),
		UploadRouter:       logic.NewUploadRouter(svcCtx),
		OperationLogRouter: logic.NewOperationLogRouter(svcCtx),
		RemarkRouter:       logic.NewRemarkRouter(svcCtx),
	}
}
