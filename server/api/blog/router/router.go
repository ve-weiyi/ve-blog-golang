package router

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/router/logic"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/router/svc"
)

type AppRouter struct {
	svcCtx           *svc.RouterContext //持有的controller引用
	AuthRouter       *logic.AuthRouter  //权限认证
	UserRouter       *logic.UserRouter  //权限认证
	ApiRouter        *logic.ApiRouter
	MenuRouter       *logic.MenuRouter //菜单
	RoleRouter       *logic.RoleRouter
	ArticleRouter    *logic.ArticleRouter    //文章
	CategoryRouter   *logic.CategoryRouter   //文章分类
	FriendLinkRouter *logic.FriendLinkRouter //文章分类
	TagRouter        *logic.TagRouter        //文章标签
	PageRouter       *logic.PageRouter       //页面
	CommentRouter    *logic.CommentRouter    //评论
	PhotoRouter      *logic.PhotoRouter      //照片
	PhotoAlbumRouter *logic.PhotoAlbumRouter //相册
	TalkRouter       *logic.TalkRouter       //说说
	CaptchaRouter    *logic.CaptchaRouter    //验证码
	UploadRouter     *logic.UploadRouter     //文件上传
}

func NewRouter(svcCtx *svc.RouterContext) *AppRouter {
	return &AppRouter{
		svcCtx:           svcCtx,
		AuthRouter:       logic.NewAuthRouter(svcCtx),
		UserRouter:       logic.NewUserRouter(svcCtx),
		ApiRouter:        logic.NewApiRouter(svcCtx),
		MenuRouter:       logic.NewMenuRouter(svcCtx),
		RoleRouter:       logic.NewRoleRouter(svcCtx),
		ArticleRouter:    logic.NewArticleRouter(svcCtx),
		CategoryRouter:   logic.NewCategoryRouter(svcCtx),
		FriendLinkRouter: logic.NewFriendLinkRouter(svcCtx),
		TagRouter:        logic.NewTagRouter(svcCtx),
		PageRouter:       logic.NewPageRouter(svcCtx),
		CommentRouter:    logic.NewCommentRouter(svcCtx),
		PhotoRouter:      logic.NewPhotoRouter(svcCtx),
		PhotoAlbumRouter: logic.NewPhotoAlbumRouter(svcCtx),
		TalkRouter:       logic.NewTalkRouter(svcCtx),
		CaptchaRouter:    logic.NewCaptchaRouter(svcCtx),
		UploadRouter:     logic.NewUploadRouter(svcCtx),
	}
}
