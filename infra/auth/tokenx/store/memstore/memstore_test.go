package memstore_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/store/memstore"
)

type fakeClock struct {
	now time.Time
}

func (c *fakeClock) Now() time.Time { return c.now }

func newSession(id, userID, deviceID, familyID string, now time.Time) tokenx.Session {
	return tokenx.Session{
		ID:                id,
		UserID:            userID,
		DeviceID:          deviceID,
		FamilyID:          familyID,
		AccessCodec:       "opaque-v3",
		RefreshCodec:      "opaque-v3",
		AccessLookupHash:  "at-" + id,
		RefreshLookupHash: "rt-" + id,
		RefreshTokenHash:  "rh-" + id,
		CreatedAt:         now,
		LastUsedAt:        now,
		AccessExpiresAt:   now.Add(time.Hour),
		RefreshExpiresAt:  now.Add(24 * time.Hour),
		AbsoluteExpiresAt: now.Add(30 * 24 * time.Hour),
	}
}

func TestCreateMultiPointReplacesSameDevice(t *testing.T) {
	ctx := context.Background()
	now := time.Unix(1700000000, 0)
	store := memstore.NewWithClock(&fakeClock{now: now})

	if err := store.Create(ctx, newSession("s1", "u", "iphone", "f1", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	if err := store.Create(ctx, newSession("s2", "u", "ipad", "f2", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	if err := store.Create(ctx, newSession("s3", "u", "iphone", "f3", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}

	if _, err := store.GetBySession(ctx, "u", "s1"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("同设备旧会话 = %v, want ErrSessionNotFound", err)
	}
	if _, err := store.GetByTokenHash(ctx, "at-s1"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("旧会话 token 索引 = %v, want ErrSessionNotFound", err)
	}
	if _, err := store.GetBySession(ctx, "u", "s2"); err != nil {
		t.Fatalf("另一设备会话: %v", err)
	}
	if _, err := store.GetBySession(ctx, "u", "s3"); err != nil {
		t.Fatalf("同设备新会话: %v", err)
	}
}

func TestCreateSinglePointClearsAll(t *testing.T) {
	ctx := context.Background()
	now := time.Unix(1700000000, 0)
	store := memstore.NewWithClock(&fakeClock{now: now})

	if err := store.Create(ctx, newSession("s1", "u", "iphone", "f1", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	if err := store.Create(ctx, newSession("s2", "u", "ipad", "f2", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	if err := store.Create(ctx, newSession("s3", "u", "ipad", "f3", now), tokenx.SinglePoint); err != nil {
		t.Fatal(err)
	}

	sessions, err := store.ListByUser(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 || sessions[0].ID != "s3" {
		t.Fatalf("单点登录后会话 = %+v, want s3", sessions)
	}
	if _, err := store.GetByTokenHash(ctx, "at-s1"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("旧 token 索引 = %v, want ErrSessionNotFound", err)
	}
}

func TestRotateAndLookupRotation(t *testing.T) {
	ctx := context.Background()
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	store := memstore.NewWithClock(clock)

	if err := store.Create(ctx, newSession("s1", "u", "d", "f1", clock.now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}

	next := newSession("s2", "u", "d", "ignored", clock.now)
	err := store.Rotate(ctx, newSession("s1", "u", "d", "f1", clock.now), next, tokenx.RotateOptions{
		GracePayload: []byte("sealed-pair"),
		RotationTTL:  time.Hour,
		GraceTTL:     30 * time.Second,
	})
	if err != nil {
		t.Fatalf("Rotate: %v", err)
	}

	if _, err := store.GetBySession(ctx, "u", "s1"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("旧会话 = %v, want ErrSessionNotFound", err)
	}
	if _, err := store.GetByTokenHash(ctx, "at-s1"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("旧 token 索引 = %v, want ErrSessionNotFound", err)
	}
	current, err := store.GetByTokenHash(ctx, "at-s2")
	if err != nil {
		t.Fatalf("新 token 索引: %v", err)
	}
	if current.ID != "s2" || current.FamilyID != "f1" {
		t.Fatalf("新会话 = %+v, want s2/f1", current)
	}

	rotation, found, err := store.LookupRotation(ctx, "u", "rh-s1")
	if err != nil || !found {
		t.Fatalf("LookupRotation = (%+v, %v, %v)", rotation, found, err)
	}
	if rotation.FamilyID != "f1" || string(rotation.GracePayload) != "sealed-pair" {
		t.Fatalf("rotation = %+v", rotation)
	}

	if _, found, err := store.LookupRotation(ctx, "other", "rh-s1"); err != nil || found {
		t.Fatalf("错误用户的 rotation = (found %v, err %v)", found, err)
	}

	clock.now = clock.now.Add(31 * time.Second)
	rotation, found, err = store.LookupRotation(ctx, "u", "rh-s1")
	if err != nil || !found {
		t.Fatalf("grace 过期后仍应找得到轮换记录: (%+v, %v, %v)", rotation, found, err)
	}
	if len(rotation.GracePayload) != 0 {
		t.Fatal("grace 过期后不应返回 payload")
	}

	clock.now = clock.now.Add(time.Hour)
	if _, found, err := store.LookupRotation(ctx, "u", "rh-s1"); err != nil || found {
		t.Fatalf("rotation 过期后 = (found %v, err %v), want 未找到", found, err)
	}
}

func TestRotateStaleSession(t *testing.T) {
	ctx := context.Background()
	now := time.Unix(1700000000, 0)
	store := memstore.NewWithClock(&fakeClock{now: now})

	if err := store.Create(ctx, newSession("s1", "u", "d", "f1", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	old := newSession("s1", "u", "d", "f1", now)
	if err := store.Rotate(ctx, old, newSession("s2", "u", "d", "f1", now), tokenx.RotateOptions{}); err != nil {
		t.Fatal(err)
	}
	if err := store.Rotate(ctx, old, newSession("s3", "u", "d", "f1", now), tokenx.RotateOptions{}); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("重复轮换 = %v, want ErrSessionNotFound", err)
	}
}

func TestRevokeFamily(t *testing.T) {
	ctx := context.Background()
	now := time.Unix(1700000000, 0)
	store := memstore.NewWithClock(&fakeClock{now: now})

	if err := store.Create(ctx, newSession("s1", "u", "d", "f1", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	old := newSession("s1", "u", "d", "f1", now)
	if err := store.Rotate(ctx, old, newSession("s2", "u", "d", "f1", now), tokenx.RotateOptions{}); err != nil {
		t.Fatal(err)
	}
	if err := store.RevokeFamily(ctx, "u", "f1"); err != nil {
		t.Fatal(err)
	}

	if _, err := store.GetByTokenHash(ctx, "rt-s2"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("撤销 family 后 token = %v, want ErrSessionNotFound", err)
	}
	sessions, err := store.ListByUser(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 0 {
		t.Fatalf("撤销 family 后会话 = %+v, want 空", sessions)
	}
}

func TestDeleteAndList(t *testing.T) {
	ctx := context.Background()
	now := time.Unix(1700000000, 0)
	store := memstore.NewWithClock(&fakeClock{now: now})

	if err := store.Create(ctx, newSession("s1", "u", "iphone", "f1", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	if err := store.Create(ctx, newSession("s2", "u", "ipad", "f2", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}

	if err := store.Delete(ctx, "u2", "s1"); err != nil {
		t.Fatalf("错误用户删除应无副作用: %v", err)
	}
	if _, err := store.GetBySession(ctx, "u", "s1"); err != nil {
		t.Fatalf("错误用户不应删除他人会话: %v", err)
	}

	if err := store.Delete(ctx, "u", "s1"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if err := store.Delete(ctx, "u", "s1"); err != nil {
		t.Fatalf("重复删除应幂等: %v", err)
	}
	sessions, err := store.ListByUser(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 || sessions[0].ID != "s2" {
		t.Fatalf("删除后会话 = %+v", sessions)
	}

	if err := store.DeleteByDevice(ctx, "u", "ipad"); err != nil {
		t.Fatal(err)
	}
	sessions, err = store.ListByUser(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 0 {
		t.Fatalf("按设备删除后会话数 = %d, want 0", len(sessions))
	}
}

func TestDeleteByUser(t *testing.T) {
	ctx := context.Background()
	now := time.Unix(1700000000, 0)
	store := memstore.NewWithClock(&fakeClock{now: now})

	if err := store.Create(ctx, newSession("s1", "u", "iphone", "f1", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	if err := store.Create(ctx, newSession("s2", "u2", "ipad", "f2", now), tokenx.MultiPoint); err != nil {
		t.Fatal(err)
	}
	if err := store.DeleteByUser(ctx, "u"); err != nil {
		t.Fatal(err)
	}

	if _, err := store.GetBySession(ctx, "u", "s1"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("删除用户后会话 = %v", err)
	}
	if _, err := store.GetBySession(ctx, "u2", "s2"); err != nil {
		t.Fatalf("其他用户会话不应受影响: %v", err)
	}
}

func TestGetMissing(t *testing.T) {
	store := memstore.New()
	if _, err := store.GetBySession(context.Background(), "u", "missing"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("GetBySession missing = %v, want ErrSessionNotFound", err)
	}
	if _, err := store.GetByTokenHash(context.Background(), "missing"); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("GetByTokenHash missing = %v, want ErrSessionNotFound", err)
	}
}
