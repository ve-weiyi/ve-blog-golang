package tokenx_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/hmaccodec"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/jwtcodec"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/opaque"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/store/memstore"
)

const (
	testSecret  = "0123456789abcdef0123456789abcdef"
	accessTTL   = 120 * time.Minute
	idleTTL     = 7 * 24 * time.Hour
	absoluteTTL = 30 * 24 * time.Hour
	graceTTL    = tokenx.GraceWindow
)

type fakeClock struct {
	mu  sync.Mutex
	now time.Time
}

func newFakeClock() *fakeClock {
	return &fakeClock{now: time.Unix(1700000000, 0)}
}

func (c *fakeClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *fakeClock) Advance(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(d)
}

func newService(t *testing.T, clock tokenx.Clock, access, refresh tokenx.Codec, mode tokenx.LoginMode) tokenx.Service {
	t.Helper()

	svc, err := tokenx.New(tokenx.Config{
		AccessCodec:        access,
		RefreshCodec:       refresh,
		Store:              memstore.NewWithClock(clock),
		LoginMode:          mode,
		AccessTTL:          accessTTL,
		RefreshIdleTTL:     idleTTL,
		RefreshAbsoluteTTL: absoluteTTL,
		Clock:              clock,
	})
	if err != nil {
		t.Fatal(err)
	}
	return svc
}

func newOpaqueService(t *testing.T, clock tokenx.Clock, mode tokenx.LoginMode) tokenx.Service {
	t.Helper()
	return newService(t, clock, opaque.New(), opaque.New(), mode)
}

func newJWTService(t *testing.T, clock tokenx.Clock, mode tokenx.LoginMode) tokenx.Service {
	t.Helper()
	codec, err := jwtcodec.New(testSecret, jwtcodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}
	return newService(t, clock, codec, codec, mode)
}

func newHMACService(t *testing.T, clock tokenx.Clock, mode tokenx.LoginMode) tokenx.Service {
	t.Helper()
	codec, err := hmaccodec.New(testSecret, hmaccodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}
	return newService(t, clock, codec, codec, mode)
}

func refreshRequest(pair *tokenx.TokenPair) tokenx.RefreshRequest {
	return tokenx.RefreshRequest{
		RefreshToken: pair.RefreshToken,
		UserID:       pair.UserID,
		DeviceID:     pair.DeviceID,
	}
}

func issue(t *testing.T, svc tokenx.Service, userID, deviceID string) *tokenx.TokenPair {
	t.Helper()
	pair, err := svc.Issue(context.Background(), tokenx.IssueRequest{
		UserID:   userID,
		DeviceID: deviceID,
		Issuer:   "blog-api",
	})
	if err != nil {
		t.Fatalf("Issue: %v", err)
	}
	return pair
}

func TestIssueValidate(t *testing.T) {
	cases := []struct {
		name string
		new  func(*testing.T, tokenx.Clock, tokenx.LoginMode) tokenx.Service
	}{
		{"opaque", newOpaqueService},
		{"jwt", newJWTService},
		{"hmac", newHMACService},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			clock := newFakeClock()
			svc := tc.new(t, clock, tokenx.MultiPoint)
			pair := issue(t, svc, "user-1", "device-1")

			claims, err := svc.Validate(context.Background(), pair.AccessToken)
			if err != nil {
				t.Fatalf("Validate: %v", err)
			}
			if claims.UserID != "user-1" || claims.DeviceID != "device-1" {
				t.Fatalf("claims identity = (%s, %s)", claims.UserID, claims.DeviceID)
			}
			if claims.SessionID != pair.SessionID || claims.Kind != tokenx.KindAccess {
				t.Fatalf("claims session/kind = (%s, %s)", claims.SessionID, claims.Kind)
			}
			if claims.Issuer != "blog-api" {
				t.Fatalf("claims issuer = %s", claims.Issuer)
			}
			if pair.AccessExpiresIn != int64(accessTTL/time.Second) {
				t.Fatalf("access ttl = %d, want %d", pair.AccessExpiresIn, int64(accessTTL/time.Second))
			}
		})
	}
}

func TestClientIDCarriedThroughRefresh(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair, err := svc.Issue(ctx, tokenx.IssueRequest{
		UserID:   "u",
		DeviceID: "d",
		ClientID: "blog-app",
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := svc.Refresh(ctx, refreshRequest(pair)); err != nil {
		t.Fatal(err)
	}

	sessions, err := svc.ListSessions(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 || sessions[0].ClientID != "blog-app" {
		t.Fatalf("会话 = %+v, want client_id=blog-app", sessions)
	}
}

func TestMultiPointKeepsOtherDevices(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	first := issue(t, svc, "u", "iphone")
	second := issue(t, svc, "u", "ipad")
	replaced := issue(t, svc, "u", "iphone")

	if _, err := svc.Validate(ctx, first.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("同设备旧 access = %v, want ErrSessionNotFound", err)
	}
	if _, err := svc.Validate(ctx, replaced.AccessToken); err != nil {
		t.Fatalf("同设备新 access: %v", err)
	}
	if _, err := svc.Validate(ctx, second.AccessToken); err != nil {
		t.Fatalf("其他设备 access 应保留: %v", err)
	}
}

func TestSinglePointRevokesAllDevices(t *testing.T) {
	clock := newFakeClock()
	svc := newJWTService(t, clock, tokenx.SinglePoint)
	ctx := context.Background()

	first := issue(t, svc, "u", "iphone")
	second := issue(t, svc, "u", "ipad")

	if _, err := svc.Validate(ctx, first.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("旧设备 access = %v, want ErrSessionNotFound", err)
	}
	if _, err := svc.Validate(ctx, second.AccessToken); err != nil {
		t.Fatalf("新设备 access: %v", err)
	}
}

func TestRevokeIsImmediateAndIdempotent(t *testing.T) {
	clock := newFakeClock()
	svc := newJWTService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	if err := svc.Revoke(ctx, "u", pair.SessionID); err != nil {
		t.Fatalf("Revoke: %v", err)
	}
	if _, err := svc.Validate(ctx, pair.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("撤销后 access = %v, want ErrSessionNotFound", err)
	}
	if err := svc.Revoke(ctx, "u", pair.SessionID); err != nil {
		t.Fatalf("重复撤销应幂等: %v", err)
	}
}

func TestRevokeDeviceAndAll(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	iphone := issue(t, svc, "u", "iphone")
	ipad := issue(t, svc, "u", "ipad")

	if err := svc.RevokeDevice(ctx, "u", "iphone"); err != nil {
		t.Fatalf("RevokeDevice: %v", err)
	}
	if _, err := svc.Validate(ctx, iphone.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("撤销设备后 access = %v", err)
	}
	if _, err := svc.Validate(ctx, ipad.AccessToken); err != nil {
		t.Fatalf("另一设备 access 应保留: %v", err)
	}

	if err := svc.RevokeAll(ctx, "u"); err != nil {
		t.Fatalf("RevokeAll: %v", err)
	}
	if _, err := svc.Validate(ctx, ipad.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("RevokeAll 后 access = %v", err)
	}
}

func TestRefreshRotates(t *testing.T) {
	clock := newFakeClock()
	svc := newJWTService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	refreshed, err := svc.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if refreshed.UserID != pair.UserID || refreshed.DeviceID != pair.DeviceID {
		t.Fatalf("refresh identity = (%s, %s)", refreshed.UserID, refreshed.DeviceID)
	}
	if refreshed.SessionID == pair.SessionID {
		t.Fatal("刷新应生成新的 session id")
	}

	if _, err := svc.Validate(ctx, pair.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("旧 access = %v, want ErrSessionNotFound", err)
	}
	if _, err := svc.Validate(ctx, refreshed.AccessToken); err != nil {
		t.Fatalf("新 access: %v", err)
	}
}

func TestRefreshWithinGraceReturnsSamePair(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	first, err := svc.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatalf("first refresh: %v", err)
	}

	clock.Advance(graceTTL - time.Second)
	replayed, err := svc.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatalf("grace 内重放: %v", err)
	}
	if replayed.AccessToken != first.AccessToken || replayed.RefreshToken != first.RefreshToken {
		t.Fatal("grace 内重放应返回同一 pair")
	}
	if _, err := svc.Validate(ctx, first.AccessToken); err != nil {
		t.Fatalf("grace 内重放不应撤销新会话: %v", err)
	}
}

func TestRefreshOutsideGraceRevokesFamily(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	refreshed, err := svc.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}

	clock.Advance(graceTTL + time.Second)
	if _, err := svc.Refresh(ctx, refreshRequest(pair)); !errors.Is(err, tokenx.ErrRefreshReused) {
		t.Fatalf("grace 外重放 = %v, want ErrRefreshReused", err)
	}
	if _, err := svc.Validate(ctx, refreshed.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("family 撤销后 access = %v, want ErrSessionNotFound", err)
	}
	if _, err := svc.Refresh(ctx, refreshRequest(refreshed)); !errors.Is(err, tokenx.ErrRefreshReused) {
		t.Fatalf("撤销后新 refresh = %v, want ErrRefreshReused", err)
	}
}

func TestRefreshRejectsMismatchedIdentity(t *testing.T) {
	clock := newFakeClock()
	svc := newJWTService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")

	wrongUser := refreshRequest(pair)
	wrongUser.UserID = "other"
	if _, err := svc.Refresh(ctx, wrongUser); !errors.Is(err, tokenx.ErrTokenInvalid) {
		t.Fatalf("错误 userID = %v, want ErrTokenInvalid", err)
	}

	wrongDevice := refreshRequest(pair)
	wrongDevice.DeviceID = "other"
	if _, err := svc.Refresh(ctx, wrongDevice); !errors.Is(err, tokenx.ErrTokenInvalid) {
		t.Fatalf("错误 deviceID = %v, want ErrTokenInvalid", err)
	}

	if _, err := svc.Validate(ctx, pair.AccessToken); err != nil {
		t.Fatalf("失败的身份校验不应轮换会话: %v", err)
	}
}

func TestRefreshForgedTokenDoesNotRevokeFamily(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	refreshed, err := svc.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatal(err)
	}

	forged, err := opaque.New().Encode(ctx, tokenx.Claims{Kind: tokenx.KindRefresh})
	if err != nil {
		t.Fatal(err)
	}
	_, err = svc.Refresh(ctx, tokenx.RefreshRequest{
		RefreshToken: forged.Value,
		UserID:       pair.UserID,
		DeviceID:     pair.DeviceID,
	})
	if !errors.Is(err, tokenx.ErrRefreshReused) {
		t.Fatalf("伪造 refresh = %v, want ErrRefreshReused", err)
	}
	if _, err := svc.Validate(ctx, refreshed.AccessToken); err != nil {
		t.Fatalf("伪造 token 不应撤销合法会话: %v", err)
	}
}

func TestRefreshAfterLogoutDoesNotRevokeNewLogin(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	oldPair := issue(t, svc, "u", "d")
	if err := svc.RevokeDevice(ctx, "u", "d"); err != nil {
		t.Fatal(err)
	}
	newPair := issue(t, svc, "u", "d")

	if _, err := svc.Refresh(ctx, refreshRequest(oldPair)); !errors.Is(err, tokenx.ErrRefreshReused) {
		t.Fatalf("登出前 refresh = %v, want ErrRefreshReused", err)
	}
	if _, err := svc.Validate(ctx, newPair.AccessToken); err != nil {
		t.Fatalf("新登录不应被误撤销: %v", err)
	}
}

func TestAccessAndRefreshExpiry(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	clock.Advance(accessTTL + tokenx.ClockLeeway + time.Second)
	if _, err := svc.Validate(ctx, pair.AccessToken); !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("过期 access = %v, want ErrTokenExpired", err)
	}

	refreshed, err := svc.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatalf("access 过期后仍应能刷新: %v", err)
	}

	clock.Advance(idleTTL + tokenx.ClockLeeway + time.Second)
	if _, err := svc.Refresh(ctx, refreshRequest(refreshed)); !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("idle 过期 refresh = %v, want ErrTokenExpired", err)
	}
}

func TestRefreshAbsoluteLifetime(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	step := 6 * 24 * time.Hour
	for i := 0; i < 4; i++ {
		clock.Advance(step)
		refreshed, err := svc.Refresh(ctx, refreshRequest(pair))
		if err != nil {
			t.Fatalf("第 %d 次刷新: %v", i+1, err)
		}
		pair = refreshed
	}

	clock.Advance(step + tokenx.ClockLeeway + time.Minute)
	if _, err := svc.Refresh(ctx, refreshRequest(pair)); !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("absolute 过期 refresh = %v, want ErrTokenExpired", err)
	}
}

func TestRefreshRejectsAbsoluteExpiry(t *testing.T) {
	clock := newFakeClock()
	svc, err := tokenx.New(tokenx.Config{
		AccessCodec:        opaque.New(),
		Store:              memstore.NewWithClock(clock),
		AccessTTL:          accessTTL,
		RefreshIdleTTL:     idleTTL,
		RefreshAbsoluteTTL: time.Hour,
		Clock:              clock,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	clock.Advance(time.Hour + time.Second)
	if _, err := svc.Refresh(ctx, refreshRequest(pair)); !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("absolute 后刷新 = %v, want ErrTokenExpired", err)
	}
}

func TestRotationRecordCoversGraceWindow(t *testing.T) {
	clock := newFakeClock()
	store := memstore.NewWithClock(clock)
	svc, err := tokenx.New(tokenx.Config{
		AccessCodec:        opaque.New(),
		Store:              store,
		AccessTTL:          accessTTL,
		RefreshIdleTTL:     time.Second,
		RefreshAbsoluteTTL: time.Hour,
		Clock:              clock,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	clock.Advance(10 * time.Second) // refresh 已过期 9 秒，仍在 ClockLeeway 内
	if _, err := svc.Refresh(ctx, refreshRequest(pair)); err != nil {
		t.Fatalf("leeway 内刷新: %v", err)
	}

	hash := tokenx.HashToken(pair.RefreshToken)
	if _, found, err := store.LookupRotation(ctx, "u", hash); err != nil || !found {
		t.Fatalf("轮换记录应存在: found=%v err=%v", found, err)
	}

	clock.Advance(graceTTL + time.Second)
	if _, found, err := store.LookupRotation(ctx, "u", hash); err != nil || found {
		t.Fatalf("轮换记录应在 grace 窗口后过期: found=%v err=%v", found, err)
	}
}

func TestListSessionsFiltersExpired(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	issue(t, svc, "u", "iphone")
	issue(t, svc, "u", "ipad")

	sessions, err := svc.ListSessions(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 2 {
		t.Fatalf("会话数 = %d, want 2", len(sessions))
	}

	if err := svc.RevokeDevice(ctx, "u", "iphone"); err != nil {
		t.Fatal(err)
	}
	sessions, err = svc.ListSessions(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 || sessions[0].DeviceID != "ipad" {
		t.Fatalf("撤销后会话 = %+v, want 只剩 ipad", sessions)
	}

	clock.Advance(absoluteTTL + tokenx.ClockLeeway + time.Second)
	sessions, err = svc.ListSessions(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 0 {
		t.Fatalf("过期后会话数 = %d, want 0", len(sessions))
	}
}

func TestOpaqueTamperedTokenIsRejected(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	last := pair.AccessToken[len(pair.AccessToken)-1]
	replacement := byte('A')
	if last == 'A' {
		replacement = 'B'
	}
	tampered := pair.AccessToken[:len(pair.AccessToken)-1] + string(replacement)

	if _, err := svc.Validate(ctx, tampered); !errors.Is(err, tokenx.ErrSessionNotFound) {
		t.Fatalf("篡改的 opaque token = %v, want ErrSessionNotFound", err)
	}
}

func TestDefaultDeviceID(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)

	pair := issue(t, svc, "u", "")
	claims, err := svc.Validate(context.Background(), pair.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if claims.DeviceID != tokenx.DefaultDeviceID {
		t.Fatalf("device id = %q, want %q", claims.DeviceID, tokenx.DefaultDeviceID)
	}
}

func TestConcurrentSinglePointKeepsOneSession(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.SinglePoint)
	ctx := context.Background()

	const n = 16
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = svc.Issue(ctx, tokenx.IssueRequest{UserID: "u", DeviceID: "d"})
		}()
	}
	wg.Wait()

	sessions, err := svc.ListSessions(ctx, "u")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 {
		t.Fatalf("并发单点登录后会话数 = %d, want 1", len(sessions))
	}
}

func TestConcurrentRefreshWithinGraceReturnsSamePair(t *testing.T) {
	clock := newFakeClock()
	svc := newOpaqueService(t, clock, tokenx.MultiPoint)
	ctx := context.Background()

	pair := issue(t, svc, "u", "d")
	results := make([]*tokenx.TokenPair, 2)
	errs := make([]error, 2)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i], errs[i] = svc.Refresh(ctx, refreshRequest(pair))
		}(i)
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			t.Fatalf("并发刷新 %d: %v", i, err)
		}
	}
	if results[0].AccessToken != results[1].AccessToken ||
		results[0].RefreshToken != results[1].RefreshToken {
		t.Fatal("grace 内并发刷新应返回同一 pair")
	}
	if _, err := svc.Validate(ctx, results[0].AccessToken); err != nil {
		t.Fatalf("并发刷新后 access: %v", err)
	}
}

func TestNewValidation(t *testing.T) {
	base := tokenx.Config{
		AccessCodec: opaque.New(),
		Store:       memstore.New(),
	}
	if _, err := tokenx.New(base); err != nil {
		t.Fatalf("最小配置应可用: %v", err)
	}

	if _, err := tokenx.New(tokenx.Config{Store: memstore.New()}); err == nil {
		t.Fatal("缺少 AccessCodec 应被拒绝")
	}
	if _, err := tokenx.New(tokenx.Config{AccessCodec: opaque.New()}); err == nil {
		t.Fatal("缺少 Store 应被拒绝")
	}
}

func TestGraceAcrossInstances(t *testing.T) {
	clock := newFakeClock()
	store := memstore.NewWithClock(clock)
	newSvc := func() tokenx.Service {
		svc, err := tokenx.New(tokenx.Config{
			AccessCodec:        opaque.New(),
			Store:              store,
			AccessTTL:          accessTTL,
			RefreshIdleTTL:     idleTTL,
			RefreshAbsoluteTTL: absoluteTTL,
			Clock:              clock,
		})
		if err != nil {
			t.Fatal(err)
		}
		return svc
	}
	first, second := newSvc(), newSvc()
	ctx := context.Background()

	pair := issue(t, first, "u", "d")
	refreshed, err := first.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}

	// 密钥由旧 refresh token 派生，另一实例无需共享配置即可解开 grace payload。
	replayed, err := second.Refresh(ctx, refreshRequest(pair))
	if err != nil {
		t.Fatalf("跨实例重放: %v", err)
	}
	if replayed.AccessToken != refreshed.AccessToken || replayed.RefreshToken != refreshed.RefreshToken {
		t.Fatal("跨实例重放应返回同一 pair")
	}
}
