package response

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

type Login struct {
	Token            string        `json:"token"`
	Userinfo         *UserDetail   `json:"userinfo"`
	LastLoginHistory *LoginHistory `json:"last_login_history"`
}

type OauthLoginUrl struct {
	Url string `json:"url" example:""` // 授权地址
}

type UserDetail struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
	Email    string `json:"email"`

	Roles []*entity.Role
}

type LoginHistory struct {
	LoginType string `json:"login_type"` // 登录类型
	IpAddress string `json:"ip_address"` // ip host
	IpSource  string `json:"ip_source"`  // ip 源
	LoginTime string `json:"login_time"` // 创建时间
}
