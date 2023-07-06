package captcha

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

// 验证码存储
type RedisStore struct {
	Redis      *redis.Client // 缓存 15分钟
	Expiration time.Duration // 过期时间
	PrefixKey  string        // 前缀
	Context    context.Context
}

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: 15 * 60 * time.Second,
		PrefixKey:  "captcha:",
		Redis:      global.REDIS,
		Context:    context.Background(),
	}
}

func (rs *RedisStore) Set(key string, value string) error {
	cacheKey := rs.PrefixKey + key
	err := rs.Redis.Set(rs.Context, cacheKey, value, rs.Expiration).Err()
	if err != nil {
		global.LOG.Errorf("RedisStoreGetError!cacheKey:%v ,err:%v", cacheKey, err)
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	cacheKey := rs.PrefixKey + key
	val, err := rs.Redis.Get(rs.Context, cacheKey).Result()
	if err != nil {
		global.LOG.Errorf("RedisStoreGetError!cacheKey:%v ,err:%v", cacheKey, err)
		return ""
	}
	if clear {
		err := rs.Redis.Del(rs.Context, cacheKey).Err()
		if err != nil {
			global.LOG.Errorf("RedisStoreGetError!cacheKey:%v ,err:%v", cacheKey, err)
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(key, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	return v == answer
}
