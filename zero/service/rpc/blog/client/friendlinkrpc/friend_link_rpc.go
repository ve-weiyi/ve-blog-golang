// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package friendlinkrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

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
	CommentUserInfo          = blog.CommentUserInfo
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
	GetLogoutAtReq           = blog.GetLogoutAtReq
	IdReq                    = blog.IdReq
	IdsReq                   = blog.IdsReq
	LoginHistory             = blog.LoginHistory
	LoginHistoryPageResp     = blog.LoginHistoryPageResp
	LoginReq                 = blog.LoginReq
	LoginResp                = blog.LoginResp
	LogoffReq                = blog.LogoffReq
	LogoutReq                = blog.LogoutReq
	LogoutResp               = blog.LogoutResp
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
	UpdateUserInfoReq        = blog.UpdateUserInfoReq
	UpdateUserRoleReq        = blog.UpdateUserRoleReq
	UpdateUserStatusReq      = blog.UpdateUserStatusReq
	UploadRecordReq          = blog.UploadRecordReq
	UploadRecordResp         = blog.UploadRecordResp
	User                     = blog.User
	UserEmailReq             = blog.UserEmailReq
	UserInfoResp             = blog.UserInfoResp
	UserPageResp             = blog.UserPageResp
	UserReq                  = blog.UserReq
	UserVisit                = blog.UserVisit
	UserVisitPageRsp         = blog.UserVisitPageRsp

	FriendLinkRpc interface {
		// 创建友链
		AddFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error)
		// 更新友链
		UpdateFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error)
		// 删除友链
		DeleteFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除友链
		DeleteFriendLinkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询友链
		FindFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*FriendLink, error)
		// 查询友链列表
		FindFriendLinkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*FriendLinkPageResp, error)
		// 查询友链数量
		FindFriendLinkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
	}

	defaultFriendLinkRpc struct {
		cli zrpc.Client
	}
)

func NewFriendLinkRpc(cli zrpc.Client) FriendLinkRpc {
	return &defaultFriendLinkRpc{
		cli: cli,
	}
}

// 创建友链
func (m *defaultFriendLinkRpc) AddFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.AddFriendLink(ctx, in, opts...)
}

// 更新友链
func (m *defaultFriendLinkRpc) UpdateFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.UpdateFriendLink(ctx, in, opts...)
}

// 删除友链
func (m *defaultFriendLinkRpc) DeleteFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.DeleteFriendLink(ctx, in, opts...)
}

// 批量删除友链
func (m *defaultFriendLinkRpc) DeleteFriendLinkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.DeleteFriendLinkList(ctx, in, opts...)
}

// 查询友链
func (m *defaultFriendLinkRpc) FindFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*FriendLink, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.FindFriendLink(ctx, in, opts...)
}

// 查询友链列表
func (m *defaultFriendLinkRpc) FindFriendLinkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*FriendLinkPageResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.FindFriendLinkList(ctx, in, opts...)
}

// 查询友链数量
func (m *defaultFriendLinkRpc) FindFriendLinkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.FindFriendLinkCount(ctx, in, opts...)
}
