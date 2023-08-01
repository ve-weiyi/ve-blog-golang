package svc

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rabbitmq/handler"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/upload"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/glog"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	*repository.AppRepository

	Config *config.Config
	//MainDB *gorm.DB
	//DBList map[string]*gorm.DB
	//Cache  *redis.Client
	Log            *glog.Glogger
	Token          *jjwt.JwtToken
	RBAC           *rbac.CachedEnforcer
	Captcha        *captcha.CaptchaRepository
	EmailPublisher rabbitmq.MessagePublisher
	Uploader       upload.Uploader
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	ctx := svc.NewRepositoryContext(cfg)
	repo := repository.NewRepository(ctx)
	if repo == nil {
		panic("repository cannot be null")
	}

	email := handler.NewEmailHandler(cfg.RabbitMQ.GetUrl())
	return &ServiceContext{
		AppRepository: repo,
		Config:        cfg,
		//MainDB: global.DB,
		//DBList: global.DBList,
		//Cache:  global.REDIS,
		Log:   global.LOG,
		Token: global.JWT,
		//RBAC:           global.RbacEnforcer,
		Captcha:        captcha.NewCaptchaRepository(),
		EmailPublisher: email.Publisher(),
		Uploader:       upload.NewOss(&cfg.Upload),
	}
}
