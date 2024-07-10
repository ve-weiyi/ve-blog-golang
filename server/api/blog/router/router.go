package router

import (
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type AppRouter struct {
	WebsiteRouter      *WebsiteRouter      // 博客
	WebsocketRouter    *WebsocketRouter    // websocket
	AuthRouter         *AuthRouter         // 权限认证
	UserRouter         *UserRouter         // 用户
	ApiRouter          *ApiRouter          // api
	MenuRouter         *MenuRouter         // 菜单
	RoleRouter         *RoleRouter         // 角色
	ArticleRouter      *ArticleRouter      // 文章
	CategoryRouter     *CategoryRouter     // 文章分类
	FriendLinkRouter   *FriendLinkRouter   // 文章分类
	TagRouter          *TagRouter          // 文章标签
	PageRouter         *PageRouter         // 页面
	CommentRouter      *CommentRouter      // 评论
	PhotoRouter        *PhotoRouter        // 照片
	PhotoAlbumRouter   *PhotoAlbumRouter   // 相册
	TalkRouter         *TalkRouter         // 说说
	CaptchaRouter      *CaptchaRouter      // 验证码
	UploadRouter       *UploadRouter       // 文件上传
	OperationLogRouter *OperationLogRouter // 操作记录
	RemarkRouter       *RemarkRouter       // 留言
	AIRouter           *AIRouter           // AI
}

func NewRouter(svcCtx *svc.ServiceContext) *AppRouter {
	return &AppRouter{
		WebsiteRouter:      NewWebsiteRouter(svcCtx),
		WebsocketRouter:    NewWebsocketRouter(svcCtx),
		AuthRouter:         NewAuthRouter(svcCtx),
		UserRouter:         NewUserRouter(svcCtx),
		ApiRouter:          NewApiRouter(svcCtx),
		MenuRouter:         NewMenuRouter(svcCtx),
		RoleRouter:         NewRoleRouter(svcCtx),
		ArticleRouter:      NewArticleRouter(svcCtx),
		CategoryRouter:     NewCategoryRouter(svcCtx),
		FriendLinkRouter:   NewFriendLinkRouter(svcCtx),
		TagRouter:          NewTagRouter(svcCtx),
		PageRouter:         NewPageRouter(svcCtx),
		CommentRouter:      NewCommentRouter(svcCtx),
		PhotoRouter:        NewPhotoRouter(svcCtx),
		PhotoAlbumRouter:   NewPhotoAlbumRouter(svcCtx),
		TalkRouter:         NewTalkRouter(svcCtx),
		CaptchaRouter:      NewCaptchaRouter(svcCtx),
		UploadRouter:       NewUploadRouter(svcCtx),
		OperationLogRouter: NewOperationLogRouter(svcCtx),
		RemarkRouter:       NewRemarkRouter(svcCtx),
		AIRouter:           NewAIRouter(svcCtx),
	}
}
