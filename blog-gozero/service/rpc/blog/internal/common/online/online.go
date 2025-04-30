package online

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type OnlineUserService struct {
	rdb          *redis.Client
	key          string
	expirePeriod time.Duration
}

func NewOnlineUserService(rdb *redis.Client, expire int) *OnlineUserService {
	return &OnlineUserService{
		rdb:          rdb,
		key:          "online_users",
		expirePeriod: time.Duration(expire) * time.Second,
	}
}

// 用户登录（加入 ZSet，score 为当前时间戳）
func (s *OnlineUserService) Login(ctx context.Context, userID string) error {
	score := float64(time.Now().Unix())
	return s.rdb.ZAdd(ctx, s.key, redis.Z{
		Score:  score,
		Member: userID,
	}).Err()
}

// 用户退出（从 ZSet 移除）
func (s *OnlineUserService) Logout(ctx context.Context, userID string) error {
	return s.rdb.ZRem(ctx, s.key, userID).Err()
}

// 清除过期用户（30分钟未活动）
func (s *OnlineUserService) CleanExpired(ctx context.Context) error {
	expireBefore := float64(time.Now().Add(-s.expirePeriod).Unix())
	return s.rdb.ZRemRangeByScore(ctx, s.key, "-inf", fmt.Sprintf("%f", expireBefore)).Err()
}

// 获取在线用户 ID 列表（先清理，再返回）
func (s *OnlineUserService) GetOnlineUsers(ctx context.Context, offset, limit int64) ([]string, error) {
	if err := s.CleanExpired(ctx); err != nil {
		return nil, err
	}
	return s.rdb.ZRevRange(ctx, s.key, offset, offset+limit-1).Result()
}

// 获取在线用户数量（先清理）
func (s *OnlineUserService) GetOnlineUserCount(ctx context.Context) (int64, error) {
	if err := s.CleanExpired(ctx); err != nil {
		return 0, err
	}
	return s.rdb.ZCard(ctx, s.key).Result()
}
