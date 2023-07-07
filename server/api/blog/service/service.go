package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/logic"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type AppService struct {
	svcCtx                  *svc.ServiceContext            //持有的repository层引用
	AuthService             *logic.AuthService             //登录权限认证
	UserService             *logic.UserService             //用户
	UserAccountService      *logic.UserAccountService      //用户登录信息
	UserLoginHistoryService *logic.UserLoginHistoryService //用户登录历史
	CaptchaService          *logic.CaptchaService          //验证码
	UploadService           *logic.UploadService           //文件上传
	ChatRecordService       *logic.ChatRecordService       //聊天记录
	RoleService             *logic.RoleService             //角色
	MenuService             *logic.MenuService             //菜单
	ApiService              *logic.ApiService              //接口
	ArticleService          *logic.ArticleService          //文章
	TagService              *logic.TagService              //文章标签
	CategoryService         *logic.CategoryService         //文章分类
	CommentService          *logic.CommentService          //评论
	PhotoService            *logic.PhotoService            //相片
	PhotoAlbumService       *logic.PhotoAlbumService       //相册
	PageService             *logic.PageService             //页面
	TalkService             *logic.TalkService             //说说
	FriendLinkService       *logic.FriendLinkService       //友链
}

func NewService(svcCtx *svc.ServiceContext) *AppService {
	return &AppService{
		svcCtx:                  svcCtx,
		AuthService:             logic.NewAuthService(svcCtx),
		UserService:             logic.NewUserService(svcCtx),
		UserAccountService:      logic.NewUserAccountService(svcCtx),
		UserLoginHistoryService: logic.NewUserLoginHistoryService(svcCtx),
		CaptchaService:          logic.NewCaptchaService(svcCtx),
		UploadService:           logic.NewUploadService(svcCtx),
		ChatRecordService:       logic.NewChatRecordService(svcCtx),
		RoleService:             logic.NewRoleService(svcCtx),
		MenuService:             logic.NewMenuService(svcCtx),
		ApiService:              logic.NewApiService(svcCtx),
		ArticleService:          logic.NewArticleService(svcCtx),
		TagService:              logic.NewTagService(svcCtx),
		CategoryService:         logic.NewCategoryService(svcCtx),
		CommentService:          logic.NewCommentService(svcCtx),
		PhotoService:            logic.NewPhotoService(svcCtx),
		PhotoAlbumService:       logic.NewPhotoAlbumService(svcCtx),
		PageService:             logic.NewPageService(svcCtx),
		TalkService:             logic.NewTalkService(svcCtx),
		FriendLinkService:       logic.NewFriendLinkService(svcCtx),
	}
}
