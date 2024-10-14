package svctx

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/orca-zhang/ecache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"

	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/middleware"
	"github.com/ve-weiyi/ve-blog-golang/server/initialize"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	Config *config.Config

	DbEngin             *gorm.DB
	RedisEngin          *redis.Client
	LocalCache          *ecache.Cache
	Token               *jtoken.JwtInstance
	Oauth               map[string]oauth.Oauth
	CaptchaHolder       *captcha.CaptchaHolder
	AIChatGPT           *chatgpt.AIChatGPT
	EmailPublisher      rabbitmq.MessagePublisher
	Uploader            upload.Uploader
	MiddlewareSignToken gin.HandlerFunc
	MiddlewareJwtToken  gin.HandlerFunc
	MiddlewareOperation gin.HandlerFunc
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

	tk := jtoken.NewJWTInstance([]byte(c.JWT.SigningKey))

	return &ServiceContext{
		Config:              c,
		DbEngin:             db,
		RedisEngin:          rdb,
		LocalCache:          cache,
		Token:               tk,
		Oauth:               initialize.InitOauth(c.Oauth),
		CaptchaHolder:       ch,
		AIChatGPT:           gpt,
		EmailPublisher:      mq,
		Uploader:            up,
		MiddlewareSignToken: middleware.SignToken(),
		MiddlewareJwtToken:  middleware.JwtToken(tk),
		MiddlewareOperation: middleware.GinLogger(),
	}
}
