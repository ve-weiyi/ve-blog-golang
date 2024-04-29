// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package commentrpc

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
	UploadRecordReq          = blog.UploadRecordReq
	UploadRecordResp         = blog.UploadRecordResp
	User                     = blog.User
	UserEmailReq             = blog.UserEmailReq
	UserInfoPageResp         = blog.UserInfoPageResp
	UserInfoResp             = blog.UserInfoResp
	UserReq                  = blog.UserReq

	CommentRpc interface {
		// 创建评论
		CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error)
		// 更新评论
		UpdateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error)
		// 删除评论
		DeleteComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除评论
		DeleteCommentList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询评论
		FindComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Comment, error)
		// 分页获取评论列表
		FindCommentList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CommentPageResp, error)
		// 分页获取评论回复列表
		FindCommentReplyList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CommentReplyPageResp, error)
		// 查询评论数量
		FindCommentCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
		// 点赞评论
		LikeComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error)
	}

	defaultCommentRpc struct {
		cli zrpc.Client
	}
)

func NewCommentRpc(cli zrpc.Client) CommentRpc {
	return &defaultCommentRpc{
		cli: cli,
	}
}

// 创建评论
func (m *defaultCommentRpc) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.CreateComment(ctx, in, opts...)
}

// 更新评论
func (m *defaultCommentRpc) UpdateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.UpdateComment(ctx, in, opts...)
}

// 删除评论
func (m *defaultCommentRpc) DeleteComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

// 批量删除评论
func (m *defaultCommentRpc) DeleteCommentList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.DeleteCommentList(ctx, in, opts...)
}

// 查询评论
func (m *defaultCommentRpc) FindComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Comment, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.FindComment(ctx, in, opts...)
}

// 分页获取评论列表
func (m *defaultCommentRpc) FindCommentList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CommentPageResp, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.FindCommentList(ctx, in, opts...)
}

// 分页获取评论回复列表
func (m *defaultCommentRpc) FindCommentReplyList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CommentReplyPageResp, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.FindCommentReplyList(ctx, in, opts...)
}

// 查询评论数量
func (m *defaultCommentRpc) FindCommentCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.FindCommentCount(ctx, in, opts...)
}

// 点赞评论
func (m *defaultCommentRpc) LikeComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewCommentRpcClient(m.cli.Conn())
	return client.LikeComment(ctx, in, opts...)
}
