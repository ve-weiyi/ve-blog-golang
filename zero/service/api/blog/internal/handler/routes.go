// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	album "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/album"
	article "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/article"
	auth "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/auth"
	category "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/category"
	chat "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/chat"
	comment "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/comment"
	friend "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/friend"
	page "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/page"
	remark "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/remark"
	tag "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/tag"
	talk "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/talk"
	upload "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/upload"
	user "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/user"
	website "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/website"
	websocket "github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler/websocket"
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
					// 获取相册列表
					Method:  http.MethodPost,
					Path:    "/album/find_album_list",
					Handler: album.FindAlbumListHandler(serverCtx),
				},
				{
					// 获取相册下的照片列表
					Method:  http.MethodPost,
					Path:    "/album/find_photo_list",
					Handler: album.FindPhotoListHandler(serverCtx),
				},
				{
					// 获取相册
					Method:  http.MethodPost,
					Path:    "/album/get_album",
					Handler: album.GetAlbumHandler(serverCtx),
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
					// 文章归档(时间轴)
					Method:  http.MethodPost,
					Path:    "/article/get_article_archives",
					Handler: article.FindArticleArchivesHandler(serverCtx),
				},
				{
					// 通过分类获取文章列表
					Method:  http.MethodPost,
					Path:    "/article/get_article_classify_category",
					Handler: article.FindArticleClassifyCategoryHandler(serverCtx),
				},
				{
					// 通过标签获取文章列表
					Method:  http.MethodPost,
					Path:    "/article/get_article_classify_tag",
					Handler: article.FindArticleClassifyTagHandler(serverCtx),
				},
				{
					// 获取文章详情
					Method:  http.MethodPost,
					Path:    "/article/get_article_details",
					Handler: article.GetArticleDetailsHandler(serverCtx),
				},
				{
					// 获取首页文章列表
					Method:  http.MethodPost,
					Path:    "/article/get_article_home_list",
					Handler: article.FindArticleHomeListHandler(serverCtx),
				},
				{
					// 获取首页推荐文章列表
					Method:  http.MethodPost,
					Path:    "/article/get_article_recommend",
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
				// 第三方登录授权地址
				Method:  http.MethodPost,
				Path:    "/oauth_authorize_url",
				Handler: auth.OauthAuthorizeUrlHandler(serverCtx),
			},
			{
				// 第三方登录
				Method:  http.MethodPost,
				Path:    "/oauth_login",
				Handler: auth.OauthLoginHandler(serverCtx),
			},
			{
				// 注册
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: auth.RegisterHandler(serverCtx),
			},
			{
				// 发送注册账号邮件
				Method:  http.MethodPost,
				Path:    "/send_register_email",
				Handler: auth.SendRegisterEmailHandler(serverCtx),
			},
			{
				// 重置密码
				Method:  http.MethodPost,
				Path:    "/user/reset_password",
				Handler: auth.ResetPasswordHandler(serverCtx),
			},
			{
				// 发送重置密码邮件
				Method:  http.MethodPost,
				Path:    "/user/send_reset_email",
				Handler: auth.SendResetEmailHandler(serverCtx),
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
					// 查询聊天记录
					Method:  http.MethodPost,
					Path:    "/chat/records",
					Handler: chat.GetChatRecordsHandler(serverCtx),
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
					// 查询最新评论回复列表
					Method:  http.MethodPost,
					Path:    "/comment/find_comment_recent_list",
					Handler: comment.FindCommentRecentListHandler(serverCtx),
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
					Path:    "/comment/add_comment",
					Handler: comment.AddCommentHandler(serverCtx),
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
					// 分页获取友链列表
					Method:  http.MethodPost,
					Path:    "/friend_link/find_friend_list",
					Handler: friend.FindFriendListHandler(serverCtx),
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
					// 分页获取页面列表
					Method:  http.MethodPost,
					Path:    "/page/find_page_list",
					Handler: page.FindPageListHandler(serverCtx),
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
					// 分页获取留言列表
					Method:  http.MethodPost,
					Path:    "/remark/find_remark_list",
					Handler: remark.FindRemarkListHandler(serverCtx),
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
					// 创建留言
					Method:  http.MethodPost,
					Path:    "/remark/add_remark",
					Handler: remark.AddRemarkHandler(serverCtx),
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

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken},
			[]rest.Route{
				{
					// 分页获取说说列表
					Method:  http.MethodPost,
					Path:    "/talk/find_talk_list",
					Handler: talk.FindTalkListHandler(serverCtx),
				},
				{
					// 查询说说
					Method:  http.MethodPost,
					Path:    "/talk/get_talk",
					Handler: talk.GetTalkHandler(serverCtx),
				},
				{
					// 点赞说说
					Method:  http.MethodPut,
					Path:    "/talk/like_talk",
					Handler: talk.LikeTalkHandler(serverCtx),
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
					// 上传文件
					Method:  http.MethodPost,
					Path:    "/upload/upload_file",
					Handler: upload.UploadFileHandler(serverCtx),
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
					// 获取用户信息
					Method:  http.MethodGet,
					Path:    "/user/get_user_info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					// 发送绑定邮箱验证码
					Method:  http.MethodPost,
					Path:    "/user/send_bind_email",
					Handler: user.SendBindEmailHandler(serverCtx),
				},
				{
					// 修改用户信息
					Method:  http.MethodPost,
					Path:    "/user/update_user_avatar",
					Handler: user.UpdateUserAvatarHandler(serverCtx),
				},
				{
					// 修改用户邮箱
					Method:  http.MethodPost,
					Path:    "/user/update_user_email",
					Handler: user.UpdateUserEmailHandler(serverCtx),
				},
				{
					// 修改用户信息
					Method:  http.MethodPost,
					Path:    "/user/update_user_info",
					Handler: user.UpdateUserInfoHandler(serverCtx),
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
					// 获取博客前台首页信息
					Method:  http.MethodGet,
					Path:    "/blog",
					Handler: website.GetBlogHomeInfoHandler(serverCtx),
				},
				{
					// 获取关于我的信息
					Method:  http.MethodGet,
					Path:    "/blog/about_me",
					Handler: website.GetAboutMeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// WebSocket消息
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: websocket.WebSocketHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
