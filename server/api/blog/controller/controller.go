package controller

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/logic"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
)

type AppController struct {
	svcCtx               *svc.ControllerContext      //持有的service层引用
	BlogController       *logic.BlogController       //博客
	AuthController       *logic.AuthController       //登录权限认证
	UserController       *logic.UserController       //用户登录注册
	RbacController       *logic.RBACController       //rbac权限认证
	ApiController        *logic.ApiController        //api路由
	MenuController       *logic.MenuController       //菜单
	RoleController       *logic.RoleController       //角色
	ArticleController    *logic.ArticleController    //文章
	CategoryController   *logic.CategoryController   //文章分类
	FriendLinkController *logic.FriendLinkController //文章分类
	TagController        *logic.TagController        //文章标签
	PageController       *logic.PageController       //页面
	CommentController    *logic.CommentController    //评论
	PhotoController      *logic.PhotoController      //照片
	PhotoAlbumController *logic.PhotoAlbumController //相册
	TalkController       *logic.TalkController       //说说
	CaptchaController    *logic.CaptchaController    //验证码
	UploadController     *logic.UploadController     //文件上传
}

func NewController(svcCtx *svc.ControllerContext) *AppController {
	return &AppController{
		svcCtx:               svcCtx,
		BlogController:       logic.NewBlogController(svcCtx),
		AuthController:       logic.NewAuthController(svcCtx),
		RbacController:       logic.NewRBACController(svcCtx),
		UserController:       logic.NewUserController(svcCtx),
		ApiController:        logic.NewApiController(svcCtx),
		MenuController:       logic.NewMenuController(svcCtx),
		RoleController:       logic.NewRoleController(svcCtx),
		ArticleController:    logic.NewArticleController(svcCtx),
		CategoryController:   logic.NewCategoryController(svcCtx),
		FriendLinkController: logic.NewFriendLinkController(svcCtx),
		TagController:        logic.NewTagController(svcCtx),
		PageController:       logic.NewPageController(svcCtx),
		CommentController:    logic.NewCommentController(svcCtx),
		PhotoController:      logic.NewPhotoController(svcCtx),
		PhotoAlbumController: logic.NewPhotoAlbumController(svcCtx),
		TalkController:       logic.NewTalkController(svcCtx),
		CaptchaController:    logic.NewCaptchaController(svcCtx),
		UploadController:     logic.NewUploadController(svcCtx),
	}
}
