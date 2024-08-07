// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package articlerpc

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

	ArticleRpc interface {
		// 创建文章
		CreateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
		// 更新文章
		UpdateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
		// 删除文章
		DeleteArticle(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除文章
		DeleteArticleList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询文章
		FindArticle(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Article, error)
		// 查询文章列表
		FindArticleList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ArticlePageResp, error)
		// 查询文章数量
		FindArticleCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
		// 查询文章列表
		FindArticleByTag(ctx context.Context, in *FindArticleByTagReq, opts ...grpc.CallOption) (*ArticlePageResp, error)
		// 查询文章列表
		FindArticleByCategory(ctx context.Context, in *FindArticleByCategoryReq, opts ...grpc.CallOption) (*ArticlePageResp, error)
	}

	defaultArticleRpc struct {
		cli zrpc.Client
	}
)

func NewArticleRpc(cli zrpc.Client) ArticleRpc {
	return &defaultArticleRpc{
		cli: cli,
	}
}

// 创建文章
func (m *defaultArticleRpc) CreateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.CreateArticle(ctx, in, opts...)
}

// 更新文章
func (m *defaultArticleRpc) UpdateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.UpdateArticle(ctx, in, opts...)
}

// 删除文章
func (m *defaultArticleRpc) DeleteArticle(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.DeleteArticle(ctx, in, opts...)
}

// 批量删除文章
func (m *defaultArticleRpc) DeleteArticleList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.DeleteArticleList(ctx, in, opts...)
}

// 查询文章
func (m *defaultArticleRpc) FindArticle(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Article, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.FindArticle(ctx, in, opts...)
}

// 查询文章列表
func (m *defaultArticleRpc) FindArticleList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*ArticlePageResp, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.FindArticleList(ctx, in, opts...)
}

// 查询文章数量
func (m *defaultArticleRpc) FindArticleCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.FindArticleCount(ctx, in, opts...)
}

// 查询文章列表
func (m *defaultArticleRpc) FindArticleByTag(ctx context.Context, in *FindArticleByTagReq, opts ...grpc.CallOption) (*ArticlePageResp, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.FindArticleByTag(ctx, in, opts...)
}

// 查询文章列表
func (m *defaultArticleRpc) FindArticleByCategory(ctx context.Context, in *FindArticleByCategoryReq, opts ...grpc.CallOption) (*ArticlePageResp, error) {
	client := blog.NewArticleRpcClient(m.cli.Conn())
	return client.FindArticleByCategory(ctx, in, opts...)
}
