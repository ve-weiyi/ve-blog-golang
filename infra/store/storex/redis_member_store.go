package storex

import (
	"context"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisMemberStore 基于 Redis ZSET 的 MemberStore 实现。
// member 的过期时间记在各自的 score 上，因此互不影响。
type RedisMemberStore struct {
	client *redis.Client
}

// NewRedisMemberStore 创建 RedisMemberStore。
func NewRedisMemberStore(client *redis.Client) *RedisMemberStore {
	return &RedisMemberStore{client: client}
}

func (s *RedisMemberStore) Add(ctx context.Context, key, member string, ttl time.Duration) error {
	// 过期时间只记在 member 的 score 上。给整个 key 设 TTL 会让短 TTL 成员
	// 连带删掉长 TTL 成员，与「member 之间独立过期」的契约相反。
	return s.client.ZAdd(ctx, key, redis.Z{Score: expiryScore(ttl), Member: member}).Err()
}

func (s *RedisMemberStore) Members(ctx context.Context, key string) ([]string, error) {
	expiries, err := s.Expiries(ctx, key)
	if err != nil {
		return nil, err
	}

	return sortedByExpiry(expiries), nil
}

func (s *RedisMemberStore) Expiries(ctx context.Context, key string) (map[string]time.Time, error) {
	// 清理已过期 member
	if err := s.client.ZRemRangeByScore(ctx, key, "-inf", formatFloat(nowSeconds())).Err(); err != nil {
		return nil, err
	}

	items, err := s.client.ZRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	expiries := make(map[string]time.Time, len(items))
	for _, item := range items {
		member, ok := item.Member.(string)
		if !ok {
			continue
		}
		expiries[member] = scoreToExpiry(item.Score)
	}

	return expiries, nil
}

func (s *RedisMemberStore) Has(ctx context.Context, key, member string) (bool, error) {
	score, err := s.client.ZScore(ctx, key, member).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if score <= nowSeconds() {
		// 惰性清理
		_ = s.client.ZRem(ctx, key, member).Err()
		return false, nil
	}

	return true, nil
}

func (s *RedisMemberStore) Remove(ctx context.Context, key string, members ...string) error {
	if len(members) == 0 {
		return s.client.Del(ctx, key).Err()
	}

	return s.client.ZRem(ctx, key, members).Err()
}

// ---- helpers ----------------------------------------------------------------

// expiryScore 把 TTL 换算成 ZSET 的过期时间戳。
// ttl <= 0 表示不过期，用 +Inf 表示 —— Redis 的 score 允许取无穷大。
// 使用亚秒精度，短 TTL 才不会因为整秒截断而延迟过期。
func expiryScore(ttl time.Duration) float64 {
	if ttl <= 0 {
		return math.Inf(1)
	}

	return nowSeconds() + ttl.Seconds()
}

func nowSeconds() float64 {
	return float64(time.Now().UnixNano()) / float64(time.Second)
}

// scoreToExpiry 是 expiryScore 的逆：+Inf 还原成零值时刻（表示不过期）。
func scoreToExpiry(score float64) time.Time {
	if math.IsInf(score, 1) {
		return time.Time{}
	}

	sec, frac := math.Modf(score)
	return time.Unix(int64(sec), int64(frac*float64(time.Second)))
}

// sortedByExpiry 按过期时刻升序排成员，同刻按成员名 —— 与 ZSET 按 score 的顺序一致，
// 不过期（零值时刻）的排最后，对应 score 为 +Inf。
func sortedByExpiry(expiries map[string]time.Time) []string {
	members := make([]string, 0, len(expiries))
	for member := range expiries {
		members = append(members, member)
	}

	sort.Slice(members, func(i, j int) bool {
		a, b := expirySortKey(expiries[members[i]]), expirySortKey(expiries[members[j]])
		if a == b {
			return members[i] < members[j]
		}
		return a < b
	})

	return members
}

// expirySortKey 给不过期的成员一个比任何时刻都大的键。
func expirySortKey(expireAt time.Time) float64 {
	if expireAt.IsZero() {
		return math.Inf(1)
	}

	return float64(expireAt.UnixNano())
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
