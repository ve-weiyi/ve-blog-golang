// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package categoryrpc

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
	CommentDetails           = blog.CommentDetails
	CommentDetailsPageResp   = blog.CommentDetailsPageResp
	CommentPageResp          = blog.CommentPageResp
	CommentReply             = blog.CommentReply
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

	CategoryRpc interface {
		// 创建文章分类
		CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
		// 更新文章分类
		UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
		// 删除文章分类
		DeleteCategory(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除文章分类
		DeleteCategoryList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询文章分类
		FindCategory(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Category, error)
		// 分页获取文章分类列表
		FindCategoryList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CategoryPageResp, error)
		// 查询文章分类数量
		FindCategoryCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
	}

	defaultCategoryRpc struct {
		cli zrpc.Client
	}
)

func NewCategoryRpc(cli zrpc.Client) CategoryRpc {
	return &defaultCategoryRpc{
		cli: cli,
	}
}

// 创建文章分类
func (m *defaultCategoryRpc) CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	client := blog.NewCategoryRpcClient(m.cli.Conn())
	return client.CreateCategory(ctx, in, opts...)
}

// 更新文章分类
func (m *defaultCategoryRpc) UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	client := blog.NewCategoryRpcClient(m.cli.Conn())
	return client.UpdateCategory(ctx, in, opts...)
}

// 删除文章分类
func (m *defaultCategoryRpc) DeleteCategory(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewCategoryRpcClient(m.cli.Conn())
	return client.DeleteCategory(ctx, in, opts...)
}

// 批量删除文章分类
func (m *defaultCategoryRpc) DeleteCategoryList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewCategoryRpcClient(m.cli.Conn())
	return client.DeleteCategoryList(ctx, in, opts...)
}

// 查询文章分类
func (m *defaultCategoryRpc) FindCategory(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Category, error) {
	client := blog.NewCategoryRpcClient(m.cli.Conn())
	return client.FindCategory(ctx, in, opts...)
}

// 分页获取文章分类列表
func (m *defaultCategoryRpc) FindCategoryList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CategoryPageResp, error) {
	client := blog.NewCategoryRpcClient(m.cli.Conn())
	return client.FindCategoryList(ctx, in, opts...)
}

// 查询文章分类数量
func (m *defaultCategoryRpc) FindCategoryCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewCategoryRpcClient(m.cli.Conn())
	return client.FindCategoryCount(ctx, in, opts...)
}
