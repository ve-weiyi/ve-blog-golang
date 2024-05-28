package svc

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	*repository.AppRepository

	Config         *config.Config
	Token          *jjwt.JwtToken
	RBAC           *rbac.CachedEnforcer
	CaptchaHolder  *captcha.CaptchaHolder
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
		Token:          global.JWT,
		CaptchaHolder:  captcha.NewCaptchaHolder(captcha.NewRedisStore(global.REDIS)),
		EmailPublisher: global.EmailMQ,
		Uploader:       global.Uploader,
		AIChatGPT:      global.AIChatGPT,
	}
}
