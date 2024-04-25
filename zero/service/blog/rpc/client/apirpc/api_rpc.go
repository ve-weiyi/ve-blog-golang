// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package apirpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Api                    = blog.Api
	ApiDetails             = blog.ApiDetails
	ApiPageResp            = blog.ApiPageResp
	Article                = blog.Article
	ArticlePageResp        = blog.ArticlePageResp
	BatchResp              = blog.BatchResp
	Category               = blog.Category
	CategoryPageResp       = blog.CategoryPageResp
	CountResp              = blog.CountResp
	EmptyReq               = blog.EmptyReq
	EmptyResp              = blog.EmptyResp
	FindCategoryByNameReq  = blog.FindCategoryByNameReq
	FindConfigReq          = blog.FindConfigReq
	FindConfigResp         = blog.FindConfigResp
	FindTagArticleCountReq = blog.FindTagArticleCountReq
	FindTagByNameReq       = blog.FindTagByNameReq
	FriendLink             = blog.FriendLink
	FriendLinkPageResp     = blog.FriendLinkPageResp
	IdReq                  = blog.IdReq
	IdsReq                 = blog.IdsReq
	LoginHistory           = blog.LoginHistory
	LoginHistoryPageResp   = blog.LoginHistoryPageResp
	LoginReq               = blog.LoginReq
	LoginResp              = blog.LoginResp
	Menu                   = blog.Menu
	MenuDetails            = blog.MenuDetails
	MenuPageResp           = blog.MenuPageResp
	OauthLoginReq          = blog.OauthLoginReq
	OauthLoginUrlResp      = blog.OauthLoginUrlResp
	PageCondition          = blog.PageCondition
	PageLimit              = blog.PageLimit
	PageQuery              = blog.PageQuery
	PageResp               = blog.PageResp
	PageSort               = blog.PageSort
	Remark                 = blog.Remark
	RemarkPageResp         = blog.RemarkPageResp
	ResetPasswordReq       = blog.ResetPasswordReq
	Role                   = blog.Role
	RoleDetails            = blog.RoleDetails
	RoleLabel              = blog.RoleLabel
	RolePageResp           = blog.RolePageResp
	RoleResourcesResp      = blog.RoleResourcesResp
	SaveConfigReq          = blog.SaveConfigReq
	SyncMenuRequest        = blog.SyncMenuRequest
	Tag                    = blog.Tag
	TagPageResp            = blog.TagPageResp
	UpdateRoleApisReq      = blog.UpdateRoleApisReq
	UpdateRoleMenusReq     = blog.UpdateRoleMenusReq
	UpdateUserAvatarReq    = blog.UpdateUserAvatarReq
	UpdateUserInfoReq      = blog.UpdateUserInfoReq
	UpdateUserRoleReq      = blog.UpdateUserRoleReq
	UpdateUserStatusReq    = blog.UpdateUserStatusReq
	User                   = blog.User
	UserEmailReq           = blog.UserEmailReq
	UserInfoPageResp       = blog.UserInfoPageResp
	UserInfoResp           = blog.UserInfoResp

	ApiRpc interface {
		// 创建接口
		CreateApi(ctx context.Context, in *Api, opts ...grpc.CallOption) (*Api, error)
		// 更新接口
		UpdateApi(ctx context.Context, in *Api, opts ...grpc.CallOption) (*Api, error)
		// 删除接口
		DeleteApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除接口
		DeleteApiList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询接口
		FindApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Api, error)
		// 分页获取接口列表
		FindApiList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ApiPageResp, error)
		// 同步接口列表
		SyncApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 清空接口列表
		CleanApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error)
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
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.CreateApi(ctx, in, opts...)
}

// 更新接口
func (m *defaultApiRpc) UpdateApi(ctx context.Context, in *Api, opts ...grpc.CallOption) (*Api, error) {
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.UpdateApi(ctx, in, opts...)
}

// 删除接口
func (m *defaultApiRpc) DeleteApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

// 批量删除接口
func (m *defaultApiRpc) DeleteApiList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.DeleteApiList(ctx, in, opts...)
}

// 查询接口
func (m *defaultApiRpc) FindApi(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Api, error) {
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.FindApi(ctx, in, opts...)
}

// 分页获取接口列表
func (m *defaultApiRpc) FindApiList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ApiPageResp, error) {
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.FindApiList(ctx, in, opts...)
}

// 同步接口列表
func (m *defaultApiRpc) SyncApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.SyncApiList(ctx, in, opts...)
}

// 清空接口列表
func (m *defaultApiRpc) CleanApiList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewApiRpcClient(m.cli.Conn())
	return client.CleanApiList(ctx, in, opts...)
}
