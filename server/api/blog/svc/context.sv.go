package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	Config *config.Config

	DbEngin        *gorm.DB
	DBList         map[string]*gorm.DB
	Cache          *redis.Client
	Token          *jjwt.JwtToken
	RBAC           *rbac.CachedEnforcer
	CaptchaHolder  *captcha.CaptchaHolder
	EmailPublisher rabbitmq.MessagePublisher
	Uploader       upload.Uploader
	AIChatGPT      *chatgpt.AIChatGPT

	ApiRepository              *repository.ApiRepository              //api路由
	ArticleRepository          *repository.ArticleRepository          //文章
	MenuRepository             *repository.MenuRepository             //菜单
	RoleRepository             *repository.RoleRepository             //角色
	RoleApiRepository          *repository.RoleApiRepository          //角色-api关联
	RoleMenuRepository         *repository.RoleMenuRepository         //角色-菜单关联
	UserAccountRepository      *repository.UserAccountRepository      //用户登录信息
	UserInformationRepository  *repository.UserInformationRepository  //用户信息
	UserLoginHistoryRepository *repository.UserLoginHistoryRepository //用户登录历史
	UserOauthRepository        *repository.UserOauthRepository        //第三方登录信息
	UserRoleRepository         *repository.UserRoleRepository         //用户-角色关联
	CategoryRepository         *repository.CategoryRepository         //文章分类
	FriendLinkRepository       *repository.FriendLinkRepository       //文章分类
	TagRepository              *repository.TagRepository              //文章标签
	PageRepository             *repository.PageRepository             //页面
	CommentRepository          *repository.CommentRepository          //评论
	PhotoRepository            *repository.PhotoRepository            //照片
	PhotoAlbumRepository       *repository.PhotoAlbumRepository       //相册
	TalkRepository             *repository.TalkRepository             //说说
	ArticleTagRepository       *repository.ArticleTagRepository       //文章标签映射
	ChatRecordRepository       *repository.ChatRecordRepository       //聊天记录
	UniqueViewRepository       *repository.UniqueViewRepository       //页面访问数量
	OperationLogRepository     *repository.OperationLogRepository     //操作记录
	RemarkRepository           *repository.RemarkRepository           //留言
	WebsiteConfigRepository    *repository.WebsiteConfigRepository    //网站设置
	UploadRecordRepository     *repository.UploadRecordRepository     //文件上传
	ChatSessionRepository      *repository.ChatSessionRepository      //聊天会话
	ChatMessageRepository      *repository.ChatMessageRepository      //聊天消息
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	db := global.DB
	rdb := global.REDIS

	return &ServiceContext{
		Config:         cfg,
		Token:          global.JWT,
		CaptchaHolder:  captcha.NewCaptchaHolder(captcha.NewRedisStore(global.REDIS)),
		EmailPublisher: global.EmailMQ,
		Uploader:       global.Uploader,
		AIChatGPT:      global.AIChatGPT,
		DbEngin:        global.DB,
		DBList:         global.DBList,
		Cache:          global.REDIS,

		ApiRepository:              repository.NewApiRepository(db, rdb),
		ArticleRepository:          repository.NewArticleRepository(db, rdb),
		MenuRepository:             repository.NewMenuRepository(db, rdb),
		RoleRepository:             repository.NewRoleRepository(db, rdb),
		RoleApiRepository:          repository.NewRoleApiRepository(db, rdb),
		RoleMenuRepository:         repository.NewRoleMenuRepository(db, rdb),
		UserAccountRepository:      repository.NewUserAccountRepository(db, rdb),
		UserInformationRepository:  repository.NewUserInformationRepository(db, rdb),
		UserLoginHistoryRepository: repository.NewUserLoginHistoryRepository(db, rdb),
		UserOauthRepository:        repository.NewUserOauthRepository(db, rdb),
		UserRoleRepository:         repository.NewUserRoleRepository(db, rdb),
		CategoryRepository:         repository.NewCategoryRepository(db, rdb),
		FriendLinkRepository:       repository.NewFriendLinkRepository(db, rdb),
		TagRepository:              repository.NewTagRepository(db, rdb),
		PageRepository:             repository.NewPageRepository(db, rdb),
		CommentRepository:          repository.NewCommentRepository(db, rdb),
		PhotoRepository:            repository.NewPhotoRepository(db, rdb),
		PhotoAlbumRepository:       repository.NewPhotoAlbumRepository(db, rdb),
		TalkRepository:             repository.NewTalkRepository(db, rdb),
		ArticleTagRepository:       repository.NewArticleTagRepository(db, rdb),
		ChatRecordRepository:       repository.NewChatRecordRepository(db, rdb),
		UniqueViewRepository:       repository.NewUniqueViewRepository(db, rdb),
		OperationLogRepository:     repository.NewOperationLogRepository(db, rdb),
		RemarkRepository:           repository.NewRemarkRepository(db, rdb),
		WebsiteConfigRepository:    repository.NewWebsiteConfigRepository(db, rdb),
		UploadRecordRepository:     repository.NewUploadRecordRepository(db, rdb),
		ChatSessionRepository:      repository.NewChatSessionRepository(db, rdb),
		ChatMessageRepository:      repository.NewChatMessageRepository(db, rdb),
	}
}
