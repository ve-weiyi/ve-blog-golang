// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package friendlinkrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AnalysisArticleResp       = blog.AnalysisArticleResp
	Api                       = blog.Api
	ApiDetails                = blog.ApiDetails
	ArticleCategory           = blog.ArticleCategory
	ArticleDetails            = blog.ArticleDetails
	ArticleNew                = blog.ArticleNew
	ArticleTag                = blog.ArticleTag
	BatchResp                 = blog.BatchResp
	BindUserEmailReq          = blog.BindUserEmailReq
	CategoryDetails           = blog.CategoryDetails
	CategoryNew               = blog.CategoryNew
	ChatRecord                = blog.ChatRecord
	Comment                   = blog.Comment
	CommentUserInfo           = blog.CommentUserInfo
	CountResp                 = blog.CountResp
	EmptyReq                  = blog.EmptyReq
	EmptyResp                 = blog.EmptyResp
	FindApiListResp           = blog.FindApiListResp
	FindArticleListReq        = blog.FindArticleListReq
	FindArticleListResp       = blog.FindArticleListResp
	FindArticlesByCategoryReq = blog.FindArticlesByCategoryReq
	FindArticlesByTagReq      = blog.FindArticlesByTagReq
	FindCategoryListReq       = blog.FindCategoryListReq
	FindCategoryListResp      = blog.FindCategoryListResp
	FindChatRecordListResp    = blog.FindChatRecordListResp
	FindCommentListResp       = blog.FindCommentListResp
	FindConfigReq             = blog.FindConfigReq
	FindConfigResp            = blog.FindConfigResp
	FindFriendLinkListResp    = blog.FindFriendLinkListResp
	FindLikeArticleResp       = blog.FindLikeArticleResp
	FindLikeCommentResp       = blog.FindLikeCommentResp
	FindLikeTalkResp          = blog.FindLikeTalkResp
	FindLoginHistoryListReq   = blog.FindLoginHistoryListReq
	FindLoginHistoryListResp  = blog.FindLoginHistoryListResp
	FindMenuListResp          = blog.FindMenuListResp
	FindOperationLogListResp  = blog.FindOperationLogListResp
	FindPageListResp          = blog.FindPageListResp
	FindPhotoAlbumListResp    = blog.FindPhotoAlbumListResp
	FindPhotoListResp         = blog.FindPhotoListResp
	FindRemarkListResp        = blog.FindRemarkListResp
	FindRoleListResp          = blog.FindRoleListResp
	FindTagArticleCountReq    = blog.FindTagArticleCountReq
	FindTagByNameReq          = blog.FindTagByNameReq
	FindTagListReq            = blog.FindTagListReq
	FindTagListResp           = blog.FindTagListResp
	FindTalkListResp          = blog.FindTalkListResp
	FindUserListReq           = blog.FindUserListReq
	FindUserListResp          = blog.FindUserListResp
	FindUserRegionListResp    = blog.FindUserRegionListResp
	FriendLink                = blog.FriendLink
	GetLogoutAtReq            = blog.GetLogoutAtReq
	IdReq                     = blog.IdReq
	IdsReq                    = blog.IdsReq
	ListResp                  = blog.ListResp
	LoginReq                  = blog.LoginReq
	LoginResp                 = blog.LoginResp
	LogoffReq                 = blog.LogoffReq
	LogoutReq                 = blog.LogoutReq
	LogoutResp                = blog.LogoutResp
	Menu                      = blog.Menu
	MenuDetails               = blog.MenuDetails
	OauthLoginReq             = blog.OauthLoginReq
	OauthLoginUrlResp         = blog.OauthLoginUrlResp
	OperationLog              = blog.OperationLog
	Page                      = blog.Page
	PageCondition             = blog.PageCondition
	PageLimit                 = blog.PageLimit
	PageQuery                 = blog.PageQuery
	PageSort                  = blog.PageSort
	Photo                     = blog.Photo
	PhotoAlbum                = blog.PhotoAlbum
	RecycleArticleReq         = blog.RecycleArticleReq
	RegisterReq               = blog.RegisterReq
	Remark                    = blog.Remark
	ResetPasswordReq          = blog.ResetPasswordReq
	Role                      = blog.Role
	RoleDetails               = blog.RoleDetails
	RoleResourcesResp         = blog.RoleResourcesResp
	SaveConfigReq             = blog.SaveConfigReq
	SyncMenuReq               = blog.SyncMenuReq
	TagDetails                = blog.TagDetails
	TagMapResp                = blog.TagMapResp
	TagNew                    = blog.TagNew
	Talk                      = blog.Talk
	TalkDetailsDTO            = blog.TalkDetailsDTO
	TopArticleReq             = blog.TopArticleReq
	UpdateRoleApisReq         = blog.UpdateRoleApisReq
	UpdateRoleMenusReq        = blog.UpdateRoleMenusReq
	UpdateUserInfoReq         = blog.UpdateUserInfoReq
	UpdateUserRoleReq         = blog.UpdateUserRoleReq
	UpdateUserStatusReq       = blog.UpdateUserStatusReq
	UploadRecordReq           = blog.UploadRecordReq
	UploadRecordResp          = blog.UploadRecordResp
	UserDetails               = blog.UserDetails
	UserEmailReq              = blog.UserEmailReq
	UserIdReq                 = blog.UserIdReq
	UserInfoResp              = blog.UserInfoResp
	UserLoginHistory          = blog.UserLoginHistory
	UserRegion                = blog.UserRegion
	UserRoleLabel             = blog.UserRoleLabel
	UserVisit                 = blog.UserVisit
	UserVisitPageRsp          = blog.UserVisitPageRsp

	FriendLinkRpc interface {
		// 创建友链
		AddFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error)
		// 更新友链
		UpdateFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error)
		// 删除友链
		DeleteFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除友链
		DeleteFriendLinkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询友链
		FindFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*FriendLink, error)
		// 查询友链列表
		FindFriendLinkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*FindFriendLinkListResp, error)
		// 查询友链数量
		FindFriendLinkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error)
	}

	defaultFriendLinkRpc struct {
		cli zrpc.Client
	}
)

func NewFriendLinkRpc(cli zrpc.Client) FriendLinkRpc {
	return &defaultFriendLinkRpc{
		cli: cli,
	}
}

// 创建友链
func (m *defaultFriendLinkRpc) AddFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.AddFriendLink(ctx, in, opts...)
}

// 更新友链
func (m *defaultFriendLinkRpc) UpdateFriendLink(ctx context.Context, in *FriendLink, opts ...grpc.CallOption) (*FriendLink, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.UpdateFriendLink(ctx, in, opts...)
}

// 删除友链
func (m *defaultFriendLinkRpc) DeleteFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.DeleteFriendLink(ctx, in, opts...)
}

// 批量删除友链
func (m *defaultFriendLinkRpc) DeleteFriendLinkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.DeleteFriendLinkList(ctx, in, opts...)
}

// 查询友链
func (m *defaultFriendLinkRpc) FindFriendLink(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*FriendLink, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.FindFriendLink(ctx, in, opts...)
}

// 查询友链列表
func (m *defaultFriendLinkRpc) FindFriendLinkList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*FindFriendLinkListResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.FindFriendLinkList(ctx, in, opts...)
}

// 查询友链数量
func (m *defaultFriendLinkRpc) FindFriendLinkCount(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*CountResp, error) {
	client := blog.NewFriendLinkRpcClient(m.cli.Conn())
	return client.FindFriendLinkCount(ctx, in, opts...)
}
