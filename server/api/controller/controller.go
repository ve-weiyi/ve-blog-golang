package controller

import (
	logic2 "github.com/ve-weiyi/ve-blog-golang/server/api/controller/logic"
	"github.com/ve-weiyi/ve-blog-golang/server/api/controller/svc"
)

type AppController struct {
	svcCtx                 *svc.ControllerContext         //持有的service层引用
	BlogController         *logic2.BlogController         //博客
	AuthController         *logic2.AuthController         //登录权限认证
	UserController         *logic2.UserController         //用户登录注册
	AdminController        *logic2.AdminController        //管理员
	ApiController          *logic2.ApiController          //api路由
	MenuController         *logic2.MenuController         //菜单
	RoleController         *logic2.RoleController         //角色
	ArticleController      *logic2.ArticleController      //文章
	CategoryController     *logic2.CategoryController     //文章分类
	FriendLinkController   *logic2.FriendLinkController   //文章分类
	TagController          *logic2.TagController          //文章标签
	PageController         *logic2.PageController         //页面
	CommentController      *logic2.CommentController      //评论
	PhotoController        *logic2.PhotoController        //照片
	PhotoAlbumController   *logic2.PhotoAlbumController   //相册
	TalkController         *logic2.TalkController         //说说
	CaptchaController      *logic2.CaptchaController      //验证码
	UploadController       *logic2.UploadController       //文件上传
	OperationLogController *logic2.OperationLogController //操作记录
	RemarkController       *logic2.RemarkController       //留言
}

func NewController(svcCtx *svc.ControllerContext) *AppController {
	return &AppController{
		svcCtx:                 svcCtx,
		BlogController:         logic2.NewBlogController(svcCtx),
		AuthController:         logic2.NewAuthController(svcCtx),
		AdminController:        logic2.NewAdminController(svcCtx),
		UserController:         logic2.NewUserController(svcCtx),
		ApiController:          logic2.NewApiController(svcCtx),
		MenuController:         logic2.NewMenuController(svcCtx),
		RoleController:         logic2.NewRoleController(svcCtx),
		ArticleController:      logic2.NewArticleController(svcCtx),
		CategoryController:     logic2.NewCategoryController(svcCtx),
		FriendLinkController:   logic2.NewFriendLinkController(svcCtx),
		TagController:          logic2.NewTagController(svcCtx),
		PageController:         logic2.NewPageController(svcCtx),
		CommentController:      logic2.NewCommentController(svcCtx),
		PhotoController:        logic2.NewPhotoController(svcCtx),
		PhotoAlbumController:   logic2.NewPhotoAlbumController(svcCtx),
		TalkController:         logic2.NewTalkController(svcCtx),
		CaptchaController:      logic2.NewCaptchaController(svcCtx),
		UploadController:       logic2.NewUploadController(svcCtx),
		OperationLogController: logic2.NewOperationLogController(svcCtx),
		RemarkController:       logic2.NewRemarkController(svcCtx),
	}
}
