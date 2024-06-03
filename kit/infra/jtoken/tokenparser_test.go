package jtoken

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Test_JWTInstance_GenerateJWT(t *testing.T) {
	t.Log(jwt.RegisteredClaims{
		Issuer:    "iss",
		Subject:   "sub",
		Audience:  jwt.ClaimStrings{"aud"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(3600 * time.Second)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        "id",
	})
	t.Log(jwt.MapClaims{})

	jt := NewJWTInstance([]byte("2024/3/23"))
	token, _ := jt.CreateToken(
		WithIssuer("test"),
		WithSubject("test"),
		WithAudience("test"),
		WithExpiresAt(time.Now().Unix()+3600),
		WithNotBefore(time.Now().Unix()),
		WithIssuedAt(time.Now().Unix()),
		WithId("test"),
		WithClaimExt("test", "test"),
	) // 生成有效期为24小时的 JWT
	t.Log(token)

	tk, err := jt.ParseToken(token)
	t.Log(tk)
	t.Log(err)
}
