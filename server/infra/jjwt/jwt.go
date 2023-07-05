package jjwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JwtToken struct {
	SigningKey  []byte
	TokenPrefix string
	Issuer      string
}

type JwtClaims struct {
	Uid      int      `json:"uid"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	//UserClaims         interface{} //用户信息
	jwt.StandardClaims //标准荷载
}

var (
	TokenExpired     = errors.New("token 已过期")
	TokenNotValidYet = errors.New("token 无效")
	TokenMalformed   = errors.New("token 格式错误")
	TokenInvalid     = errors.New("token 不可用")
)

// createToken 生成token
func (j *JwtToken) createToken(claims JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// parserToken 解析token
func (j *JwtToken) parserToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// 根据用户登录信息生成token，
func (j *JwtToken) CreateClaims(userId int, username string, roles []string) (string, error) {
	claims := JwtClaims{
		//UserClaims: info,
		Uid:      userId,
		Username: username,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    j.Issuer,
		},
	}

	return j.createToken(claims)
}

func (j *JwtToken) ParseTokenByGin(c *gin.Context) (*JwtClaims, error) {
	tokenHeader := c.Request.Header.Get("Authorization")

	//token是空
	if tokenHeader == "" {
		//tokenHeader, _ = c.Cookie("token")
		//global.LOG.Info("get token by cookie :" + tokenHeader)
		return nil, errors.New("token is null")
	}

	//验证token是否 Bearer 开头的
	ok := strings.HasPrefix(tokenHeader, j.TokenPrefix)
	if !ok {
		return nil, errors.New("token must be has prefix :" + j.TokenPrefix)
	}

	token := strings.TrimPrefix(tokenHeader, j.TokenPrefix)

	// 解析token
	claims, err := j.parserToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
