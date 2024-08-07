// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package remarkrpc

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
	UpdateUserAvatarReq      = blog.UpdateUserAvatarReq
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

	RemarkRpc interface {
		// 创建留言
		CreateRemark(ctx context.Context, in *Remark, opts ...grpc.CallOption) (*Remark, error)
		// 更新留言
		UpdateRemark(ctx context.Context, in *Remark, opts ...grpc.CallOption) (*Remark, error)
		// 删除留言
		DeleteRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除留言
		DeleteRemarkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询留言
		FindRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Remark, error)
		// 查询留言列表
		FindRemarkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*RemarkPageResp, error)
		// 查询留言数量
		FindRemarkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
	}

	defaultRemarkRpc struct {
		cli zrpc.Client
	}
)

func NewRemarkRpc(cli zrpc.Client) RemarkRpc {
	return &defaultRemarkRpc{
		cli: cli,
	}
}

// 创建留言
func (m *defaultRemarkRpc) CreateRemark(ctx context.Context, in *Remark, opts ...grpc.CallOption) (*Remark, error) {
	client := blog.NewRemarkRpcClient(m.cli.Conn())
	return client.CreateRemark(ctx, in, opts...)
}

// 更新留言
func (m *defaultRemarkRpc) UpdateRemark(ctx context.Context, in *Remark, opts ...grpc.CallOption) (*Remark, error) {
	client := blog.NewRemarkRpcClient(m.cli.Conn())
	return client.UpdateRemark(ctx, in, opts...)
}

// 删除留言
func (m *defaultRemarkRpc) DeleteRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewRemarkRpcClient(m.cli.Conn())
	return client.DeleteRemark(ctx, in, opts...)
}

// 批量删除留言
func (m *defaultRemarkRpc) DeleteRemarkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewRemarkRpcClient(m.cli.Conn())
	return client.DeleteRemarkList(ctx, in, opts...)
}

// 查询留言
func (m *defaultRemarkRpc) FindRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Remark, error) {
	client := blog.NewRemarkRpcClient(m.cli.Conn())
	return client.FindRemark(ctx, in, opts...)
}

// 查询留言列表
func (m *defaultRemarkRpc) FindRemarkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*RemarkPageResp, error) {
	client := blog.NewRemarkRpcClient(m.cli.Conn())
	return client.FindRemarkList(ctx, in, opts...)
}

// 查询留言数量
func (m *defaultRemarkRpc) FindRemarkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewRemarkRpcClient(m.cli.Conn())
	return client.FindRemarkCount(ctx, in, opts...)
}
