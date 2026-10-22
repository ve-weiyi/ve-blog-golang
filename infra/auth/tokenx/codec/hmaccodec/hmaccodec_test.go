package hmaccodec_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/hmaccodec"
)

const testSecret = "0123456789abcdef0123456789abcdef"

type fakeClock struct {
	now time.Time
}

func (c *fakeClock) Now() time.Time { return c.now }

func testClaims(now time.Time) tokenx.Claims {
	return tokenx.Claims{
		UserID:    "user-1",
		DeviceID:  "device-1",
		SessionID: "session-1",
		Issuer:    "app-api",
		Kind:      tokenx.KindAccess,
		IssuedAt:  now,
		ExpiresAt: now.Add(time.Hour),
	}
}

func TestRoundTrip(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	codec, err := hmaccodec.New(testSecret, hmaccodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}

	for _, kind := range []tokenx.TokenKind{tokenx.KindAccess, tokenx.KindRefresh} {
		claims := testClaims(clock.Now())
		claims.Kind = kind
		encoded, err := codec.Encode(context.Background(), claims)
		if err != nil {
			t.Fatalf("Encode(%s): %v", kind, err)
		}
		if !strings.HasPrefix(encoded.Value, "v1.") || encoded.LookupHash != "" {
			t.Fatalf("encoded = %+v", encoded)
		}

		decoded, err := codec.Decode(context.Background(), encoded.Value)
		if err != nil {
			t.Fatalf("Decode(%s): %v", kind, err)
		}
		if decoded.Kind != kind || decoded.Claims == nil {
			t.Fatalf("decoded = %+v", decoded)
		}
		if decoded.Claims.UserID != "user-1" || decoded.Claims.SessionID != "session-1" {
			t.Fatalf("claims = %+v", decoded.Claims)
		}
		if decoded.Claims.TokenID == "" {
			t.Fatal("token id should be generated")
		}
	}
}

func TestDecodeRejects(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	codec, err := hmaccodec.New(testSecret, hmaccodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}
	other, err := hmaccodec.New("fedcba9876543210fedcba9876543210", hmaccodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}

	encoded, err := codec.Encode(context.Background(), testClaims(clock.Now()))
	if err != nil {
		t.Fatal(err)
	}
	parts := strings.Split(encoded.Value, ".")
	tampered := parts[0] + "." + parts[1] + "." + "AAAA"

	cases := []struct {
		name  string
		codec tokenx.Codec
		raw   string
		want  error
	}{
		{"wrong secret", other, encoded.Value, tokenx.ErrTokenInvalid},
		{"tampered mac", codec, tampered, tokenx.ErrTokenInvalid},
		{"wrong version", codec, "v2." + parts[1] + "." + parts[2], tokenx.ErrTokenInvalid},
		{"malformed", codec, "not-a-token", tokenx.ErrTokenInvalid},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := tc.codec.Decode(context.Background(), tc.raw); !errors.Is(err, tc.want) {
				t.Fatalf("Decode = %v, want %v", err, tc.want)
			}
		})
	}
}

func TestExpired(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	codec, err := hmaccodec.New(testSecret, hmaccodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}

	claims := testClaims(clock.Now())
	claims.ExpiresAt = clock.Now().Add(time.Minute)
	encoded, err := codec.Encode(context.Background(), claims)
	if err != nil {
		t.Fatal(err)
	}

	clock.now = clock.now.Add(2 * time.Minute)
	if _, err := codec.Decode(context.Background(), encoded.Value); !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("expired = %v, want ErrTokenExpired", err)
	}
}

func TestNewRejectsShortSecret(t *testing.T) {
	if _, err := hmaccodec.New("short"); err == nil {
		t.Fatal("short secret should be rejected")
	}
}
