// Package hmaccodec 实现自定义 HMAC-SHA256 信封，作为可选 codec。
//
// 格式：v1.<base64url(claims JSON)>.<base64url(HMAC)>；不使用 JWT。
package hmaccodec

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
)

const (
	version   = "v1"
	name      = "hmac-v1"
	minSecret = 32
)

// Codec 是 HMAC-SHA256 自包含令牌编码器。
type Codec struct {
	secret []byte
	clock  tokenx.Clock
	leeway time.Duration
}

// Option 是 HMAC Codec 的可选配置。
type Option func(*Codec)

// WithClock 注入时间源。
func WithClock(clock tokenx.Clock) Option {
	return func(c *Codec) {
		if clock != nil {
			c.clock = clock
		}
	}
}

// WithLeeway 设置时钟偏差容忍度。
func WithLeeway(leeway time.Duration) Option {
	return func(c *Codec) {
		if leeway >= 0 {
			c.leeway = leeway
		}
	}
}

// New 创建 HMAC-SHA256 Codec。
func New(secret string, opts ...Option) (tokenx.Codec, error) {
	if len(secret) < minSecret {
		return nil, fmt.Errorf("hmaccodec: secret must be at least %d bytes", minSecret)
	}

	codec := &Codec{
		secret: []byte(secret),
		clock:  tokenx.SystemClock{},
		leeway: tokenx.ClockLeeway,
	}
	for _, opt := range opts {
		opt(codec)
	}
	return codec, nil
}

// Name 返回方案标识。
func (*Codec) Name() string { return name }

type wireClaims struct {
	UserID    string `json:"sub"`
	DeviceID  string `json:"device_id"`
	SessionID string `json:"sid"`
	Issuer    string `json:"iss,omitempty"`
	TokenID   string `json:"jti"`
	TokenType string `json:"typ"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

// Encode 签发 HMAC 自包含令牌。
func (c *Codec) Encode(_ context.Context, in tokenx.Claims) (tokenx.EncodedToken, error) {
	if in.UserID == "" || in.DeviceID == "" || in.SessionID == "" {
		return tokenx.EncodedToken{}, errors.New("hmaccodec: user ID, device ID and session ID are required")
	}
	if in.Kind != tokenx.KindAccess && in.Kind != tokenx.KindRefresh {
		return tokenx.EncodedToken{}, fmt.Errorf("hmaccodec: invalid token kind %q", in.Kind)
	}

	tokenID, err := uuid.NewV7()
	if err != nil {
		return tokenx.EncodedToken{}, fmt.Errorf("hmaccodec: generate token id: %w", err)
	}

	payload, err := json.Marshal(wireClaims{
		UserID:    in.UserID,
		DeviceID:  in.DeviceID,
		SessionID: in.SessionID,
		Issuer:    in.Issuer,
		TokenID:   tokenID.String(),
		TokenType: string(in.Kind),
		IssuedAt:  in.IssuedAt.Unix(),
		ExpiresAt: in.ExpiresAt.Unix(),
	})
	if err != nil {
		return tokenx.EncodedToken{}, fmt.Errorf("hmaccodec: marshal claims: %w", err)
	}

	encodedPayload := base64.RawURLEncoding.EncodeToString(payload)
	signingInput := version + "." + encodedPayload
	mac := hmac.New(sha256.New, c.secret)
	mac.Write([]byte(signingInput))
	encodedMAC := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return tokenx.EncodedToken{Value: signingInput + "." + encodedMAC}, nil
}

// Decode 验签并解析 HMAC 自包含令牌。
func (c *Codec) Decode(_ context.Context, raw string) (tokenx.DecodedToken, error) {
	parts := strings.Split(raw, ".")
	if len(parts) != 3 || parts[0] != version {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: hmac format", tokenx.ErrTokenInvalid)
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: hmac payload", tokenx.ErrTokenInvalid)
	}
	providedMAC, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: hmac signature", tokenx.ErrTokenInvalid)
	}

	mac := hmac.New(sha256.New, c.secret)
	mac.Write([]byte(version + "." + parts[1]))
	if subtle.ConstantTimeCompare(providedMAC, mac.Sum(nil)) != 1 {
		return tokenx.DecodedToken{}, tokenx.ErrTokenInvalid
	}

	var wire wireClaims
	if err := json.Unmarshal(payload, &wire); err != nil {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: hmac claims", tokenx.ErrTokenInvalid)
	}
	if wire.UserID == "" || wire.DeviceID == "" || wire.SessionID == "" || wire.TokenID == "" {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: missing identity claims", tokenx.ErrTokenInvalid)
	}

	kind := tokenx.TokenKind(wire.TokenType)
	if kind != tokenx.KindAccess && kind != tokenx.KindRefresh {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: invalid token kind %q", tokenx.ErrTokenInvalid, wire.TokenType)
	}
	if wire.ExpiresAt <= 0 {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: missing expiration", tokenx.ErrTokenInvalid)
	}

	expiresAt := time.Unix(wire.ExpiresAt, 0)
	if !c.clock.Now().Before(expiresAt.Add(c.leeway)) {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: hmac token expired", tokenx.ErrTokenExpired)
	}

	issuedAt := time.Time{}
	if wire.IssuedAt > 0 {
		issuedAt = time.Unix(wire.IssuedAt, 0)
	}

	return tokenx.DecodedToken{
		Kind: kind,
		Claims: &tokenx.Claims{
			UserID:    wire.UserID,
			DeviceID:  wire.DeviceID,
			SessionID: wire.SessionID,
			Issuer:    wire.Issuer,
			TokenID:   wire.TokenID,
			Kind:      kind,
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	}, nil
}
