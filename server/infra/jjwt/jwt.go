package jjwt

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtToken struct {
	SigningKey  []byte
	TokenPrefix string
	Issuer      string
}

type JwtClaims struct {
	Uid       int    `json:"uid"`
	Username  string `json:"username"`
	LoginType string `json:"login_type"`
	//Roles    []string `json:"roles"`
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
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				// ValidationErrorMalformed是一个uint常量，表示token不可用
				return nil, TokenMalformed
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				// ValidationErrorExpired表示Token过期
				return nil, TokenExpired
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				// ValidationErrorNotValidYet表示无效token
				return nil, TokenNotValidYet
			default:
				return nil, TokenInvalid
			}
		}
		return nil, err
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
func (j *JwtToken) CreateClaims(userId int, username string, loginType string) (string, error) {
	claims := JwtClaims{
		Uid:       userId,
		Username:  username,
		LoginType: loginType,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    j.Issuer,
		},
	}

	return j.createToken(claims)
}

func (j *JwtToken) VerifyToken(token string, uid string) (*JwtClaims, error) {
	if token == "" {
		return nil, errors.New("token is null")
	}
	if uid == "" {
		return nil, errors.New("uid is null")
	}
	//验证token是否 Bearer 开头的
	ok := strings.HasPrefix(token, j.TokenPrefix)
	if !ok {
		return nil, errors.New("token must be has prefix :" + j.TokenPrefix)
	}

	token = strings.TrimPrefix(token, j.TokenPrefix)
	// 解析token
	claims, err := j.parserToken(token)
	if err != nil {
		return nil, err
	}

	if uid != strconv.Itoa(claims.Uid) {
		return nil, errors.New("uid is not equal")
	}
	return claims, nil
}
