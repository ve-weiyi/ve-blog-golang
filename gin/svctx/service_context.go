package svctx

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/orca-zhang/ecache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/gin/config"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/middleware"
	"github.com/ve-weiyi/ve-blog-golang/gin/initialize"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jwtx"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	Config *config.Config

	DbEngin    *gorm.DB
	RedisEngin *redis.Client
	LocalCache *ecache.Cache
	Token      *jwtx.JwtInstance

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

	cache := ecache.NewLRUCache(16, 200, 10*time.Second).LRU2(1024)

	tk := jwtx.NewJWTInstance([]byte(c.JWT.SigningKey))

	return &ServiceContext{
		Config:              c,
		DbEngin:             db,
		RedisEngin:          rdb,
		LocalCache:          cache,
		Token:               tk,
		MiddlewareSignToken: middleware.SignToken(),
		MiddlewareJwtToken:  middleware.JwtToken(tk),
		MiddlewareOperation: middleware.GinLogger(),
	}
}
