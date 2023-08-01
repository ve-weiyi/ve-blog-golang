package service

import (
	logic2 "github.com/ve-weiyi/ve-blog-golang/server/api/service/logic"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type AppService struct {
	svcCtx                  *svc.ServiceContext             //持有的repository层引用
	BlogService             *logic2.BlogService             //博客
	AuthService             *logic2.AuthService             //登录权限认证
	UserService             *logic2.UserService             //用户
	UserAccountService      *logic2.UserAccountService      //用户登录信息
	UserLoginHistoryService *logic2.UserLoginHistoryService //用户登录历史
	CaptchaService          *logic2.CaptchaService          //验证码
	UploadService           *logic2.UploadService           //文件上传
	ChatRecordService       *logic2.ChatRecordService       //聊天记录
	RoleService             *logic2.RoleService             //角色
	MenuService             *logic2.MenuService             //菜单
	ApiService              *logic2.ApiService              //接口
	ArticleService          *logic2.ArticleService          //文章
	TagService              *logic2.TagService              //文章标签
	CategoryService         *logic2.CategoryService         //文章分类
	CommentService          *logic2.CommentService          //评论
	PhotoService            *logic2.PhotoService            //相片
	PhotoAlbumService       *logic2.PhotoAlbumService       //相册
	PageService             *logic2.PageService             //页面
	TalkService             *logic2.TalkService             //说说
	FriendLinkService       *logic2.FriendLinkService       //友链
	OperationLogService     *logic2.OperationLogService     //操作记录
	RemarkService           *logic2.RemarkService           //留言
	WebsiteConfigService    *logic2.WebsiteConfigService    //网站设置
}

func NewService(svcCtx *svc.ServiceContext) *AppService {
	return &AppService{
		svcCtx:                  svcCtx,
		BlogService:             logic2.NewBlogService(svcCtx),
		AuthService:             logic2.NewAuthService(svcCtx),
		UserService:             logic2.NewUserService(svcCtx),
		UserAccountService:      logic2.NewUserAccountService(svcCtx),
		UserLoginHistoryService: logic2.NewUserLoginHistoryService(svcCtx),
		CaptchaService:          logic2.NewCaptchaService(svcCtx),
		UploadService:           logic2.NewUploadService(svcCtx),
		ChatRecordService:       logic2.NewChatRecordService(svcCtx),
		RoleService:             logic2.NewRoleService(svcCtx),
		MenuService:             logic2.NewMenuService(svcCtx),
		ApiService:              logic2.NewApiService(svcCtx),
		ArticleService:          logic2.NewArticleService(svcCtx),
		TagService:              logic2.NewTagService(svcCtx),
		CategoryService:         logic2.NewCategoryService(svcCtx),
		CommentService:          logic2.NewCommentService(svcCtx),
		PhotoService:            logic2.NewPhotoService(svcCtx),
		PhotoAlbumService:       logic2.NewPhotoAlbumService(svcCtx),
		PageService:             logic2.NewPageService(svcCtx),
		TalkService:             logic2.NewTalkService(svcCtx),
		FriendLinkService:       logic2.NewFriendLinkService(svcCtx),
		OperationLogService:     logic2.NewOperationLogService(svcCtx),
		RemarkService:           logic2.NewRemarkService(svcCtx),
		WebsiteConfigService:    logic2.NewWebsiteConfigService(svcCtx),
	}
}
