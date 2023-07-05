package service

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/logic"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type AppService struct {
	svcCtx                  *svc.ServiceContext            //持有的repository层引用
	AuthService             *logic.AuthService             //登录权限认证
	UserService             *logic.UserService             //用户
	ApiService              *logic.ApiService              //api路由
	ArticleService          *logic.ArticleService          //文章
	CasbinRuleService       *logic.CasbinRuleService       //casbin_rule
	MenuService             *logic.MenuService             //菜单
	RoleService             *logic.RoleService             //角色
	RoleApiService          *logic.RoleApiService          //角色-api关联
	RoleMenuService         *logic.RoleMenuService         //角色-菜单关联
	UserAccountService      *logic.UserAccountService      //用户登录信息
	UserInformationService  *logic.UserInformationService  //用户信息
	UserLoginHistoryService *logic.UserLoginHistoryService //用户登录历史
	UserOauthService        *logic.UserOauthService        //第三方登录信息
	UserRoleService         *logic.UserRoleService         //用户-角色关联
	CategoryService         *logic.CategoryService         //文章分类
	FriendLinkService       *logic.FriendLinkService       //文章分类
	TagService              *logic.TagService              //文章标签
	PageService             *logic.PageService             //页面
	CommentService          *logic.CommentService          //评论
	PhotoService            *logic.PhotoService            //照片
	PhotoAlbumService       *logic.PhotoAlbumService       //相册
	TalkService             *logic.TalkService             //说说
	ArticleTagService       *logic.ArticleTagService       //文章标签映射
	CaptchaService          *logic.CaptchaService          //验证码
	UploadService           *logic.UploadService           //文件上传
}

func NewService(svcCtx *svc.ServiceContext) *AppService {
	return &AppService{
		svcCtx:                  svcCtx,
		AuthService:             logic.NewAuthService(svcCtx),
		UserService:             logic.NewUserService(svcCtx),
		ApiService:              logic.NewApiService(svcCtx),
		ArticleService:          logic.NewArticleService(svcCtx),
		CasbinRuleService:       logic.NewCasbinRuleService(svcCtx),
		MenuService:             logic.NewMenuService(svcCtx),
		RoleService:             logic.NewRoleService(svcCtx),
		RoleApiService:          logic.NewRoleApiService(svcCtx),
		RoleMenuService:         logic.NewRoleMenuService(svcCtx),
		UserAccountService:      logic.NewUserAccountService(svcCtx),
		UserInformationService:  logic.NewUserInformationService(svcCtx),
		UserLoginHistoryService: logic.NewUserLoginHistoryService(svcCtx),
		UserOauthService:        logic.NewUserOauthService(svcCtx),
		UserRoleService:         logic.NewUserRoleService(svcCtx),
		CategoryService:         logic.NewCategoryService(svcCtx),
		FriendLinkService:       logic.NewFriendLinkService(svcCtx),
		TagService:              logic.NewTagService(svcCtx),
		PageService:             logic.NewPageService(svcCtx),
		CommentService:          logic.NewCommentService(svcCtx),
		PhotoService:            logic.NewPhotoService(svcCtx),
		PhotoAlbumService:       logic.NewPhotoAlbumService(svcCtx),
		TalkService:             logic.NewTalkService(svcCtx),
		ArticleTagService:       logic.NewArticleTagService(svcCtx),
		CaptchaService:          logic.NewCaptchaService(svcCtx),
		UploadService:           logic.NewUploadService(svcCtx),
	}
}
