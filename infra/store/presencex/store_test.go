package presencex_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/blog-cloud/infra/store/presencex"
	"github.com/ve-weiyi/blog-cloud/infra/store/storex"
)

// 本文件测「Store 接一个真的存储实现」这条组合路径：presence_test.go 用内存替身测领域契约，
// 这里用 miniredis 走 storex.RedisMemberStore 的真实命令 —— 两者差的正是命令层。

func newRedisStore(t *testing.T, ttl time.Duration) *presencex.Store {
	t.Helper()

	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	return presencex.NewStore(storex.NewRedisMemberStore(rdb), "blog:online:user", ttl)
}

// 同一用户两条连接在真实命令上也要分得开；断开其一后用户仍在线。
func TestStoreAgainstRedis_MultiConnection(t *testing.T) {
	ctx := context.Background()
	store := newRedisStore(t, time.Minute)

	for _, ref := range []string{"ref-a", "ref-b"} {
		if err := store.Add(ctx, "uid-1", ref, time.Minute); err != nil {
			t.Fatalf("Add(%s): %v", ref, err)
		}
	}

	stats, err := store.PresenceStats(ctx)
	if err != nil {
		t.Fatalf("PresenceStats: %v", err)
	}
	if stats.NumUsers != 1 || stats.NumClients != 2 {
		t.Errorf("stats = %+v, want {NumUsers:1 NumClients:2}", stats)
	}

	if err := store.Remove(ctx, "uid-1", "ref-a"); err != nil {
		t.Fatalf("Remove: %v", err)
	}

	users, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(users) != 1 || users[0].Login != "uid-1" || users[0].Refs != 1 {
		t.Errorf("List = %+v, want 一条 uid-1 且 Refs 为 1", users)
	}
}

// 超过存活时长后在真实命令路径上也不再可见。
func TestStoreAgainstRedis_ExpiredConnectionDisappears(t *testing.T) {
	ctx := context.Background()
	store := newRedisStore(t, time.Minute)

	if err := store.Add(ctx, "uid-1", "ref-a", 20*time.Millisecond); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if err := store.Add(ctx, "uid-2", "ref-b", time.Hour); err != nil {
		t.Fatalf("Add: %v", err)
	}

	time.Sleep(80 * time.Millisecond)

	users, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(users) != 1 || users[0].Login != "uid-2" {
		t.Errorf("超期后 List = %+v, want 只剩 uid-2", users)
	}
}
