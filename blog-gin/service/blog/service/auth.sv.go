package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AuthService struct {
	svcCtx *svctx.ServiceContext
}

func NewAuthService(svcCtx *svctx.ServiceContext) *AuthService {
	return &AuthService{
		svcCtx: svcCtx,
	}
}

// 获取游客身份信息
func (s *AuthService) GetTouristInfo(reqCtx *request.Context, in *dto.EmptyReq) (out *dto.GetTouristInfoResp, err error) {
	// todo

	return
}

// 邮箱登录
func (s *AuthService) EmailLogin(reqCtx *request.Context, in *dto.EmailLoginReq) (out *dto.LoginResp, err error) {
	// todo

	return
}

// 获取验证码
func (s *AuthService) GetCaptchaCode(reqCtx *request.Context, in *dto.GetCaptchaCodeReq) (out *dto.GetCaptchaCodeResp, err error) {
	// todo

	return
}

// 第三方登录授权地址
func (s *AuthService) GetOauthAuthorizeUrl(reqCtx *request.Context, in *dto.GetOauthAuthorizeUrlReq) (out *dto.GetOauthAuthorizeUrlResp, err error) {
	// todo

	return
}

// 登录
func (s *AuthService) Login(reqCtx *request.Context, in *dto.LoginReq) (out *dto.LoginResp, err error) {
	// todo

	return
}

// 手机登录
func (s *AuthService) PhoneLogin(reqCtx *request.Context, in *dto.PhoneLoginReq) (out *dto.LoginResp, err error) {
	// todo

	return
}

// 注册
func (s *AuthService) Register(reqCtx *request.Context, in *dto.RegisterReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 重置密码
func (s *AuthService) ResetPassword(reqCtx *request.Context, in *dto.ResetPasswordReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 发送邮件验证码
func (s *AuthService) SendEmailVerifyCode(reqCtx *request.Context, in *dto.SendEmailVerifyCodeReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 发送手机验证码
func (s *AuthService) SendPhoneVerifyCode(reqCtx *request.Context, in *dto.SendPhoneVerifyCodeReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 第三方登录
func (s *AuthService) ThirdLogin(reqCtx *request.Context, in *dto.ThirdLoginReq) (out *dto.LoginResp, err error) {
	// todo

	return
}

// 注销
func (s *AuthService) Logoff(reqCtx *request.Context, in *dto.EmptyReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 登出
func (s *AuthService) Logout(reqCtx *request.Context, in *dto.EmptyReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}
