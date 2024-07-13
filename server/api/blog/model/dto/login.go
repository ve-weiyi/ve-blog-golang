package dto

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
)

type LoginReq struct {
	Username string `json:"username" from:"username" example:"admin@qq.com"`
	Password string `json:"password" from:"password" example:"123456"`
	Code     string `json:"code" from:"code" example:""`
}

func (m LoginReq) IsValid() error {
	if m.Username == "" || m.Password == "" {
		return fmt.Errorf("用户名或密码不能为空")
	}

	// 验证邮箱格式是否正确
	if !valid.IsEmailValid(m.Username) {
		return fmt.Errorf("邮箱格式不正确")
	}

	if len(m.Password) < 6 {
		return fmt.Errorf("密码长度不能小于6位")
	}

	return nil
}

// 用户名只能是邮箱
type UserEmailReq struct {
	Username string `json:"username" from:"username" example:"admin@qq.com"`
}

func (m UserEmailReq) IsValid() error {
	// 验证邮箱格式是否正确
	if !valid.IsEmailValid(m.Username) {
		return fmt.Errorf("邮箱格式不正确")
	}

	return nil
}

type ResetPasswordReq struct {
	Username string `json:"username" from:"username" example:"admin@qq.com"`
	Password string `json:"password" from:"password" example:"123456"`
	Code     string `json:"code" from:"code" example:""`
}

// Modify password structure
type ChangePasswordReq struct {
	Id          int64  `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 旧密码
	NewPassword string `json:"newPassword"` // 新密码
}

type OauthLoginReq struct {
	Platform string `json:"platform" example:""` // 平台
	Code     string `json:"code" example:""`     // 授权码
	State    string `json:"state" example:""`    // 状态
}

// 用户登录信息
type LoginResp struct {
	Token        *Token        `json:"token"`
	UserInfo     *UserInfo     `json:"user_info"`
	LoginHistory *LoginHistory `json:"login_history"`
}

type Token struct {
	UserId           int64  `json:"user_id"`            // 用户id
	TokenType        string `json:"token_type"`         // token类型,Bearer
	AccessToken      string `json:"access_token"`       // 访问token,过期时间较短。2h
	ExpiresIn        int64  `json:"expires_in"`         // 访问token过期时间
	RefreshToken     string `json:"refresh_token"`      // 刷新token,过期时间较长。30d
	RefreshExpiresIn int64  `json:"refresh_expires_in"` // 刷新token过期时间
	Scope            string `json:"scope"`              // 作用域
}

type UserInfo struct {
	UserId   int64  `json:"user_id"`  // 用户id
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
	Intro    string `json:"intro"`    // 个人简介
	Website  string `json:"website"`  // 个人网站
	Email    string `json:"email"`    // 邮箱

	// ArticleLikeSet []string `json:"article_like_set"` // 文章点赞集合
	// CommentLikeSet []string `json:"comment_like_set"` // 评论点赞集合
	// TalkLikeSet    []string `json:"talk_like_set"`    // 说说点赞集合

	Roles []*RoleDTO `json:"roles"` // 角色列表
}

type LoginHistory struct {
	Id        int64  `json:"id"`
	LoginType string `json:"login_type"` // 登录类型
	Agent     string `json:"agent"`      // 代理
	IpAddress string `json:"ip_address"` // ip host
	IpSource  string `json:"ip_source"`  // ip 源
	LoginTime string `json:"login_time"` // 创建时间
}

type OauthLoginUrl struct {
	Url string `json:"url" example:""` // 授权地址
}
