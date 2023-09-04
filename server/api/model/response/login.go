package response

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

// 用户登录信息
type Login struct {
	*Token
	UserInfo  *UserInfo  `json:"user_info"`
	LoginInfo *LoginInfo `json:"login_info"`
}

type Token struct {
	TokenType        string `json:"token_type"`         // token类型,Bearer
	AccessToken      string `json:"access_token"`       // 访问token,过期时间较短。2h
	ExpiresIn        int64  `json:"expires_in"`         // 访问token过期时间
	RefreshToken     string `json:"refresh_token"`      // 刷新token,过期时间较长。30d
	RefreshExpiresIn int64  `json:"refresh_expires_in"` // 刷新token过期时间
	UID              int    `json:"uid"`                // 用户id
}

type UserInfo struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Status    int       `json:"status"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Intro     string    `json:"intro"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`

	Roles []*entity.Role `json:"roles"`
}

type LoginInfo struct {
	LoginType string `json:"login_type"` // 登录类型
	IpAddress string `json:"ip_address"` // ip host
	IpSource  string `json:"ip_source"`  // ip 源
	LoginTime string `json:"login_time"` // 创建时间
}

type OauthLoginUrl struct {
	Url string `json:"url" example:""` // 授权地址
}
