package response

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

type Login struct {
	Token    string      `json:"token"`
	UserInfo *UserDetail `json:"userInfo"`
}

type OauthLoginUrl struct {
	Url string `json:"url" example:""` // 授权地址
}

type UserDetail struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Intro         string `json:"intro"`
	Email         string `json:"email"`
	LoginType     string `json:"login_type"`
	IpAddress     string `json:"ip_address"`
	IpSource      string `json:"ip_source"`
	LastLoginTime string `json:"last_login_time"`

	Roles []*entity.Role
}

type LoginHistory struct {
	LoginType string // 登录类型
	IpAddress string // ip host
	IpSource  string // ip 源
	LoginTime string // 创建时间
}
