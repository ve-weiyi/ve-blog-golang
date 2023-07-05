package response

import "github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"

type Login struct {
	Token    string     `json:"token"`
	UserInfo UserDetail `json:"userInfo"`
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
	Status        int    `json:"status"`
	LoginType     int    `json:"login_type"`
	IpAddress     string `json:"ip_address"`
	IpSource      string `json:"ip_source"`
	CreatedAt     string `json:"created_at"`
	LastLoginTime string `json:"last_login_time"`

	Roles []*entity.Role
}

type LoginHistory struct {
	LoginType int    // 登录类型
	IpAddress string // ip host
	IpSource  string // ip 源
	LoginTime string // 创建时间
}

type UserInfo struct {
	ID            int    `json:"id"`
	UUID          string `json:"user_id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Intro         string `json:"intro"`
	Email         string `json:"email"`
	LoginType     int    `json:"loginType"`
	IpAddress     string `json:"ip_address"`
	IpSource      string `json:"ip_source"`
	LastLoginTime string `json:"last_login_time"`

	Roles []*entity.Role
}
