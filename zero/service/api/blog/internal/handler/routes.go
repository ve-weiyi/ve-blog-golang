// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	article "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/article"
	auth "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/auth"
	category "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/category"
	comment "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/comment"
	tag "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/tag"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// ping
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: PingHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken},
			[]rest.Route{
				{
					// 文章归档(时间轴)
					Method:  http.MethodPost,
					Path:    "/article/article_archives",
					Handler: article.FindArticleArchivesHandler(serverCtx),
				},
				{
					// 通过分类获取文章列表
					Method:  http.MethodPost,
					Path:    "/article/article_classify_category",
					Handler: article.FindArticleClassifyCategoryHandler(serverCtx),
				},
				{
					// 通过标签获取文章列表
					Method:  http.MethodPost,
					Path:    "/article/article_classify_tag",
					Handler: article.FindArticleClassifyTagHandler(serverCtx),
				},
				{
					// 获取首页文章列表
					Method:  http.MethodPost,
					Path:    "/article/find_article_home_list",
					Handler: article.FindArticleHomeListHandler(serverCtx),
				},
				{
					// 文章相关推荐
					Method:  http.MethodPost,
					Path:    "/article/find_article_recommend",
					Handler: article.FindArticleRecommendHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken, serverCtx.JwtToken},
			[]rest.Route{
				{
					// 点赞文章
					Method:  http.MethodPost,
					Path:    "/article/like_article",
					Handler: article.LikeArticleHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 登录
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: auth.LoginHandler(serverCtx),
			},
			{
				// 第三方登录
				Method:  http.MethodPost,
				Path:    "/oauth_login",
				Handler: auth.OauthLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtToken},
			[]rest.Route{
				{
					// 注销
					Method:  http.MethodPost,
					Path:    "/logoff",
					Handler: auth.LogoffHandler(serverCtx),
				},
				{
					// 登出
					Method:  http.MethodPost,
					Path:    "/logout",
					Handler: auth.LogoutHandler(serverCtx),
				},
				{
					// 第三方登录授权地址
					Method:  http.MethodPost,
					Path:    "/oauth_authorize_url",
					Handler: auth.OauthAuthorizeUrlHandler(serverCtx),
				},
				{
					// 注册
					Method:  http.MethodPost,
					Path:    "/register",
					Handler: auth.RegisterHandler(serverCtx),
				},
				{
					// 重置密码
					Method:  http.MethodPost,
					Path:    "/reset_password",
					Handler: auth.ResetPasswordHandler(serverCtx),
				},
				{
					// 发送忘记密码邮件
					Method:  http.MethodPost,
					Path:    "/send_forget_email",
					Handler: auth.SendForgetEmailHandler(serverCtx),
				},
				{
					// 发送注册账号邮件
					Method:  http.MethodPost,
					Path:    "/send_register_email",
					Handler: auth.SendRegisterEmailHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken},
			[]rest.Route{
				{
					// 分页获取文章分类列表
					Method:  http.MethodPost,
					Path:    "/category/find_category_list",
					Handler: category.FindCategoryListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken},
			[]rest.Route{
				{
					// 查询评论列表
					Method:  http.MethodPost,
					Path:    "/comment/find_comment_list",
					Handler: comment.FindCommentListHandler(serverCtx),
				},
				{
					// 查询评论回复列表
					Method:  http.MethodPost,
					Path:    "/comment/find_comment_reply_list",
					Handler: comment.FindCommentReplyListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken, serverCtx.JwtToken},
			[]rest.Route{
				{
					// 创建评论
					Method:  http.MethodPost,
					Path:    "/comment/create_comment",
					Handler: comment.CreateCommentHandler(serverCtx),
				},
				{
					// 点赞评论
					Method:  http.MethodPost,
					Path:    "/comment/like_comment",
					Handler: comment.LikeCommentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken},
			[]rest.Route{
				{
					// 分页获取标签列表
					Method:  http.MethodPost,
					Path:    "/tag/find_tag_list",
					Handler: tag.FindTagListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)
}
