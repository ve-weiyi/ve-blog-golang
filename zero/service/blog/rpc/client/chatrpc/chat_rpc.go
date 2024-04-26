// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package chatrpc

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
	PageCondition            = blog.PageCondition
	PageLimit                = blog.PageLimit
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
	User                     = blog.User
	UserEmailReq             = blog.UserEmailReq
	UserInfoPageResp         = blog.UserInfoPageResp
	UserInfoResp             = blog.UserInfoResp
	UserReq                  = blog.UserReq

	ChatRpc interface {
		// 创建聊天记录
		CreateChatRecord(ctx context.Context, in *ChatRecord, opts ...grpc.CallOption) (*ChatRecord, error)
		// 更新聊天记录
		UpdateChatRecord(ctx context.Context, in *ChatRecord, opts ...grpc.CallOption) (*ChatRecord, error)
		// 删除聊天记录
		DeleteChatRecord(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除聊天记录
		DeleteChatRecordList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询聊天记录
		FindChatRecord(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*ChatRecord, error)
		// 分页获取聊天记录列表
		FindChatRecordList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ChatRecordPageResp, error)
		// 查询聊天记录数量
		FindChatRecordCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
	}

	defaultChatRpc struct {
		cli zrpc.Client
	}
)

func NewChatRpc(cli zrpc.Client) ChatRpc {
	return &defaultChatRpc{
		cli: cli,
	}
}

// 创建聊天记录
func (m *defaultChatRpc) CreateChatRecord(ctx context.Context, in *ChatRecord, opts ...grpc.CallOption) (*ChatRecord, error) {
	client := blog.NewChatRpcClient(m.cli.Conn())
	return client.CreateChatRecord(ctx, in, opts...)
}

// 更新聊天记录
func (m *defaultChatRpc) UpdateChatRecord(ctx context.Context, in *ChatRecord, opts ...grpc.CallOption) (*ChatRecord, error) {
	client := blog.NewChatRpcClient(m.cli.Conn())
	return client.UpdateChatRecord(ctx, in, opts...)
}

// 删除聊天记录
func (m *defaultChatRpc) DeleteChatRecord(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewChatRpcClient(m.cli.Conn())
	return client.DeleteChatRecord(ctx, in, opts...)
}

// 批量删除聊天记录
func (m *defaultChatRpc) DeleteChatRecordList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewChatRpcClient(m.cli.Conn())
	return client.DeleteChatRecordList(ctx, in, opts...)
}

// 查询聊天记录
func (m *defaultChatRpc) FindChatRecord(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*ChatRecord, error) {
	client := blog.NewChatRpcClient(m.cli.Conn())
	return client.FindChatRecord(ctx, in, opts...)
}

// 分页获取聊天记录列表
func (m *defaultChatRpc) FindChatRecordList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ChatRecordPageResp, error) {
	client := blog.NewChatRpcClient(m.cli.Conn())
	return client.FindChatRecordList(ctx, in, opts...)
}

// 查询聊天记录数量
func (m *defaultChatRpc) FindChatRecordCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewChatRpcClient(m.cli.Conn())
	return client.FindChatRecordCount(ctx, in, opts...)
}
