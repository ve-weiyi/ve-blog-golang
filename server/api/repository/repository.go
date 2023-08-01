package repository

import (
	logic2 "github.com/ve-weiyi/ve-blog-golang/server/api/repository/logic"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

// 对应go-zero rpc层服务
type AppRepository struct {
	svcCtx                     *svc.RepositoryContext             //持有的repository层引用
	ApiRepository              *logic2.ApiRepository              //api路由
	ArticleRepository          *logic2.ArticleRepository          //文章
	CasbinRuleRepository       *logic2.CasbinRuleRepository       //casbin_rule
	MenuRepository             *logic2.MenuRepository             //菜单
	RoleRepository             *logic2.RoleRepository             //角色
	RoleApiRepository          *logic2.RoleApiRepository          //角色-api关联
	RoleMenuRepository         *logic2.RoleMenuRepository         //角色-菜单关联
	UserAccountRepository      *logic2.UserAccountRepository      //用户登录信息
	UserInformationRepository  *logic2.UserInformationRepository  //用户信息
	UserLoginHistoryRepository *logic2.UserLoginHistoryRepository //用户登录历史
	UserOauthRepository        *logic2.UserOauthRepository        //第三方登录信息
	UserRoleRepository         *logic2.UserRoleRepository         //用户-角色关联
	CategoryRepository         *logic2.CategoryRepository         //文章分类
	FriendLinkRepository       *logic2.FriendLinkRepository       //文章分类
	TagRepository              *logic2.TagRepository              //文章标签
	PageRepository             *logic2.PageRepository             //页面
	CommentRepository          *logic2.CommentRepository          //评论
	PhotoRepository            *logic2.PhotoRepository            //照片
	PhotoAlbumRepository       *logic2.PhotoAlbumRepository       //相册
	TalkRepository             *logic2.TalkRepository             //说说
	ArticleTagRepository       *logic2.ArticleTagRepository       //文章标签映射
	UploadRepository           *logic2.UploadRepository           //文件上传
	ChatRecordRepository       *logic2.ChatRecordRepository       //聊天记录
	UniqueViewRepository       *logic2.UniqueViewRepository       //页面访问数量
	OperationLogRepository     *logic2.OperationLogRepository     //操作记录
	RemarkRepository           *logic2.RemarkRepository           //留言
	WebsiteConfigRepository    *logic2.WebsiteConfigRepository    //网站设置
}

func NewRepository(svcCtx *svc.RepositoryContext) *AppRepository {
	return &AppRepository{
		svcCtx:                     svcCtx,
		ApiRepository:              logic2.NewApiRepository(svcCtx),
		ArticleRepository:          logic2.NewArticleRepository(svcCtx),
		CasbinRuleRepository:       logic2.NewCasbinRuleRepository(svcCtx),
		MenuRepository:             logic2.NewMenuRepository(svcCtx),
		RoleRepository:             logic2.NewRoleRepository(svcCtx),
		RoleApiRepository:          logic2.NewRoleApiRepository(svcCtx),
		RoleMenuRepository:         logic2.NewRoleMenuRepository(svcCtx),
		UserAccountRepository:      logic2.NewUserAccountRepository(svcCtx),
		UserInformationRepository:  logic2.NewUserInformationRepository(svcCtx),
		UserLoginHistoryRepository: logic2.NewUserLoginHistoryRepository(svcCtx),
		UserOauthRepository:        logic2.NewUserOauthRepository(svcCtx),
		UserRoleRepository:         logic2.NewUserRoleRepository(svcCtx),
		CategoryRepository:         logic2.NewCategoryRepository(svcCtx),
		FriendLinkRepository:       logic2.NewFriendLinkRepository(svcCtx),
		TagRepository:              logic2.NewTagRepository(svcCtx),
		PageRepository:             logic2.NewPageRepository(svcCtx),
		CommentRepository:          logic2.NewCommentRepository(svcCtx),
		PhotoRepository:            logic2.NewPhotoRepository(svcCtx),
		PhotoAlbumRepository:       logic2.NewPhotoAlbumRepository(svcCtx),
		TalkRepository:             logic2.NewTalkRepository(svcCtx),
		ArticleTagRepository:       logic2.NewArticleTagRepository(svcCtx),
		UploadRepository:           logic2.NewUploadRepository(svcCtx),
		ChatRecordRepository:       logic2.NewChatRecordRepository(svcCtx),
		UniqueViewRepository:       logic2.NewUniqueViewRepository(svcCtx),
		OperationLogRepository:     logic2.NewOperationLogRepository(svcCtx),
		RemarkRepository:           logic2.NewRemarkRepository(svcCtx),
		WebsiteConfigRepository:    logic2.NewWebsiteConfigRepository(svcCtx),
	}
}
