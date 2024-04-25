// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package tagrpc

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

	TagRpc interface {
		// 创建标签
		CreateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
		// 更新标签
		UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
		// 删除标签
		DeleteTag(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除标签
		DeleteTagList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询标签
		FindTag(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Tag, error)
		// 分页获取标签列表
		FindTagList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*TagPageResp, error)
		// 查询文章标签数量
		FindTagCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
		// 查询标签关联文章数量
		FindTagArticleCount(ctx context.Context, in *FindTagArticleCountReq, opts ...grpc.CallOption) (*CountResp, error)
		// 查询文章标签列表
		FindTagListByArticleId(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*TagPageResp, error)
	}

	defaultTagRpc struct {
		cli zrpc.Client
	}
)

func NewTagRpc(cli zrpc.Client) TagRpc {
	return &defaultTagRpc{
		cli: cli,
	}
}

// 创建标签
func (m *defaultTagRpc) CreateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.CreateTag(ctx, in, opts...)
}

// 更新标签
func (m *defaultTagRpc) UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.UpdateTag(ctx, in, opts...)
}

// 删除标签
func (m *defaultTagRpc) DeleteTag(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.DeleteTag(ctx, in, opts...)
}

// 批量删除标签
func (m *defaultTagRpc) DeleteTagList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.DeleteTagList(ctx, in, opts...)
}

// 查询标签
func (m *defaultTagRpc) FindTag(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Tag, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTag(ctx, in, opts...)
}

// 分页获取标签列表
func (m *defaultTagRpc) FindTagList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*TagPageResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagList(ctx, in, opts...)
}

// 查询文章标签数量
func (m *defaultTagRpc) FindTagCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagCount(ctx, in, opts...)
}

// 查询标签关联文章数量
func (m *defaultTagRpc) FindTagArticleCount(ctx context.Context, in *FindTagArticleCountReq, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagArticleCount(ctx, in, opts...)
}

// 查询文章标签列表
func (m *defaultTagRpc) FindTagListByArticleId(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*TagPageResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagListByArticleId(ctx, in, opts...)
}
