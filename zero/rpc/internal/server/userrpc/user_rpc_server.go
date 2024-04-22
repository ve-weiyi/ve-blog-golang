// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/logic/userrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

type UserRpcServer struct {
	svcCtx *svc.ServiceContext
	account.UnimplementedUserRpcServer
}

func NewUserRpcServer(svcCtx *svc.ServiceContext) *UserRpcServer {
	return &UserRpcServer{
		svcCtx: svcCtx,
	}
}

// 查询用户登录历史
func (s *UserRpcServer) FindUserLoginHistoryList(ctx context.Context, in *account.PageQuery) (*account.LoginHistoryPageResp, error) {
	l := userrpclogic.NewFindUserLoginHistoryListLogic(ctx, s.svcCtx)
	return l.FindUserLoginHistoryList(in)
}

// 批量删除登录历史
func (s *UserRpcServer) DeleteUserLoginHistoryList(ctx context.Context, in *account.IdsReq) (*account.BatchResult, error) {
	l := userrpclogic.NewDeleteUserLoginHistoryListLogic(ctx, s.svcCtx)
	return l.DeleteUserLoginHistoryList(in)
}

// 获取用户接口权限
func (s *UserRpcServer) GetUserApis(ctx context.Context, in *account.EmptyReq) (*account.ApiPageResp, error) {
	l := userrpclogic.NewGetUserApisLogic(ctx, s.svcCtx)
	return l.GetUserApis(in)
}

// 获取用户菜单权限
func (s *UserRpcServer) GetUserMenus(ctx context.Context, in *account.EmptyReq) (*account.MenuPageResp, error) {
	l := userrpclogic.NewGetUserMenusLogic(ctx, s.svcCtx)
	return l.GetUserMenus(in)
}

// 获取用户角色信息
func (s *UserRpcServer) GetUserRoles(ctx context.Context, in *account.EmptyReq) (*account.RolePageResp, error) {
	l := userrpclogic.NewGetUserRolesLogic(ctx, s.svcCtx)
	return l.GetUserRoles(in)
}

// 获取用户信息
func (s *UserRpcServer) GetUserInfo(ctx context.Context, in *account.EmptyReq) (*account.UserInfoResp, error) {
	l := userrpclogic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

// 修改用户信息
func (s *UserRpcServer) UpdateUserInfo(ctx context.Context, in *account.UpdateUserInfoReq) (*account.UserInfoResp, error) {
	l := userrpclogic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

// 修改用户头像
func (s *UserRpcServer) UpdateUserAvatar(ctx context.Context, in *account.UpdateUserAvatarReq) (*account.UserInfoResp, error) {
	l := userrpclogic.NewUpdateUserAvatarLogic(ctx, s.svcCtx)
	return l.UpdateUserAvatar(in)
}

// 修改用户状态
func (s *UserRpcServer) UpdateUserStatus(ctx context.Context, in *account.UpdateUserStatusReq) (*account.EmptyResp, error) {
	l := userrpclogic.NewUpdateUserStatusLogic(ctx, s.svcCtx)
	return l.UpdateUserStatus(in)
}

// 修改用户角色
func (s *UserRpcServer) UpdateUserRole(ctx context.Context, in *account.UpdateUserRoleReq) (*account.EmptyResp, error) {
	l := userrpclogic.NewUpdateUserRoleLogic(ctx, s.svcCtx)
	return l.UpdateUserRole(in)
}

// 查找用户列表
func (s *UserRpcServer) FindUserList(ctx context.Context, in *account.PageQuery) (*account.PageUserInfoResp, error) {
	l := userrpclogic.NewFindUserListLogic(ctx, s.svcCtx)
	return l.FindUserList(in)
}
