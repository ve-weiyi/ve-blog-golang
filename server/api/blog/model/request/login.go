package request

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/fmtplus"
)

type User struct {
	Username  string `json:"username" from:"username" example:"admin@qq.com"`
	Password  string `json:"password" from:"password" example:"123456"`
	Code      string `json:"code" from:"code" example:""`
	LoginType string `json:"-" from:"-" example:""` // 登录类型
}

func (m User) IsValid() error {
	if m.Username == "" || m.Password == "" {
		return codes.NewError(codes.CodeInvalidParameter, "用户名和密码不能为null")
	}

	//验证邮箱格式是否正确
	if !fmtplus.IsEmailValid(m.Username) {
		return codes.NewError(codes.CodeInvalidParameter, "邮箱格式不正确")
	}

	if len(m.Password) < 6 {
		return codes.NewError(codes.CodeInvalidParameter, "密码长度不能小于6")
	}

	return nil
}

// 用户名只能是邮箱
type UserEmail struct {
	Username string `json:"username" from:"username" example:"admin@qq.com"`
}

func (m UserEmail) IsValid() error {
	//验证邮箱格式是否正确
	if !fmtplus.IsEmailValid(m.Username) {
		return codes.NewError(codes.CodeInvalidParameter, "邮箱格式不正确")
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
	ID          int    `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 旧密码
	NewPassword string `json:"newPassword"` // 新密码
}

type OauthLoginReq struct {
	Platform string `json:"platform" example:""` // 平台
	Code     string `json:"code" example:""`     // 授权码
	State    string `json:"state" example:""`    // 状态
}
