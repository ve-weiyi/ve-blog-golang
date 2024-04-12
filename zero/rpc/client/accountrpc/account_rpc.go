// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package accountrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Api                  = account.Api
	ApiDetailsDTO        = account.ApiDetailsDTO
	ApiPageResp          = account.ApiPageResp
	BatchResult          = account.BatchResult
	EmptyReq             = account.EmptyReq
	EmptyResp            = account.EmptyResp
	IdReq                = account.IdReq
	IdsReq               = account.IdsReq
	LoginHistory         = account.LoginHistory
	LoginHistoryPageResp = account.LoginHistoryPageResp
	LoginReq             = account.LoginReq
	LoginResp            = account.LoginResp
	Menu                 = account.Menu
	MenuDetailsDTO       = account.MenuDetailsDTO
	MenuPageResp         = account.MenuPageResp
	OauthLoginReq        = account.OauthLoginReq
	OauthLoginUrlResp    = account.OauthLoginUrlResp
	PageCondition        = account.PageCondition
	PageLimit            = account.PageLimit
	PageQuery            = account.PageQuery
	PageResult           = account.PageResult
	PageSort             = account.PageSort
	ResetPasswordReq     = account.ResetPasswordReq
	Role                 = account.Role
	RoleDTO              = account.RoleDTO
	RoleDetailsDTO       = account.RoleDetailsDTO
	RolePageResp         = account.RolePageResp
	RoleResourcesResp    = account.RoleResourcesResp
	SyncMenuRequest      = account.SyncMenuRequest
	UpdateRoleApisReq    = account.UpdateRoleApisReq
	UpdateRoleMenusReq   = account.UpdateRoleMenusReq
	UserEmailReq         = account.UserEmailReq
	UserInfoReq          = account.UserInfoReq
	UserInfoResp         = account.UserInfoResp

	AccountRpc interface {
		// 登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 登出
		Logout(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 注销
		Logoff(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 注册
		Register(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 发送注册邮件
		RegisterEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 发送忘记密码邮件
		ForgetPasswordEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 重置密码
		ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 第三方登录
		OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 获取授权地址
		GetOauthAuthorizeUrl(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthLoginUrlResp, error)
	}

	defaultAccountRpc struct {
		cli zrpc.Client
	}
)

func NewAccountRpc(cli zrpc.Client) AccountRpc {
	return &defaultAccountRpc{
		cli: cli,
	}
}

// 登录
func (m *defaultAccountRpc) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 登出
func (m *defaultAccountRpc) Logout(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}

// 注销
func (m *defaultAccountRpc) Logoff(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.Logoff(ctx, in, opts...)
}

// 注册
func (m *defaultAccountRpc) Register(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 发送注册邮件
func (m *defaultAccountRpc) RegisterEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.RegisterEmail(ctx, in, opts...)
}

// 发送忘记密码邮件
func (m *defaultAccountRpc) ForgetPasswordEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.ForgetPasswordEmail(ctx, in, opts...)
}

// 重置密码
func (m *defaultAccountRpc) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.ResetPassword(ctx, in, opts...)
}

// 第三方登录
func (m *defaultAccountRpc) OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.OauthLogin(ctx, in, opts...)
}

// 获取授权地址
func (m *defaultAccountRpc) GetOauthAuthorizeUrl(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthLoginUrlResp, error) {
	client := account.NewAccountRpcClient(m.cli.Conn())
	return client.GetOauthAuthorizeUrl(ctx, in, opts...)
}
