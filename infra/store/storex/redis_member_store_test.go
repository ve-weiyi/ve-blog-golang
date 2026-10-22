package storex_test

import (
	"context"
	"math"
	"slices"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"

	"github.com/ve-weiyi/blog-cloud/infra/store/storex"
)

// seedMember 直接把过期时间戳写进 ZSET，绕开 Add 的 TTL 计算。
func seedMember(t *testing.T, mr *miniredis.Miniredis, key, member string, expireAt time.Time) {
	t.Helper()

	if _, err := mr.ZAdd(key, float64(expireAt.Unix()), member); err != nil {
		t.Fatalf("ZAdd(%s, %s): %v", key, member, err)
	}
}

func TestRedisMemberStore_AddAndMembers(t *testing.T) {
	ctx := context.Background()
	client, _ := newTestRedis(t)
	s := storex.NewRedisMemberStore(client)

	for _, m := range []string{"d1", "d2"} {
		if err := s.Add(ctx, "u1", m, time.Hour); err != nil {
			t.Fatalf("Add(%s): %v", m, err)
		}
	}

	got, err := s.Members(ctx, "u1")
	if err != nil {
		t.Fatalf("Members: %v", err)
	}
	if want := []string{"d1", "d2"}; !slices.Equal(got, want) {
		t.Errorf("Members = %v, want %v", got, want)
	}

	// 每个 member 的 TTL 相互独立：短 TTL 成员过期后，长 TTL 成员仍在。
	if err := s.Add(ctx, "u1", "short", 20*time.Millisecond); err != nil {
		t.Fatalf("Add(short): %v", err)
	}
	time.Sleep(50 * time.Millisecond)
	got, err = s.Members(ctx, "u1")
	if err != nil {
		t.Fatalf("Members: %v", err)
	}
	if want := []string{"d1", "d2"}; !slices.Equal(got, want) {
		t.Errorf("短 TTL 成员过期后 Members = %v, want %v", got, want)
	}
}

func TestRedisMemberStore_Members_MissingKey(t *testing.T) {
	ctx := context.Background()
	client, _ := newTestRedis(t)
	s := storex.NewRedisMemberStore(client)

	got, err := s.Members(ctx, "never-set")
	if err != nil {
		t.Fatalf("Members: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("Members = %v, want 空", got)
	}
}

func TestRedisMemberStore_Has(t *testing.T) {
	ctx := context.Background()
	client, mr := newTestRedis(t)
	s := storex.NewRedisMemberStore(client)

	if err := s.Add(ctx, "u1", "d1", time.Hour); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if has, err := s.Has(ctx, "u1", "d1"); err != nil || !has {
		t.Errorf("Has(d1) = %v, %v; want true", has, err)
	}
	if has, err := s.Has(ctx, "u1", "d9"); err != nil || has {
		t.Errorf("Has(d9) = %v, %v; want false", has, err)
	}

	// 过期 member 返回 false，并顺手从 ZSET 摘掉。
	seedMember(t, mr, "u1", "old", time.Now().Add(-time.Hour))
	if has, err := s.Has(ctx, "u1", "old"); err != nil || has {
		t.Errorf("Has(过期) = %v, %v; want false", has, err)
	}
	if got, _ := s.Members(ctx, "u1"); !slices.Equal(got, []string{"d1"}) {
		t.Errorf("Members = %v, want [d1]", got)
	}
}

func TestRedisMemberStore_Remove(t *testing.T) {
	ctx := context.Background()
	client, mr := newTestRedis(t)
	s := storex.NewRedisMemberStore(client)

	for _, m := range []string{"d1", "d2"} {
		if err := s.Add(ctx, "u1", m, time.Hour); err != nil {
			t.Fatalf("Add(%s): %v", m, err)
		}
	}

	if err := s.Remove(ctx, "u1", "d1"); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	if got, _ := s.Members(ctx, "u1"); !slices.Equal(got, []string{"d2"}) {
		t.Errorf("Remove(d1) 后 Members = %v, want [d2]", got)
	}

	// 不带 member 时删除整个 key。
	if err := s.Remove(ctx, "u1"); err != nil {
		t.Fatalf("Remove(all): %v", err)
	}
	if mr.Exists("u1") {
		t.Error("底层 key 仍在，want 被删除")
	}
}

func TestRedisMemberStore_Expiries(t *testing.T) {
	ctx := context.Background()
	client, mr := newTestRedis(t)
	s := storex.NewRedisMemberStore(client)

	old := time.Now().Add(-time.Hour)
	fresh := time.Now().Add(time.Hour).Truncate(time.Second)
	seedMember(t, mr, "u1", "old", old)
	seedMember(t, mr, "u1", "fresh", fresh)

	got, err := s.Expiries(ctx, "u1")
	if err != nil {
		t.Fatalf("Expiries: %v", err)
	}
	if len(got) != 1 || !got["fresh"].Equal(fresh) {
		t.Errorf("Expiries = %v, want 只剩 fresh", got)
	}

	// ttl <= 0 表示不过期，过期时刻为零值。
	if err := s.Add(ctx, "u2", "forever", 0); err != nil {
		t.Fatalf("Add(forever): %v", err)
	}
	got, err = s.Expiries(ctx, "u2")
	if err != nil {
		t.Fatalf("Expiries: %v", err)
	}
	if expireAt, ok := got["forever"]; !ok || !expireAt.IsZero() {
		t.Errorf("Expiries[forever] = %v (存在 %v), want 零值", expireAt, ok)
	}
}

// Members 按过期时刻升序，同刻按成员名，不过期的排最后。
func TestRedisMemberStore_Members_OrderMatchesExpiry(t *testing.T) {
	ctx := context.Background()
	client, mr := newTestRedis(t)
	s := storex.NewRedisMemberStore(client)

	seedMember(t, mr, "u1", "third", time.Now().Add(2*time.Hour))
	seedMember(t, mr, "u1", "first", time.Now().Add(time.Hour))
	seedMember(t, mr, "u1", "second", time.Now().Add(time.Hour))
	if _, err := mr.ZAdd("u1", math.Inf(1), "immortal"); err != nil {
		t.Fatalf("ZAdd(immortal): %v", err)
	}

	got, err := s.Members(ctx, "u1")
	if err != nil {
		t.Fatalf("Members: %v", err)
	}
	if want := []string{"first", "second", "third", "immortal"}; !slices.Equal(got, want) {
		t.Errorf("Members = %v, want %v", got, want)
	}
}
