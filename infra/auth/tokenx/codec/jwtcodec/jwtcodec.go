// Package jwtcodec 实现 HS256 JWT codec。
//
// JWT 只负责格式与签名；会话存在性、撤销、轮换仍由 tokenx.Service 通过 SessionStore 判定。
package jwtcodec

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
)

const (
	minSecret = 32
	name      = "jwt-hs256"
)

// Codec 是 HS256 JWT 编码器。
type Codec struct {
	secret []byte
	clock  tokenx.Clock
	leeway time.Duration
}

// Option 是 JWT Codec 的可选配置。
type Option func(*Codec)

// WithClock 注入时间源，供测试使用。
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

// New 创建 HS256 JWT Codec。
func New(secret string, opts ...Option) (tokenx.Codec, error) {
	if len(secret) < minSecret {
		return nil, fmt.Errorf("jwtcodec: secret must be at least %d bytes", minSecret)
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
func (c *Codec) Name() string {
	return name
}

type jwtClaims struct {
	jwt.RegisteredClaims
	SessionID string `json:"sid"`
	DeviceID  string `json:"did"`
}

// Encode 签发 JWT。
func (c *Codec) Encode(_ context.Context, in tokenx.Claims) (tokenx.EncodedToken, error) {
	if in.UserID == "" || in.DeviceID == "" || in.SessionID == "" {
		return tokenx.EncodedToken{}, errors.New("jwtcodec: user ID, device ID and session ID are required")
	}
	if in.Kind != tokenx.KindAccess && in.Kind != tokenx.KindRefresh {
		return tokenx.EncodedToken{}, fmt.Errorf("jwtcodec: invalid token kind %q", in.Kind)
	}

	tokenID, err := uuid.NewV7()
	if err != nil {
		return tokenx.EncodedToken{}, fmt.Errorf("jwtcodec: generate token id: %w", err)
	}

	registered := jwt.RegisteredClaims{
		Subject:   in.UserID,
		ID:        tokenID.String(),
		Issuer:    in.Issuer,
		IssuedAt:  jwt.NewNumericDate(in.IssuedAt),
		ExpiresAt: jwt.NewNumericDate(in.ExpiresAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		RegisteredClaims: registered,
		SessionID:        in.SessionID,
		DeviceID:         in.DeviceID,
	})
	if in.Kind == tokenx.KindAccess {
		token.Header["typ"] = "at+jwt"
	} else {
		token.Header["typ"] = "JWT"
	}

	value, err := token.SignedString(c.secret)
	if err != nil {
		return tokenx.EncodedToken{}, fmt.Errorf("jwtcodec: sign token: %w", err)
	}
	return tokenx.EncodedToken{Value: value}, nil
}

// Decode 验签并解析 JWT。
func (c *Codec) Decode(_ context.Context, raw string) (tokenx.DecodedToken, error) {
	parsed := &jwtClaims{}
	token, err := jwt.ParseWithClaims(raw, parsed, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("jwtcodec: unexpected signing method %q", t.Method.Alg())
		}
		return c.secret, nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(c.leeway),
		jwt.WithTimeFunc(c.clock.Now),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return tokenx.DecodedToken{}, fmt.Errorf("%w: %v", tokenx.ErrTokenExpired, err)
		}
		return tokenx.DecodedToken{}, fmt.Errorf("%w: %v", tokenx.ErrTokenInvalid, err)
	}
	if !token.Valid || parsed.ExpiresAt == nil {
		return tokenx.DecodedToken{}, tokenx.ErrTokenInvalid
	}

	var kind tokenx.TokenKind
	switch token.Header["typ"] {
	case "at+jwt":
		kind = tokenx.KindAccess
	case "JWT":
		kind = tokenx.KindRefresh
	default:
		return tokenx.DecodedToken{}, fmt.Errorf("%w: invalid typ %v", tokenx.ErrTokenInvalid, token.Header["typ"])
	}
	if parsed.Subject == "" || parsed.DeviceID == "" || parsed.SessionID == "" {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: missing identity claims", tokenx.ErrTokenInvalid)
	}

	issuedAt := time.Time{}
	if parsed.IssuedAt != nil {
		issuedAt = parsed.IssuedAt.Time
	}

	claims := &tokenx.Claims{
		UserID:    parsed.Subject,
		DeviceID:  parsed.DeviceID,
		SessionID: parsed.SessionID,
		Issuer:    parsed.Issuer,
		TokenID:   parsed.ID,
		Kind:      kind,
		IssuedAt:  issuedAt,
		ExpiresAt: parsed.ExpiresAt.Time,
	}
	return tokenx.DecodedToken{Kind: kind, Claims: claims}, nil
}
