package storex_test

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/blog-cloud/infra/store/storex"
)

// newTestRedis 起一个 miniredis，返回连上它的客户端与实例本身。
func newTestRedis(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
	t.Helper()

	mr := miniredis.RunT(t)

	return redis.NewClient(&redis.Options{Addr: mr.Addr()}), mr
}

func TestRedisStore_SetGet(t *testing.T) {
	ctx := context.Background()
	client, _ := newTestRedis(t)
	s := storex.NewRedisStore(client)

	if err := s.Set(ctx, "k", "v", time.Hour); err != nil {
		t.Fatalf("Set: %v", err)
	}

	val, found, err := s.Get(ctx, "k")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !found || val != "v" {
		t.Errorf("Get = (%q, %v), want (v, true)", val, found)
	}

	// found 区分「没写过」与「写了个空串」。
	if err := s.Set(ctx, "empty", "", time.Hour); err != nil {
		t.Fatalf("Set(empty): %v", err)
	}
	if _, found, _ := s.Get(ctx, "never-set"); found {
		t.Error("未写过的 key 不应 found")
	}
	if val, found, _ := s.Get(ctx, "empty"); !found || val != "" {
		t.Errorf("空串 key = (%q, %v), want (\"\", true)", val, found)
	}
}

func TestRedisStore_TTL(t *testing.T) {
	t.Run("expires", func(t *testing.T) {
		ctx := context.Background()
		client, mr := newTestRedis(t)
		s := storex.NewRedisStore(client)

		if err := s.Set(ctx, "k", "v", time.Second); err != nil {
			t.Fatalf("Set: %v", err)
		}
		mr.FastForward(2 * time.Second)

		if _, found, _ := s.Get(ctx, "k"); found {
			t.Error("过期 key 仍 found")
		}
	})

	t.Run("non-positive never expires", func(t *testing.T) {
		for _, ttl := range []time.Duration{0, -time.Hour} {
			ctx := context.Background()
			client, mr := newTestRedis(t)
			s := storex.NewRedisStore(client)

			if err := s.Set(ctx, "k", "v", ttl); err != nil {
				t.Fatalf("Set(ttl=%v): %v", ttl, err)
			}
			mr.FastForward(365 * 24 * time.Hour)

			if _, found, _ := s.Get(ctx, "k"); !found {
				t.Errorf("ttl=%v: key 过期了，want 永不过期", ttl)
			}
		}
	})
}

func TestRedisStore_Delete(t *testing.T) {
	ctx := context.Background()
	client, _ := newTestRedis(t)
	s := storex.NewRedisStore(client)

	keys := []string{"k1", "k2", "k3"}
	for _, k := range keys {
		if err := s.Set(ctx, k, "v", time.Hour); err != nil {
			t.Fatalf("Set(%s): %v", k, err)
		}
	}

	if err := s.Delete(ctx, keys[0], keys[1:]...); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	for _, k := range keys {
		if _, found, _ := s.Get(ctx, k); found {
			t.Errorf("%s 删除后仍能读到", k)
		}
	}

	if err := s.Delete(ctx, "never-set"); err != nil {
		t.Errorf("删不存在的 key 报错: %v", err)
	}
}

func TestRedisStore_Keys(t *testing.T) {
	ctx := context.Background()
	client, _ := newTestRedis(t)
	s := storex.NewRedisStore(client)

	for _, k := range []string{"app:access:u1:d1", "app:access:u1:d2", "app:refresh:u2:d1"} {
		if err := s.Set(ctx, k, "v", time.Hour); err != nil {
			t.Fatalf("Set(%s): %v", k, err)
		}
	}

	got, err := s.Keys(ctx, "app:access:*")
	if err != nil {
		t.Fatalf("Keys: %v", err)
	}
	slices.Sort(got)
	if want := []string{"app:access:u1:d1", "app:access:u1:d2"}; !slices.Equal(got, want) {
		t.Errorf("Keys = %v, want %v", got, want)
	}

	if got, err := s.Keys(ctx, "nonexistent:*"); err != nil || len(got) != 0 {
		t.Errorf("无匹配 Keys = %v, %v; want 空", got, err)
	}
}
