// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package userrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Api                      = blog.Api
	ApiDetails               = blog.ApiDetails
	ApiPageResp              = blog.ApiPageResp
	Article                  = blog.Article
	ArticlePageResp          = blog.ArticlePageResp
	BatchResp                = blog.BatchResp
	Category                 = blog.Category
	CategoryPageResp         = blog.CategoryPageResp
	ChatRecord               = blog.ChatRecord
	ChatRecordPageResp       = blog.ChatRecordPageResp
	Comment                  = blog.Comment
	CommentPageResp          = blog.CommentPageResp
	CommentReply             = blog.CommentReply
	CommentReplyPageResp     = blog.CommentReplyPageResp
	CountResp                = blog.CountResp
	EmptyReq                 = blog.EmptyReq
	EmptyResp                = blog.EmptyResp
	FindArticleByCategoryReq = blog.FindArticleByCategoryReq
	FindArticleByTagReq      = blog.FindArticleByTagReq
	FindCategoryByNameReq    = blog.FindCategoryByNameReq
	FindConfigReq            = blog.FindConfigReq
	FindConfigResp           = blog.FindConfigResp
	FindTagArticleCountReq   = blog.FindTagArticleCountReq
	FindTagByNameReq         = blog.FindTagByNameReq
	FriendLink               = blog.FriendLink
	FriendLinkPageResp       = blog.FriendLinkPageResp
	IdReq                    = blog.IdReq
	IdsReq                   = blog.IdsReq
	LoginHistory             = blog.LoginHistory
	LoginHistoryPageResp     = blog.LoginHistoryPageResp
	LoginReq                 = blog.LoginReq
	LoginResp                = blog.LoginResp
	LogoffReq                = blog.LogoffReq
	LogoutReq                = blog.LogoutReq
	Menu                     = blog.Menu
	MenuDetails              = blog.MenuDetails
	MenuPageResp             = blog.MenuPageResp
	OauthLoginReq            = blog.OauthLoginReq
	OauthLoginUrlResp        = blog.OauthLoginUrlResp
	OperationLog             = blog.OperationLog
	OperationLogPageResp     = blog.OperationLogPageResp
	Page                     = blog.Page
	PageCondition            = blog.PageCondition
	PageLimit                = blog.PageLimit
	PagePageResp             = blog.PagePageResp
	PageQuery                = blog.PageQuery
	PageResp                 = blog.PageResp
	PageSort                 = blog.PageSort
	Photo                    = blog.Photo
	PhotoAlbum               = blog.PhotoAlbum
	PhotoAlbumPageResp       = blog.PhotoAlbumPageResp
	PhotoPageResp            = blog.PhotoPageResp
	Remark                   = blog.Remark
	RemarkPageResp           = blog.RemarkPageResp
	ResetPasswordReq         = blog.ResetPasswordReq
	Role                     = blog.Role
	RoleDetails              = blog.RoleDetails
	RoleLabel                = blog.RoleLabel
	RolePageResp             = blog.RolePageResp
	RoleResourcesResp        = blog.RoleResourcesResp
	SaveConfigReq            = blog.SaveConfigReq
	SyncMenuReq              = blog.SyncMenuReq
	Tag                      = blog.Tag
	TagMapResp               = blog.TagMapResp
	TagPageResp              = blog.TagPageResp
	Talk                     = blog.Talk
	TalkDetailsDTO           = blog.TalkDetailsDTO
	TalkPageResp             = blog.TalkPageResp
	UpdateRoleApisReq        = blog.UpdateRoleApisReq
	UpdateRoleMenusReq       = blog.UpdateRoleMenusReq
	UpdateUserAvatarReq      = blog.UpdateUserAvatarReq
	UpdateUserInfoReq        = blog.UpdateUserInfoReq
	UpdateUserRoleReq        = blog.UpdateUserRoleReq
	UpdateUserStatusReq      = blog.UpdateUserStatusReq
	UploadRecordReq          = blog.UploadRecordReq
	UploadRecordResp         = blog.UploadRecordResp
	User                     = blog.User
	UserEmailReq             = blog.UserEmailReq
	UserInfoPageResp         = blog.UserInfoPageResp
	UserInfoResp             = blog.UserInfoResp
	UserReq                  = blog.UserReq

	UserRpc interface {
		// 查询用户登录历史
		FindUserLoginHistoryList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*LoginHistoryPageResp, error)
		// 批量删除登录历史
		DeleteUserLoginHistoryList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 获取用户接口权限
		FindUserApis(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ApiPageResp, error)
		// 获取用户菜单权限
		FindUserMenus(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*MenuPageResp, error)
		// 获取用户角色信息
		FindUserRoles(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*RolePageResp, error)
		// 获取用户信息
		FindUserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		// 修改用户信息
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		// 修改用户头像
		UpdateUserAvatar(ctx context.Context, in *UpdateUserAvatarReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		// 修改用户状态
		UpdateUserStatus(ctx context.Context, in *UpdateUserStatusReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 修改用户角色
		UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 查找用户列表
		FindUserList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*UserInfoPageResp, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

// 查询用户登录历史
func (m *defaultUserRpc) FindUserLoginHistoryList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*LoginHistoryPageResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.FindUserLoginHistoryList(ctx, in, opts...)
}

// 批量删除登录历史
func (m *defaultUserRpc) DeleteUserLoginHistoryList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.DeleteUserLoginHistoryList(ctx, in, opts...)
}

// 获取用户接口权限
func (m *defaultUserRpc) FindUserApis(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ApiPageResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.FindUserApis(ctx, in, opts...)
}

// 获取用户菜单权限
func (m *defaultUserRpc) FindUserMenus(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*MenuPageResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.FindUserMenus(ctx, in, opts...)
}

// 获取用户角色信息
func (m *defaultUserRpc) FindUserRoles(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*RolePageResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.FindUserRoles(ctx, in, opts...)
}

// 获取用户信息
func (m *defaultUserRpc) FindUserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.FindUserInfo(ctx, in, opts...)
}

// 修改用户信息
func (m *defaultUserRpc) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

// 修改用户头像
func (m *defaultUserRpc) UpdateUserAvatar(ctx context.Context, in *UpdateUserAvatarReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.UpdateUserAvatar(ctx, in, opts...)
}

// 修改用户状态
func (m *defaultUserRpc) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in, opts...)
}

// 修改用户角色
func (m *defaultUserRpc) UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.UpdateUserRole(ctx, in, opts...)
}

// 查找用户列表
func (m *defaultUserRpc) FindUserList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*UserInfoPageResp, error) {
	client := blog.NewUserRpcClient(m.cli.Conn())
	return client.FindUserList(ctx, in, opts...)
}
