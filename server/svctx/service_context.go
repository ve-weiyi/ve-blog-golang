package svc

import (
	"time"

	"github.com/orca-zhang/ecache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
	"github.com/ve-weiyi/ve-blog-golang/server/initialize"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	Config *config.Config

	DbEngin        *gorm.DB
	RedisEngin     *redis.Client
	LocalCache     *ecache.Cache
	Token          *jtoken.JwtInstance
	Oauth          map[string]oauth.Oauth
	CaptchaHolder  *captcha.CaptchaHolder
	AIChatGPT      *chatgpt.AIChatGPT
	EmailPublisher rabbitmq.MessagePublisher
	Uploader       upload.Uploader
	RbacHolder     rbac.RbacHolder //RBAC角色访问控制器

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

func NewServiceContext(c *config.Config) *ServiceContext {
	db, err := initialize.ConnectGorm(c.Mysql)
	if err != nil {
		panic(err)
	}

	rdb, err := initialize.ConnectRedis(c.Redis)
	if err != nil {
		panic(err)
	}

	mq, err := initialize.ConnectRabbitMq(c.RabbitMQ)
	if err != nil {
		panic(err)
	}

	up, err := initialize.Upload(c.Upload)
	if err != nil {
		panic(err)
	}

	cache := ecache.NewLRUCache(16, 200, 10*time.Second).LRU2(1024)

	ch := captcha.NewCaptchaHolder(captcha.WithRedisStore(rdb))

	gpt := chatgpt.NewAIChatGPT(
		chatgpt.WithApiKey(c.ChatGPT.ApiKey),
		chatgpt.WithApiHost(c.ChatGPT.ApiHost),
		chatgpt.WithModel(c.ChatGPT.Model),
	)

	return &ServiceContext{
		Config:         c,
		DbEngin:        db,
		RedisEngin:     rdb,
		LocalCache:     cache,
		Token:          jtoken.NewJWTInstance([]byte(c.JWT.SigningKey)),
		Oauth:          initialize.InitOauth(c.Oauth),
		CaptchaHolder:  ch,
		AIChatGPT:      gpt,
		EmailPublisher: mq,
		Uploader:       up,
		RbacHolder:     rbac.NewPermissionHolder(db),

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
