// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package rolerpc

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
	SyncMenuRequest          = blog.SyncMenuRequest
	Tag                      = blog.Tag
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

	RoleRpc interface {
		// 创建角色
		CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
		// 更新角色
		UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
		// 删除角色
		DeleteRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除角色
		DeleteRoleList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询角色
		FindRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Role, error)
		// 分页获取角色列表
		FindRoleList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*RolePageResp, error)
		// 查询角色
		FindRoleResources(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleResourcesResp, error)
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
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

// 更新角色
func (m *defaultRoleRpc) UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

// 删除角色
func (m *defaultRoleRpc) DeleteRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

// 批量删除角色
func (m *defaultRoleRpc) DeleteRoleList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.DeleteRoleList(ctx, in, opts...)
}

// 查询角色
func (m *defaultRoleRpc) FindRole(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Role, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.FindRole(ctx, in, opts...)
}

// 分页获取角色列表
func (m *defaultRoleRpc) FindRoleList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*RolePageResp, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.FindRoleList(ctx, in, opts...)
}

// 查询角色
func (m *defaultRoleRpc) FindRoleResources(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleResourcesResp, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.FindRoleResources(ctx, in, opts...)
}

// 更新角色菜单
func (m *defaultRoleRpc) UpdateRoleMenus(ctx context.Context, in *UpdateRoleMenusReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.UpdateRoleMenus(ctx, in, opts...)
}

// 更新角色资源
func (m *defaultRoleRpc) UpdateRoleApis(ctx context.Context, in *UpdateRoleApisReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewRoleRpcClient(m.cli.Conn())
	return client.UpdateRoleApis(ctx, in, opts...)
}
