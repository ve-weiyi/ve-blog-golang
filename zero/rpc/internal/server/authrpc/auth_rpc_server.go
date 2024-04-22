// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/logic/authrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

type AuthRpcServer struct {
	svcCtx *svc.ServiceContext
	account.UnimplementedAuthRpcServer
}

func NewAuthRpcServer(svcCtx *svc.ServiceContext) *AuthRpcServer {
	return &AuthRpcServer{
		svcCtx: svcCtx,
	}
}

// 登录
func (s *AuthRpcServer) Login(ctx context.Context, in *account.LoginReq) (*account.LoginResp, error) {
	l := authrpclogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 登出
func (s *AuthRpcServer) Logout(ctx context.Context, in *account.EmptyReq) (*account.EmptyResp, error) {
	l := authrpclogic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

// 注销
func (s *AuthRpcServer) Logoff(ctx context.Context, in *account.EmptyReq) (*account.EmptyResp, error) {
	l := authrpclogic.NewLogoffLogic(ctx, s.svcCtx)
	return l.Logoff(in)
}

// 注册
func (s *AuthRpcServer) Register(ctx context.Context, in *account.LoginReq) (*account.EmptyResp, error) {
	l := authrpclogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 发送注册邮件
func (s *AuthRpcServer) RegisterEmail(ctx context.Context, in *account.UserEmailReq) (*account.EmptyResp, error) {
	l := authrpclogic.NewRegisterEmailLogic(ctx, s.svcCtx)
	return l.RegisterEmail(in)
}

// 发送忘记密码邮件
func (s *AuthRpcServer) ForgetPasswordEmail(ctx context.Context, in *account.UserEmailReq) (*account.EmptyResp, error) {
	l := authrpclogic.NewForgetPasswordEmailLogic(ctx, s.svcCtx)
	return l.ForgetPasswordEmail(in)
}

// 重置密码
func (s *AuthRpcServer) ResetPassword(ctx context.Context, in *account.ResetPasswordReq) (*account.EmptyResp, error) {
	l := authrpclogic.NewResetPasswordLogic(ctx, s.svcCtx)
	return l.ResetPassword(in)
}

// 第三方登录
func (s *AuthRpcServer) OauthLogin(ctx context.Context, in *account.OauthLoginReq) (*account.LoginResp, error) {
	l := authrpclogic.NewOauthLoginLogic(ctx, s.svcCtx)
	return l.OauthLogin(in)
}

// 获取授权地址
func (s *AuthRpcServer) GetOauthAuthorizeUrl(ctx context.Context, in *account.OauthLoginReq) (*account.OauthLoginUrlResp, error) {
	l := authrpclogic.NewGetOauthAuthorizeUrlLogic(ctx, s.svcCtx)
	return l.GetOauthAuthorizeUrl(in)
}
