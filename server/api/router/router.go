package router

import (
	logic2 "github.com/ve-weiyi/ve-blog-golang/server/api/router/logic"
	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type AppRouter struct {
	svcCtx             *svc.RouterContext         //持有的controller引用
	BlogRouter         *logic2.BlogRouter         //博客
	AuthRouter         *logic2.AuthRouter         //权限认证
	AdminRouter        *logic2.AdminRouter        //管理员
	UserRouter         *logic2.UserRouter         //用户
	ApiRouter          *logic2.ApiRouter          //api
	MenuRouter         *logic2.MenuRouter         //菜单
	RoleRouter         *logic2.RoleRouter         //角色
	ArticleRouter      *logic2.ArticleRouter      //文章
	CategoryRouter     *logic2.CategoryRouter     //文章分类
	FriendLinkRouter   *logic2.FriendLinkRouter   //文章分类
	TagRouter          *logic2.TagRouter          //文章标签
	PageRouter         *logic2.PageRouter         //页面
	CommentRouter      *logic2.CommentRouter      //评论
	PhotoRouter        *logic2.PhotoRouter        //照片
	PhotoAlbumRouter   *logic2.PhotoAlbumRouter   //相册
	TalkRouter         *logic2.TalkRouter         //说说
	CaptchaRouter      *logic2.CaptchaRouter      //验证码
	UploadRouter       *logic2.UploadRouter       //文件上传
	OperationLogRouter *logic2.OperationLogRouter //操作记录
	RemarkRouter       *logic2.RemarkRouter       //留言
}

func NewRouter(svcCtx *svc.RouterContext) *AppRouter {
	return &AppRouter{
		svcCtx:             svcCtx,
		BlogRouter:         logic2.NewBlogRouter(svcCtx),
		AuthRouter:         logic2.NewAuthRouter(svcCtx),
		AdminRouter:        logic2.NewAdminRouter(svcCtx),
		UserRouter:         logic2.NewUserRouter(svcCtx),
		ApiRouter:          logic2.NewApiRouter(svcCtx),
		MenuRouter:         logic2.NewMenuRouter(svcCtx),
		RoleRouter:         logic2.NewRoleRouter(svcCtx),
		ArticleRouter:      logic2.NewArticleRouter(svcCtx),
		CategoryRouter:     logic2.NewCategoryRouter(svcCtx),
		FriendLinkRouter:   logic2.NewFriendLinkRouter(svcCtx),
		TagRouter:          logic2.NewTagRouter(svcCtx),
		PageRouter:         logic2.NewPageRouter(svcCtx),
		CommentRouter:      logic2.NewCommentRouter(svcCtx),
		PhotoRouter:        logic2.NewPhotoRouter(svcCtx),
		PhotoAlbumRouter:   logic2.NewPhotoAlbumRouter(svcCtx),
		TalkRouter:         logic2.NewTalkRouter(svcCtx),
		CaptchaRouter:      logic2.NewCaptchaRouter(svcCtx),
		UploadRouter:       logic2.NewUploadRouter(svcCtx),
		OperationLogRouter: logic2.NewOperationLogRouter(svcCtx),
		RemarkRouter:       logic2.NewRemarkRouter(svcCtx),
	}
}
