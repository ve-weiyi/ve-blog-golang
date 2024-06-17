package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type AppService struct {
	svcCtx              *svctx.ServiceContext // 持有的repository层引用
	WebsiteService      *WebsiteService       // 博客
	AuthService         *AuthService          // 登录权限认证
	UserService         *UserService          // 用户
	UserAccountService  *UserAccountService   // 用户信息
	CaptchaService      *CaptchaService       // 验证码
	UploadService       *UploadService        // 文件上传
	ChatRecordService   *ChatRecordService    // 聊天记录
	RoleService         *RoleService          // 角色
	MenuService         *MenuService          // 菜单
	ApiService          *ApiService           // 接口
	ArticleService      *ArticleService       // 文章
	TagService          *TagService           // 文章标签
	CategoryService     *CategoryService      // 文章分类
	CommentService      *CommentService       // 评论
	PhotoService        *PhotoService         // 相片
	PhotoAlbumService   *PhotoAlbumService    // 相册
	PageService         *PageService          // 页面
	TalkService         *TalkService          // 说说
	FriendLinkService   *FriendLinkService    // 友链
	OperationLogService *OperationLogService  // 操作记录
	RemarkService       *RemarkService        // 留言
	AIChatService       *AIService            // AI聊天
}

func NewService(svcCtx *svctx.ServiceContext) *AppService {
	return &AppService{
		svcCtx:              svcCtx,
		WebsiteService:      NewWebsiteService(svcCtx),
		AuthService:         NewAuthService(svcCtx),
		UserService:         NewUserService(svcCtx),
		UserAccountService:  NewUserAccountService(svcCtx),
		CaptchaService:      NewCaptchaService(svcCtx),
		UploadService:       NewUploadService(svcCtx),
		ChatRecordService:   NewChatRecordService(svcCtx),
		RoleService:         NewRoleService(svcCtx),
		MenuService:         NewMenuService(svcCtx),
		ApiService:          NewApiService(svcCtx),
		ArticleService:      NewArticleService(svcCtx),
		TagService:          NewTagService(svcCtx),
		CategoryService:     NewCategoryService(svcCtx),
		CommentService:      NewCommentService(svcCtx),
		PhotoService:        NewPhotoService(svcCtx),
		PhotoAlbumService:   NewPhotoAlbumService(svcCtx),
		PageService:         NewPageService(svcCtx),
		TalkService:         NewTalkService(svcCtx),
		FriendLinkService:   NewFriendLinkService(svcCtx),
		OperationLogService: NewOperationLogService(svcCtx),
		RemarkService:       NewRemarkService(svcCtx),
		AIChatService:       NewAIService(svcCtx),
	}
}
