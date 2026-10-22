package jwtcodec_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/jwtcodec"
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
		Issuer:    "admin-api",
		Kind:      tokenx.KindAccess,
		IssuedAt:  now,
		ExpiresAt: now.Add(time.Hour),
	}
}

func TestRoundTrip(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	codec, err := jwtcodec.New(testSecret, jwtcodec.WithClock(clock))
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
		if encoded.LookupHash != "" {
			t.Fatal("JWT should not return LookupHash")
		}

		decoded, err := codec.Decode(context.Background(), encoded.Value)
		if err != nil {
			t.Fatalf("Decode(%s): %v", kind, err)
		}
		if decoded.Kind != kind || decoded.Claims == nil {
			t.Fatalf("decoded = %+v", decoded)
		}
		if decoded.Claims.UserID != "user-1" || decoded.Claims.SessionID != "session-1" || decoded.Claims.DeviceID != "device-1" {
			t.Fatalf("claims = %+v", decoded.Claims)
		}
		if decoded.Claims.TokenID == "" {
			t.Fatal("JWT should carry jti")
		}
	}
}

func TestExpired(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	codec, err := jwtcodec.New(testSecret, jwtcodec.WithClock(clock))
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

func TestWrongSecret(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	signer, err := jwtcodec.New(testSecret, jwtcodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}
	verifier, err := jwtcodec.New("fedcba9876543210fedcba9876543210", jwtcodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}

	encoded, err := signer.Encode(context.Background(), testClaims(clock.Now()))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := verifier.Decode(context.Background(), encoded.Value); !errors.Is(err, tokenx.ErrTokenInvalid) {
		t.Fatalf("wrong secret = %v, want ErrTokenInvalid", err)
	}
}

func TestRejectsMissingExpiration(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	codec, err := jwtcodec.New(testSecret, jwtcodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "u", "sid": "s", "did": "d",
	})
	token.Header["typ"] = "at+jwt"
	raw, err := token.SignedString([]byte(testSecret))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := codec.Decode(context.Background(), raw); !errors.Is(err, tokenx.ErrTokenInvalid) {
		t.Fatalf("missing exp = %v, want ErrTokenInvalid", err)
	}
}

func TestRejectsNoneAlgorithm(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1700000000, 0)}
	codec, err := jwtcodec.New(testSecret, jwtcodec.WithClock(clock))
	if err != nil {
		t.Fatal(err)
	}

	raw, err := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"sub": "u", "sid": "s", "did": "d",
		"exp": clock.Now().Add(time.Hour).Unix(),
	}).SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := codec.Decode(context.Background(), raw); !errors.Is(err, tokenx.ErrTokenInvalid) {
		t.Fatalf("alg=none = %v, want ErrTokenInvalid", err)
	}
}

func TestNewRejectsShortSecret(t *testing.T) {
	if _, err := jwtcodec.New("short"); err == nil {
		t.Fatal("short secret should be rejected")
	}
}
