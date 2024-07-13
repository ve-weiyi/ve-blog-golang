// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package menurpc

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
	GetLogoutAtReq           = blog.GetLogoutAtReq
	GetLogoutAtResp          = blog.GetLogoutAtResp
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

	MenuRpc interface {
		// 创建菜单
		AddMenu(ctx context.Context, in *Menu, opts ...grpc.CallOption) (*Menu, error)
		// 更新菜单
		UpdateMenu(ctx context.Context, in *Menu, opts ...grpc.CallOption) (*Menu, error)
		// 删除菜单
		DeleteMenu(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除菜单
		DeleteMenuList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询菜单
		FindMenu(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Menu, error)
		// 查询菜单列表
		FindMenuList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*MenuPageResp, error)
		// 同步菜单列表
		SyncMenuList(ctx context.Context, in *SyncMenuReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 清空菜单列表
		CleanMenuList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error)
	}

	defaultMenuRpc struct {
		cli zrpc.Client
	}
)

func NewMenuRpc(cli zrpc.Client) MenuRpc {
	return &defaultMenuRpc{
		cli: cli,
	}
}

// 创建菜单
func (m *defaultMenuRpc) AddMenu(ctx context.Context, in *Menu, opts ...grpc.CallOption) (*Menu, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.AddMenu(ctx, in, opts...)
}

// 更新菜单
func (m *defaultMenuRpc) UpdateMenu(ctx context.Context, in *Menu, opts ...grpc.CallOption) (*Menu, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

// 删除菜单
func (m *defaultMenuRpc) DeleteMenu(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

// 批量删除菜单
func (m *defaultMenuRpc) DeleteMenuList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.DeleteMenuList(ctx, in, opts...)
}

// 查询菜单
func (m *defaultMenuRpc) FindMenu(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Menu, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.FindMenu(ctx, in, opts...)
}

// 查询菜单列表
func (m *defaultMenuRpc) FindMenuList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*MenuPageResp, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.FindMenuList(ctx, in, opts...)
}

// 同步菜单列表
func (m *defaultMenuRpc) SyncMenuList(ctx context.Context, in *SyncMenuReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.SyncMenuList(ctx, in, opts...)
}

// 清空菜单列表
func (m *defaultMenuRpc) CleanMenuList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewMenuRpcClient(m.cli.Conn())
	return client.CleanMenuList(ctx, in, opts...)
}
