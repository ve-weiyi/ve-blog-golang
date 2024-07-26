// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package talkrpc

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

	TalkRpc interface {
		// 创建说说
		AddTalk(ctx context.Context, in *Talk, opts ...grpc.CallOption) (*Talk, error)
		// 更新说说
		UpdateTalk(ctx context.Context, in *Talk, opts ...grpc.CallOption) (*Talk, error)
		// 删除说说
		DeleteTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除说说
		DeleteTalkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询说说
		FindTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Talk, error)
		// 查询说说列表
		FindTalkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*TalkPageResp, error)
		// 查询说说数量
		FindTalkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
		// 点赞说说
		LikeTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error)
	}

	defaultTalkRpc struct {
		cli zrpc.Client
	}
)

func NewTalkRpc(cli zrpc.Client) TalkRpc {
	return &defaultTalkRpc{
		cli: cli,
	}
}

// 创建说说
func (m *defaultTalkRpc) AddTalk(ctx context.Context, in *Talk, opts ...grpc.CallOption) (*Talk, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.AddTalk(ctx, in, opts...)
}

// 更新说说
func (m *defaultTalkRpc) UpdateTalk(ctx context.Context, in *Talk, opts ...grpc.CallOption) (*Talk, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.UpdateTalk(ctx, in, opts...)
}

// 删除说说
func (m *defaultTalkRpc) DeleteTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.DeleteTalk(ctx, in, opts...)
}

// 批量删除说说
func (m *defaultTalkRpc) DeleteTalkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.DeleteTalkList(ctx, in, opts...)
}

// 查询说说
func (m *defaultTalkRpc) FindTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Talk, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.FindTalk(ctx, in, opts...)
}

// 查询说说列表
func (m *defaultTalkRpc) FindTalkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*TalkPageResp, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.FindTalkList(ctx, in, opts...)
}

// 查询说说数量
func (m *defaultTalkRpc) FindTalkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.FindTalkCount(ctx, in, opts...)
}

// 点赞说说
func (m *defaultTalkRpc) LikeTalk(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewTalkRpcClient(m.cli.Conn())
	return client.LikeTalk(ctx, in, opts...)
}
