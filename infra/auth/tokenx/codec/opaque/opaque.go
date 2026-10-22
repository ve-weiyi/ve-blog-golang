// Package opaque 实现不透明 reference token。
//
// token 只携带 kind 前缀与随机 secret，不携带 uid/sid/claims；
// 服务端通过 SHA-256(kind:token) 查找会话。
package opaque

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
)

const (
	accessPrefix  = "at_"
	refreshPrefix = "rt_"
	secretBytes   = 32
	name          = "opaque-v3"
)

// Codec 是不透明令牌编码器。
type Codec struct{}

// New 创建不透明令牌 Codec。
func New() tokenx.Codec {
	return Codec{}
}

// Name 返回方案标识。
func (Codec) Name() string {
	return name
}

// Encode 生成 at_/rt_ + 43 字符随机串。
func (Codec) Encode(_ context.Context, claims tokenx.Claims) (tokenx.EncodedToken, error) {
	prefix, err := prefixOf(claims.Kind)
	if err != nil {
		return tokenx.EncodedToken{}, err
	}

	secret := make([]byte, secretBytes)
	if _, err := rand.Read(secret); err != nil {
		return tokenx.EncodedToken{}, fmt.Errorf("opaque: generate secret: %w", err)
	}
	value := prefix + base64.RawURLEncoding.EncodeToString(secret)

	return tokenx.EncodedToken{
		Value:      value,
		LookupHash: tokenx.HashToken(string(claims.Kind) + ":" + value),
	}, nil
}

// Decode 校验前缀与格式，返回查找哈希；不验证 token 是否真实存在。
func (Codec) Decode(_ context.Context, raw string) (tokenx.DecodedToken, error) {
	kind, err := kindOf(raw)
	if err != nil {
		return tokenx.DecodedToken{}, err
	}
	if len(raw) != len(prefixOfKind(kind))+43 {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: opaque length", tokenx.ErrTokenInvalid)
	}
	secret, err := base64.RawURLEncoding.DecodeString(raw[len(prefixOfKind(kind)):])
	if err != nil || len(secret) != secretBytes {
		return tokenx.DecodedToken{}, fmt.Errorf("%w: opaque secret", tokenx.ErrTokenInvalid)
	}

	return tokenx.DecodedToken{
		Kind:       kind,
		LookupHash: tokenx.HashToken(string(kind) + ":" + raw),
	}, nil
}

func prefixOf(kind tokenx.TokenKind) (string, error) {
	switch kind {
	case tokenx.KindAccess:
		return accessPrefix, nil
	case tokenx.KindRefresh:
		return refreshPrefix, nil
	default:
		return "", fmt.Errorf("%w: opaque kind %q", tokenx.ErrTokenInvalid, kind)
	}
}

func prefixOfKind(kind tokenx.TokenKind) string {
	if kind == tokenx.KindRefresh {
		return refreshPrefix
	}
	return accessPrefix
}

func kindOf(raw string) (tokenx.TokenKind, error) {
	switch {
	case strings.HasPrefix(raw, accessPrefix):
		return tokenx.KindAccess, nil
	case strings.HasPrefix(raw, refreshPrefix):
		return tokenx.KindRefresh, nil
	default:
		return "", fmt.Errorf("%w: opaque prefix", tokenx.ErrTokenInvalid)
	}
}
