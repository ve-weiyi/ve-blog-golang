package svc

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	*repository.AppRepository

	Config         *config.Config
	Log            *glog.Glogger
	Token          *jjwt.JwtToken
	RBAC           *rbac.CachedEnforcer
	Captcha        *captcha.CaptchaHolder
	EmailPublisher rabbitmq.MessagePublisher
	Uploader       upload.Uploader
	AIChatGPT      *chatgpt.AIChatGPT
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	ctx := svc.NewRepositoryContext(cfg)
	repo := repository.NewRepository(ctx)
	if repo == nil {
		panic("repository cannot be null")
	}

	return &ServiceContext{
		AppRepository:  repo,
		Config:         cfg,
		Log:            global.LOG,
		Token:          global.JWT,
		Captcha:        captcha.NewCaptchaHolder(captcha.NewDefaultRedisStore(global.REDIS)),
		EmailPublisher: global.EmailMQ,
		Uploader:       global.Uploader,
		AIChatGPT:      global.AIChatGPT,
	}
}
