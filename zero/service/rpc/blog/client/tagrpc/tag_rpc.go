// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package tagrpc

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

	TagRpc interface {
		// 创建标签
		AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
		// 更新标签
		UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
		// 删除标签
		DeleteTag(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除标签
		DeleteTagList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询标签
		FindTag(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Tag, error)
		// 查询标签列表
		FindTagList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*TagPageResp, error)
		// 查询标签数量
		FindTagCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
		// 查询标签关联文章数量
		FindTagArticleCount(ctx context.Context, in *FindTagArticleCountReq, opts ...grpc.CallOption) (*CountResp, error)
		// 查询文章标签列表(通过文章ids)
		FindTagMapByArticleIds(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*TagMapResp, error)
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
func (m *defaultTagRpc) AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.AddTag(ctx, in, opts...)
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

// 查询标签列表
func (m *defaultTagRpc) FindTagList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*TagPageResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagList(ctx, in, opts...)
}

// 查询标签数量
func (m *defaultTagRpc) FindTagCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagCount(ctx, in, opts...)
}

// 查询标签关联文章数量
func (m *defaultTagRpc) FindTagArticleCount(ctx context.Context, in *FindTagArticleCountReq, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagArticleCount(ctx, in, opts...)
}

// 查询文章标签列表(通过文章ids)
func (m *defaultTagRpc) FindTagMapByArticleIds(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*TagMapResp, error) {
	client := blog.NewTagRpcClient(m.cli.Conn())
	return client.FindTagMapByArticleIds(ctx, in, opts...)
}
