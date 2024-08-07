// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package photorpc

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

	PhotoRpc interface {
		// 创建照片
		CreatePhoto(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Photo, error)
		// 更新照片
		UpdatePhoto(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Photo, error)
		// 删除照片
		DeletePhoto(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除照片
		DeletePhotoList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询照片
		FindPhoto(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Photo, error)
		// 查询照片列表
		FindPhotoList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PhotoPageResp, error)
		// 查询照片数量
		FindPhotoCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
		// 创建相册
		CreatePhotoAlbum(ctx context.Context, in *PhotoAlbum, opts ...grpc.CallOption) (*PhotoAlbum, error)
		// 更新相册
		UpdatePhotoAlbum(ctx context.Context, in *PhotoAlbum, opts ...grpc.CallOption) (*PhotoAlbum, error)
		// 删除相册
		DeletePhotoAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除相册
		DeletePhotoAlbumList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询相册
		FindPhotoAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PhotoAlbum, error)
		// 查询相册列表
		FindPhotoAlbumList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PhotoAlbumPageResp, error)
		// 查询相册数量
		FindPhotoAlbumCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
	}

	defaultPhotoRpc struct {
		cli zrpc.Client
	}
)

func NewPhotoRpc(cli zrpc.Client) PhotoRpc {
	return &defaultPhotoRpc{
		cli: cli,
	}
}

// 创建照片
func (m *defaultPhotoRpc) CreatePhoto(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Photo, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.CreatePhoto(ctx, in, opts...)
}

// 更新照片
func (m *defaultPhotoRpc) UpdatePhoto(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Photo, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.UpdatePhoto(ctx, in, opts...)
}

// 删除照片
func (m *defaultPhotoRpc) DeletePhoto(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.DeletePhoto(ctx, in, opts...)
}

// 批量删除照片
func (m *defaultPhotoRpc) DeletePhotoList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.DeletePhotoList(ctx, in, opts...)
}

// 查询照片
func (m *defaultPhotoRpc) FindPhoto(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Photo, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.FindPhoto(ctx, in, opts...)
}

// 查询照片列表
func (m *defaultPhotoRpc) FindPhotoList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PhotoPageResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.FindPhotoList(ctx, in, opts...)
}

// 查询照片数量
func (m *defaultPhotoRpc) FindPhotoCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.FindPhotoCount(ctx, in, opts...)
}

// 创建相册
func (m *defaultPhotoRpc) CreatePhotoAlbum(ctx context.Context, in *PhotoAlbum, opts ...grpc.CallOption) (*PhotoAlbum, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.CreatePhotoAlbum(ctx, in, opts...)
}

// 更新相册
func (m *defaultPhotoRpc) UpdatePhotoAlbum(ctx context.Context, in *PhotoAlbum, opts ...grpc.CallOption) (*PhotoAlbum, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.UpdatePhotoAlbum(ctx, in, opts...)
}

// 删除相册
func (m *defaultPhotoRpc) DeletePhotoAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.DeletePhotoAlbum(ctx, in, opts...)
}

// 批量删除相册
func (m *defaultPhotoRpc) DeletePhotoAlbumList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.DeletePhotoAlbumList(ctx, in, opts...)
}

// 查询相册
func (m *defaultPhotoRpc) FindPhotoAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PhotoAlbum, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.FindPhotoAlbum(ctx, in, opts...)
}

// 查询相册列表
func (m *defaultPhotoRpc) FindPhotoAlbumList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PhotoAlbumPageResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.FindPhotoAlbumList(ctx, in, opts...)
}

// 查询相册数量
func (m *defaultPhotoRpc) FindPhotoAlbumCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewPhotoRpcClient(m.cli.Conn())
	return client.FindPhotoAlbumCount(ctx, in, opts...)
}
