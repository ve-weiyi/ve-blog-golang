// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package logrpc

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

	LogRpc interface {
		// 创建操作记录
		AddOperationLog(ctx context.Context, in *OperationLog, opts ...grpc.CallOption) (*OperationLog, error)
		// 更新操作记录
		UpdateOperationLog(ctx context.Context, in *OperationLog, opts ...grpc.CallOption) (*OperationLog, error)
		// 删除操作记录
		DeleteOperationLog(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除操作记录
		DeleteOperationLogList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询操作记录
		FindOperationLog(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*OperationLog, error)
		// 查询操作记录列表
		FindOperationLogList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*OperationLogPageResp, error)
		// 查询操作记录数量
		FindOperationLogCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
	}

	defaultLogRpc struct {
		cli zrpc.Client
	}
)

func NewLogRpc(cli zrpc.Client) LogRpc {
	return &defaultLogRpc{
		cli: cli,
	}
}

// 创建操作记录
func (m *defaultLogRpc) AddOperationLog(ctx context.Context, in *OperationLog, opts ...grpc.CallOption) (*OperationLog, error) {
	client := blog.NewLogRpcClient(m.cli.Conn())
	return client.AddOperationLog(ctx, in, opts...)
}

// 更新操作记录
func (m *defaultLogRpc) UpdateOperationLog(ctx context.Context, in *OperationLog, opts ...grpc.CallOption) (*OperationLog, error) {
	client := blog.NewLogRpcClient(m.cli.Conn())
	return client.UpdateOperationLog(ctx, in, opts...)
}

// 删除操作记录
func (m *defaultLogRpc) DeleteOperationLog(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewLogRpcClient(m.cli.Conn())
	return client.DeleteOperationLog(ctx, in, opts...)
}

// 批量删除操作记录
func (m *defaultLogRpc) DeleteOperationLogList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewLogRpcClient(m.cli.Conn())
	return client.DeleteOperationLogList(ctx, in, opts...)
}

// 查询操作记录
func (m *defaultLogRpc) FindOperationLog(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*OperationLog, error) {
	client := blog.NewLogRpcClient(m.cli.Conn())
	return client.FindOperationLog(ctx, in, opts...)
}

// 查询操作记录列表
func (m *defaultLogRpc) FindOperationLogList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*OperationLogPageResp, error) {
	client := blog.NewLogRpcClient(m.cli.Conn())
	return client.FindOperationLogList(ctx, in, opts...)
}

// 查询操作记录数量
func (m *defaultLogRpc) FindOperationLogCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewLogRpcClient(m.cli.Conn())
	return client.FindOperationLogCount(ctx, in, opts...)
}
