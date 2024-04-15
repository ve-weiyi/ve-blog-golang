// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package apirpc

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
	PageUserInfoResp     = account.PageUserInfoResp
	ResetPasswordReq     = account.ResetPasswordReq
	Role                 = account.Role
	RoleDetailsDTO       = account.RoleDetailsDTO
	RoleLabelDTO         = account.RoleLabelDTO
	RolePageResp         = account.RolePageResp
	RoleResourcesResp    = account.RoleResourcesResp
	SyncMenuRequest      = account.SyncMenuRequest
	UpdateRoleApisReq    = account.UpdateRoleApisReq
	UpdateRoleMenusReq   = account.UpdateRoleMenusReq
	UpdateUserAvatarReq  = account.UpdateUserAvatarReq
	UpdateUserInfoReq    = account.UpdateUserInfoReq
	UpdateUserRoleReq    = account.UpdateUserRoleReq
	UpdateUserStatusReq  = account.UpdateUserStatusReq
	UserDTO              = account.UserDTO
	UserEmailReq         = account.UserEmailReq
	UserInfoResp         = account.UserInfoResp

	ApiRpc interface {
		// 创建接口
		CreateApi(ctx context.Context, in *Api, opts ...grpc.CallOption) (*Api, error)
		// 更新接口
		UpdateApi(ctx context.Context, in *Api, opts ...grpc.CallOption) (*Api, error)
		// 删除接口
		DeleteApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResult, error)
		// 批量删除接口
		DeleteApiList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResult, error)
		// 查询接口
		FindApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Api, error)
		// 分页获取接口列表
		FindApiList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ApiPageResp, error)
		// 同步接口列表
		SyncApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResult, error)
		// 清空接口列表
		CleanApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
	}

	defaultApiRpc struct {
		cli zrpc.Client
	}
)

func NewApiRpc(cli zrpc.Client) ApiRpc {
	return &defaultApiRpc{
		cli: cli,
	}
}

// 创建接口
func (m *defaultApiRpc) CreateApi(ctx context.Context, in *Api, opts ...grpc.CallOption) (*Api, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.CreateApi(ctx, in, opts...)
}

// 更新接口
func (m *defaultApiRpc) UpdateApi(ctx context.Context, in *Api, opts ...grpc.CallOption) (*Api, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.UpdateApi(ctx, in, opts...)
}

// 删除接口
func (m *defaultApiRpc) DeleteApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResult, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

// 批量删除接口
func (m *defaultApiRpc) DeleteApiList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResult, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.DeleteApiList(ctx, in, opts...)
}

// 查询接口
func (m *defaultApiRpc) FindApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Api, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.FindApi(ctx, in, opts...)
}

// 分页获取接口列表
func (m *defaultApiRpc) FindApiList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ApiPageResp, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.FindApiList(ctx, in, opts...)
}

// 同步接口列表
func (m *defaultApiRpc) SyncApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResult, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.SyncApiList(ctx, in, opts...)
}

// 清空接口列表
func (m *defaultApiRpc) CleanApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewApiRpcClient(m.cli.Conn())
	return client.CleanApiList(ctx, in, opts...)
}
