package tokenx

import "context"

// Codec 负责令牌的编码与解码，不负责会话持久化与登录策略。
type Codec interface {
	Name() string
	Encode(ctx context.Context, claims Claims) (EncodedToken, error)
	Decode(ctx context.Context, raw string) (DecodedToken, error)
}

// EncodedToken 是 Encode 的产物。
type EncodedToken struct {
	Value string

	// LookupHash 仅不透明令牌使用：SHA-256(kind:value)。
	// 自含令牌为空，校验时用 claims 中的 sid 查会话。
	LookupHash string
}

// DecodedToken 是 Decode 的产物。
type DecodedToken struct {
	Kind       TokenKind
	LookupHash string  // 不透明令牌
	Claims     *Claims // 自含令牌；不透明令牌为 nil
}
