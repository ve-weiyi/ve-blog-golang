// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package configrpc

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
	OperationLog           = blog.OperationLog
	OperationLogPageResp   = blog.OperationLogPageResp
	PageCondition          = blog.PageCondition
	PageLimit              = blog.PageLimit
	PageQuery              = blog.PageQuery
	PageResp               = blog.PageResp
	PageSort               = blog.PageSort
	Photo                  = blog.Photo
	PhotoAlbum             = blog.PhotoAlbum
	PhotoPageResp          = blog.PhotoPageResp
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
	Talk                   = blog.Talk
	TalkDetailsDTO         = blog.TalkDetailsDTO
	TalkPageResp           = blog.TalkPageResp
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

	ConfigRpc interface {
		SaveConfig(ctx context.Context, in *SaveConfigReq, opts ...grpc.CallOption) (*EmptyResp, error)
		FindConfig(ctx context.Context, in *FindConfigReq, opts ...grpc.CallOption) (*FindConfigResp, error)
	}

	defaultConfigRpc struct {
		cli zrpc.Client
	}
)

func NewConfigRpc(cli zrpc.Client) ConfigRpc {
	return &defaultConfigRpc{
		cli: cli,
	}
}

func (m *defaultConfigRpc) SaveConfig(ctx context.Context, in *SaveConfigReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewConfigRpcClient(m.cli.Conn())
	return client.SaveConfig(ctx, in, opts...)
}

func (m *defaultConfigRpc) FindConfig(ctx context.Context, in *FindConfigReq, opts ...grpc.CallOption) (*FindConfigResp, error) {
	client := blog.NewConfigRpcClient(m.cli.Conn())
	return client.FindConfig(ctx, in, opts...)
}
