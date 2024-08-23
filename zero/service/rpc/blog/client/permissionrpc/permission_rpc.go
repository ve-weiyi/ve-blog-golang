// Code generated by goctl. DO NOT EDIT.
// Source: permission.proto

package permissionrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiDetails         = permissionrpc.ApiDetails
	ApiNew             = permissionrpc.ApiNew
	BatchResp          = permissionrpc.BatchResp
	CountResp          = permissionrpc.CountResp
	EmptyReq           = permissionrpc.EmptyReq
	EmptyResp          = permissionrpc.EmptyResp
	FindApiListReq     = permissionrpc.FindApiListReq
	FindApiListResp    = permissionrpc.FindApiListResp
	FindMenuListReq    = permissionrpc.FindMenuListReq
	FindMenuListResp   = permissionrpc.FindMenuListResp
	FindRoleListReq    = permissionrpc.FindRoleListReq
	FindRoleListResp   = permissionrpc.FindRoleListResp
	IdReq              = permissionrpc.IdReq
	IdsReq             = permissionrpc.IdsReq
	MenuDetails        = permissionrpc.MenuDetails
	MenuNew            = permissionrpc.MenuNew
	RoleDetails        = permissionrpc.RoleDetails
	RoleNew            = permissionrpc.RoleNew
	RoleResourcesResp  = permissionrpc.RoleResourcesResp
	SyncMenuReq        = permissionrpc.SyncMenuReq
	UpdateRoleApisReq  = permissionrpc.UpdateRoleApisReq
	UpdateRoleMenusReq = permissionrpc.UpdateRoleMenusReq
	UpdateUserRoleReq  = permissionrpc.UpdateUserRoleReq
	UserIdReq          = permissionrpc.UserIdReq

	PermissionRpc interface {
		// 创建接口
		AddApi(ctx context.Context, in *ApiNew, opts ...grpc.CallOption) (*ApiDetails, error)
		// 更新接口
		UpdateApi(ctx context.Context, in *ApiNew, opts ...grpc.CallOption) (*ApiDetails, error)
		// 删除接口
		DeleteApi(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询接口列表
		FindApiList(ctx context.Context, in *FindApiListReq, opts ...grpc.CallOption) (*FindApiListResp, error)
		// 同步接口列表
		SyncApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 清空接口列表
		CleanApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 创建菜单
		AddMenu(ctx context.Context, in *MenuNew, opts ...grpc.CallOption) (*MenuDetails, error)
		// 更新菜单
		UpdateMenu(ctx context.Context, in *MenuNew, opts ...grpc.CallOption) (*MenuDetails, error)
		// 删除菜单
		DeleteMenu(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询菜单列表
		FindMenuList(ctx context.Context, in *FindMenuListReq, opts ...grpc.CallOption) (*FindMenuListResp, error)
		// 同步菜单列表
		SyncMenuList(ctx context.Context, in *SyncMenuReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 清空菜单列表
		CleanMenuList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 创建角色
		AddRole(ctx context.Context, in *RoleNew, opts ...grpc.CallOption) (*RoleDetails, error)
		// 更新角色
		UpdateRole(ctx context.Context, in *RoleNew, opts ...grpc.CallOption) (*RoleDetails, error)
		// 删除角色
		DeleteRole(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询角色列表
		FindRoleList(ctx context.Context, in *FindRoleListReq, opts ...grpc.CallOption) (*FindRoleListResp, error)
		// 更新角色菜单
		UpdateRoleMenus(ctx context.Context, in *UpdateRoleMenusReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 更新角色资源
		UpdateRoleApis(ctx context.Context, in *UpdateRoleApisReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 查询角色资源权限
		FindRoleResources(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleResourcesResp, error)
		// 修改用户角色
		UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 获取用户接口权限
		FindUserApis(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindApiListResp, error)
		// 获取用户菜单权限
		FindUserMenus(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindMenuListResp, error)
		// 获取用户角色信息
		FindUserRoles(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindRoleListResp, error)
	}

	defaultPermissionRpc struct {
		cli zrpc.Client
	}
)

func NewPermissionRpc(cli zrpc.Client) PermissionRpc {
	return &defaultPermissionRpc{
		cli: cli,
	}
}

// 创建接口
func (m *defaultPermissionRpc) AddApi(ctx context.Context, in *ApiNew, opts ...grpc.CallOption) (*ApiDetails, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.AddApi(ctx, in, opts...)
}

// 更新接口
func (m *defaultPermissionRpc) UpdateApi(ctx context.Context, in *ApiNew, opts ...grpc.CallOption) (*ApiDetails, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.UpdateApi(ctx, in, opts...)
}

// 删除接口
func (m *defaultPermissionRpc) DeleteApi(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

// 查询接口列表
func (m *defaultPermissionRpc) FindApiList(ctx context.Context, in *FindApiListReq, opts ...grpc.CallOption) (*FindApiListResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.FindApiList(ctx, in, opts...)
}

// 同步接口列表
func (m *defaultPermissionRpc) SyncApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.SyncApiList(ctx, in, opts...)
}

// 清空接口列表
func (m *defaultPermissionRpc) CleanApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.CleanApiList(ctx, in, opts...)
}

// 创建菜单
func (m *defaultPermissionRpc) AddMenu(ctx context.Context, in *MenuNew, opts ...grpc.CallOption) (*MenuDetails, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.AddMenu(ctx, in, opts...)
}

// 更新菜单
func (m *defaultPermissionRpc) UpdateMenu(ctx context.Context, in *MenuNew, opts ...grpc.CallOption) (*MenuDetails, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

// 删除菜单
func (m *defaultPermissionRpc) DeleteMenu(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

// 查询菜单列表
func (m *defaultPermissionRpc) FindMenuList(ctx context.Context, in *FindMenuListReq, opts ...grpc.CallOption) (*FindMenuListResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.FindMenuList(ctx, in, opts...)
}

// 同步菜单列表
func (m *defaultPermissionRpc) SyncMenuList(ctx context.Context, in *SyncMenuReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.SyncMenuList(ctx, in, opts...)
}

// 清空菜单列表
func (m *defaultPermissionRpc) CleanMenuList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.CleanMenuList(ctx, in, opts...)
}

// 创建角色
func (m *defaultPermissionRpc) AddRole(ctx context.Context, in *RoleNew, opts ...grpc.CallOption) (*RoleDetails, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.AddRole(ctx, in, opts...)
}

// 更新角色
func (m *defaultPermissionRpc) UpdateRole(ctx context.Context, in *RoleNew, opts ...grpc.CallOption) (*RoleDetails, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

// 删除角色
func (m *defaultPermissionRpc) DeleteRole(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

// 查询角色列表
func (m *defaultPermissionRpc) FindRoleList(ctx context.Context, in *FindRoleListReq, opts ...grpc.CallOption) (*FindRoleListResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.FindRoleList(ctx, in, opts...)
}

// 更新角色菜单
func (m *defaultPermissionRpc) UpdateRoleMenus(ctx context.Context, in *UpdateRoleMenusReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.UpdateRoleMenus(ctx, in, opts...)
}

// 更新角色资源
func (m *defaultPermissionRpc) UpdateRoleApis(ctx context.Context, in *UpdateRoleApisReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.UpdateRoleApis(ctx, in, opts...)
}

// 查询角色资源权限
func (m *defaultPermissionRpc) FindRoleResources(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleResourcesResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.FindRoleResources(ctx, in, opts...)
}

// 修改用户角色
func (m *defaultPermissionRpc) UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.UpdateUserRole(ctx, in, opts...)
}

// 获取用户接口权限
func (m *defaultPermissionRpc) FindUserApis(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindApiListResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.FindUserApis(ctx, in, opts...)
}

// 获取用户菜单权限
func (m *defaultPermissionRpc) FindUserMenus(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindMenuListResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.FindUserMenus(ctx, in, opts...)
}

// 获取用户角色信息
func (m *defaultPermissionRpc) FindUserRoles(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindRoleListResp, error) {
	client := permissionrpc.NewPermissionRpcClient(m.cli.Conn())
	return client.FindUserRoles(ctx, in, opts...)
}
