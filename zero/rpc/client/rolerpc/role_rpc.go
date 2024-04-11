// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package rolerpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Api                = account.Api
	ApiDetailsDTO      = account.ApiDetailsDTO
	ApiPageResp        = account.ApiPageResp
	BatchResult        = account.BatchResult
	EmptyReq           = account.EmptyReq
	EmptyResp          = account.EmptyResp
	IdReq              = account.IdReq
	IdsReq             = account.IdsReq
	LoginReq           = account.LoginReq
	LoginResp          = account.LoginResp
	Menu               = account.Menu
	MenuDetailsDTO     = account.MenuDetailsDTO
	MenuPageResp       = account.MenuPageResp
	OauthLoginReq      = account.OauthLoginReq
	OauthLoginUrlResp  = account.OauthLoginUrlResp
	PageCondition      = account.PageCondition
	PageLimit          = account.PageLimit
	PageQuery          = account.PageQuery
	PageResult         = account.PageResult
	PageSort           = account.PageSort
	ResetPasswordReq   = account.ResetPasswordReq
	Role               = account.Role
	RoleDTO            = account.RoleDTO
	RoleDetailsDTO     = account.RoleDetailsDTO
	RolePageResp       = account.RolePageResp
	RoleResourceResp   = account.RoleResourceResp
	SyncMenuRequest    = account.SyncMenuRequest
	UpdateRoleApisReq  = account.UpdateRoleApisReq
	UpdateRoleMenusReq = account.UpdateRoleMenusReq
	UserEmailReq       = account.UserEmailReq

	RoleRpc interface {
		// 创建角色
		CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
		// 更新角色
		UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
		// 删除角色
		DeleteRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResult, error)
		// 批量删除角色
		DeleteRoleList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResult, error)
		// 查询角色
		FindRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Role, error)
		// 分页获取角色列表
		FindRoleList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*RolePageResp, error)
		// 查询角色
		FindRoleResource(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleResourceResp, error)
		// 更新角色菜单
		UpdateRoleMenus(ctx context.Context, in *UpdateRoleMenusReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 更新角色资源
		UpdateRoleApis(ctx context.Context, in *UpdateRoleApisReq, opts ...grpc.CallOption) (*EmptyResp, error)
	}

	defaultRoleRpc struct {
		cli zrpc.Client
	}
)

func NewRoleRpc(cli zrpc.Client) RoleRpc {
	return &defaultRoleRpc{
		cli: cli,
	}
}

// 创建角色
func (m *defaultRoleRpc) CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

// 更新角色
func (m *defaultRoleRpc) UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

// 删除角色
func (m *defaultRoleRpc) DeleteRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResult, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

// 批量删除角色
func (m *defaultRoleRpc) DeleteRoleList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResult, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.DeleteRoleList(ctx, in, opts...)
}

// 查询角色
func (m *defaultRoleRpc) FindRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Role, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.FindRole(ctx, in, opts...)
}

// 分页获取角色列表
func (m *defaultRoleRpc) FindRoleList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*RolePageResp, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.FindRoleList(ctx, in, opts...)
}

// 查询角色
func (m *defaultRoleRpc) FindRoleResource(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleResourceResp, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.FindRoleResource(ctx, in, opts...)
}

// 更新角色菜单
func (m *defaultRoleRpc) UpdateRoleMenus(ctx context.Context, in *UpdateRoleMenusReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.UpdateRoleMenus(ctx, in, opts...)
}

// 更新角色资源
func (m *defaultRoleRpc) UpdateRoleApis(ctx context.Context, in *UpdateRoleApisReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := account.NewRoleRpcClient(m.cli.Conn())
	return client.UpdateRoleApis(ctx, in, opts...)
}
