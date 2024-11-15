package jwtx

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtAudience  = "aud"
	JwtExpire    = "exp"
	JwtId        = "jti"
	JwtIssueAt   = "iat"
	JwtIssuer    = "iss"
	JwtNotBefore = "nbf"
	JwtSubject   = "sub"
	JwtExt       = "ext"
)

type JwtInstance struct {
	secret []byte
}

func NewJWTInstance(secret []byte) *JwtInstance {
	return &JwtInstance{secret}
}

// GenerateJWT 生成JWT
func (that JwtInstance) CreateToken(options ...Option) (string, error) {
	claims := jwt.MapClaims{}

	for _, option := range options {
		option(claims)
	}

	// 创建一个新的JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 设置签名并获取token字符串
	token, err := jwtToken.SignedString(that.secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析JWT
func (that JwtInstance) ParseToken(tokenString string) (*jwt.Token, error) {
	// 解析JWT字符串
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return that.secret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
	// 验证token
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	return claims, nil
	//}
	//
	//return nil, fmt.Errorf("invalid token")
}
