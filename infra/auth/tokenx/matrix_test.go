package tokenx_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/hmaccodec"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/jwtcodec"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/opaque"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/store/memstore"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/store/redisstore"
)

// TestMatrix 覆盖 codec × store × LoginMode 的主流程：签发、校验、轮换、宽限重放、撤销。
func TestMatrix(t *testing.T) {
	codecs := []struct {
		name string
		new  func(*testing.T, tokenx.Clock) tokenx.Codec
	}{
		{"opaque", func(*testing.T, tokenx.Clock) tokenx.Codec { return opaque.New() }},
		{"jwt-hs256", func(t *testing.T, clock tokenx.Clock) tokenx.Codec {
			codec, err := jwtcodec.New(testSecret, jwtcodec.WithClock(clock))
			if err != nil {
				t.Fatal(err)
			}
			return codec
		}},
		{"hmac", func(t *testing.T, clock tokenx.Clock) tokenx.Codec {
			codec, err := hmaccodec.New(testSecret, hmaccodec.WithClock(clock))
			if err != nil {
				t.Fatal(err)
			}
			return codec
		}},
	}

	stores := []struct {
		name string
		new  func(*testing.T, tokenx.Clock) tokenx.SessionStore
	}{
		{"memstore", func(t *testing.T, clock tokenx.Clock) tokenx.SessionStore {
			return memstore.NewWithClock(clock)
		}},
		{"redisstore", func(t *testing.T, _ tokenx.Clock) tokenx.SessionStore {
			mr := miniredis.RunT(t)
			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
			t.Cleanup(func() { _ = client.Close() })
			store, err := redisstore.New(client, "tokenx:matrix")
			if err != nil {
				t.Fatal(err)
			}
			return store
		}},
	}

	modes := []struct {
		name string
		mode tokenx.LoginMode
	}{
		{"multi", tokenx.MultiPoint},
		{"single", tokenx.SinglePoint},
	}

	for _, codecCase := range codecs {
		for _, storeCase := range stores {
			for _, modeCase := range modes {
				t.Run(codecCase.name+"/"+storeCase.name+"/"+modeCase.name, func(t *testing.T) {
					clock := newFakeClock()
					svc := newMatrixService(t, codecCase.new(t, clock), storeCase.new(t, clock), modeCase.mode, clock)
					ctx := context.Background()

					pair := issue(t, svc, "u", "d")
					if _, err := svc.Validate(ctx, pair.AccessToken); err != nil {
						t.Fatalf("Validate: %v", err)
					}

					refreshed, err := svc.Refresh(ctx, refreshRequest(pair))
					if err != nil {
						t.Fatalf("Refresh: %v", err)
					}
					if _, err := svc.Validate(ctx, pair.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
						t.Fatalf("旧 access = %v, want ErrSessionNotFound", err)
					}

					replayed, err := svc.Refresh(ctx, refreshRequest(pair))
					if err != nil {
						t.Fatalf("grace 内重放: %v", err)
					}
					if replayed.AccessToken != refreshed.AccessToken {
						t.Fatal("grace 内重放应返回同一 pair")
					}

					if err := svc.RevokeAll(ctx, "u"); err != nil {
						t.Fatalf("RevokeAll: %v", err)
					}
					if _, err := svc.Validate(ctx, refreshed.AccessToken); !errors.Is(err, tokenx.ErrSessionNotFound) {
						t.Fatalf("撤销后 access = %v, want ErrSessionNotFound", err)
					}
				})
			}
		}
	}
}

func newMatrixService(t *testing.T, codec tokenx.Codec, store tokenx.SessionStore, mode tokenx.LoginMode, clock tokenx.Clock) tokenx.Service {
	t.Helper()

	svc, err := tokenx.New(tokenx.Config{
		AccessCodec:        codec,
		RefreshCodec:       codec,
		Store:              store,
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
