// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package authrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Api                    = blog.Api
	ApiDetails             = blog.ApiDetails
	ApiPageResp            = blog.ApiPageResp
	Article                = blog.Article
	ArticlePageResp        = blog.ArticlePageResp
	BatchResp              = blog.BatchResp
	Category               = blog.Category
	CategoryPageResp       = blog.CategoryPageResp
	ChatRecord             = blog.ChatRecord
	ChatRecordPageResp     = blog.ChatRecordPageResp
	Comment                = blog.Comment
	CommentDetails         = blog.CommentDetails
	CommentDetailsPageResp = blog.CommentDetailsPageResp
	CommentPageResp        = blog.CommentPageResp
	CommentReply           = blog.CommentReply
	CountResp              = blog.CountResp
	EmptyReq               = blog.EmptyReq
	EmptyResp              = blog.EmptyResp
	FindCategoryByNameReq  = blog.FindCategoryByNameReq
	FindConfigReq          = blog.FindConfigReq
	FindConfigResp         = blog.FindConfigResp
	FindTagArticleCountReq = blog.FindTagArticleCountReq
	FindTagByNameReq       = blog.FindTagByNameReq
	FriendLink             = blog.FriendLink
	FriendLinkPageResp     = blog.FriendLinkPageResp
	IdReq                  = blog.IdReq
	IdsReq                 = blog.IdsReq
	LoginHistory           = blog.LoginHistory
	LoginHistoryPageResp   = blog.LoginHistoryPageResp
	LoginReq               = blog.LoginReq
	LoginResp              = blog.LoginResp
	Menu                   = blog.Menu
	MenuDetails            = blog.MenuDetails
	MenuPageResp           = blog.MenuPageResp
	OauthLoginReq          = blog.OauthLoginReq
	OauthLoginUrlResp      = blog.OauthLoginUrlResp
	OperationLog           = blog.OperationLog
	OperationLogPageResp   = blog.OperationLogPageResp
	PageCondition          = blog.PageCondition
	PageLimit              = blog.PageLimit
	PageQuery              = blog.PageQuery
	PageResp               = blog.PageResp
	PageSort               = blog.PageSort
	Photo                  = blog.Photo
	PhotoAlbum             = blog.PhotoAlbum
	PhotoPageResp          = blog.PhotoPageResp
	Remark                 = blog.Remark
	RemarkPageResp         = blog.RemarkPageResp
	ResetPasswordReq       = blog.ResetPasswordReq
	Role                   = blog.Role
	RoleDetails            = blog.RoleDetails
	RoleLabel              = blog.RoleLabel
	RolePageResp           = blog.RolePageResp
	RoleResourcesResp      = blog.RoleResourcesResp
	SaveConfigReq          = blog.SaveConfigReq
	SyncMenuRequest        = blog.SyncMenuRequest
	Tag                    = blog.Tag
	TagPageResp            = blog.TagPageResp
	Talk                   = blog.Talk
	TalkDetailsDTO         = blog.TalkDetailsDTO
	TalkPageResp           = blog.TalkPageResp
	UpdateRoleApisReq      = blog.UpdateRoleApisReq
	UpdateRoleMenusReq     = blog.UpdateRoleMenusReq
	UpdateUserAvatarReq    = blog.UpdateUserAvatarReq
	UpdateUserInfoReq      = blog.UpdateUserInfoReq
	UpdateUserRoleReq      = blog.UpdateUserRoleReq
	UpdateUserStatusReq    = blog.UpdateUserStatusReq
	User                   = blog.User
	UserEmailReq           = blog.UserEmailReq
	UserInfoPageResp       = blog.UserInfoPageResp
	UserInfoResp           = blog.UserInfoResp
	UserReq                = blog.UserReq

	AuthRpc interface {
		// 登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 登出
		Logout(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 注销
		Logoff(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 注册
		Register(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 发送注册邮件
		RegisterEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 发送忘记密码邮件
		ForgetPasswordEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 重置密码
		ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 第三方登录
		OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 获取授权地址
		GetOauthAuthorizeUrl(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthLoginUrlResp, error)
	}

	defaultAuthRpc struct {
		cli zrpc.Client
	}
)

func NewAuthRpc(cli zrpc.Client) AuthRpc {
	return &defaultAuthRpc{
		cli: cli,
	}
}

// 登录
func (m *defaultAuthRpc) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 登出
func (m *defaultAuthRpc) Logout(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}

// 注销
func (m *defaultAuthRpc) Logoff(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.Logoff(ctx, in, opts...)
}

// 注册
func (m *defaultAuthRpc) Register(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 发送注册邮件
func (m *defaultAuthRpc) RegisterEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.RegisterEmail(ctx, in, opts...)
}

// 发送忘记密码邮件
func (m *defaultAuthRpc) ForgetPasswordEmail(ctx context.Context, in *UserEmailReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.ForgetPasswordEmail(ctx, in, opts...)
}

// 重置密码
func (m *defaultAuthRpc) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.ResetPassword(ctx, in, opts...)
}

// 第三方登录
func (m *defaultAuthRpc) OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.OauthLogin(ctx, in, opts...)
}

// 获取授权地址
func (m *defaultAuthRpc) GetOauthAuthorizeUrl(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthLoginUrlResp, error) {
	client := blog.NewAuthRpcClient(m.cli.Conn())
	return client.GetOauthAuthorizeUrl(ctx, in, opts...)
}
