package auth

import (
	"errors"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
)

// toToken 将令牌对转换为前端契约结构。
func toToken(tk *tokenx.TokenPair) *types.Token {
	return &types.Token{
		TokenType:        headerx.TokenTypeBearer,
		AccessToken:      tk.AccessToken,
		ExpiresIn:        tk.AccessExpiresIn,
		RefreshToken:     tk.RefreshToken,
		RefreshExpiresIn: tk.RefreshExpiresIn,
		RefreshExpiresAt: tk.RefreshExpiresAt.Unix(),
	}
}

// toAuthError 把承载凭据的失败统一映射为登录态错误，避免落成 500。
func toAuthError(err error) error {
	switch {
	case errors.Is(err, tokenx.ErrTokenExpired), errors.Is(err, tokenx.ErrRefreshReused):
		return bizerr.NewBizError(bizcode.CodeLoginExpired, "登录已过期，请重新登录")
	case errors.Is(err, tokenx.ErrTokenEmpty), errors.Is(err, tokenx.ErrTokenInvalid),
		errors.Is(err, tokenx.ErrSessionNotFound), errors.Is(err, tokenx.ErrSessionExpired),
		errors.Is(err, tokenx.ErrCodecMismatch):
		return bizerr.NewBizError(bizcode.CodeUnauthenticated, "登录状态已失效，请重新登录")
	default:
		return err
	}
}
