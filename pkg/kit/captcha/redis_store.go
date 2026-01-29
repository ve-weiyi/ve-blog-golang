package captcha

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/ve-blog-golang/pkg/kit/logz"
)

// RedisStore 验证码存储
type RedisStore struct {
	Redis      *redis.Client   // Redis 客户端
	Expiration time.Duration   // 过期时间，默认 15 分钟
	Context    context.Context // 上下文
	Logger     logz.Logger     // 日志接口
}

// NewRedisStore 创建 Redis 存储实例
func NewRedisStore(rd *redis.Client, logger logz.Logger) *RedisStore {
	if logger == nil {
		logger = logz.NewDefaultLogger() // 如果未提供 logger，使用默认实现
	}
	return &RedisStore{
		Expiration: 15 * 60 * time.Second,
		Redis:      rd,
		Context:    context.Background(),
		Logger:     logger,
	}
}

func (rs *RedisStore) Set(key string, value string) error {
	err := rs.Redis.Set(rs.Context, key, value, rs.Expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := rs.Redis.Get(rs.Context, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		err := rs.Redis.Del(rs.Context, key).Err()
		if err != nil {
			return ""
		}
	}
	return val
}

// Verify 验证验证码
func (rs *RedisStore) Verify(key, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	if v == "" {
		return false
	}
	if rs.Logger != nil {
		rs.Logger.Debugf("RedisStore Verify. key:%v, answer:%v, v:%v, clear:%v", key, answer, v, clear)
	}
	return v == answer
}
