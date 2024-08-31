// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	account "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/account"
	album "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/album"
	api "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/api"
	article "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/article"
	auth "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/auth"
	banner "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/banner"
	category "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/category"
	comment "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/comment"
	friend "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/friend"
	menu "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/menu"
	operation_log "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/operation_log"
	photo "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/photo"
	remark "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/remark"
	role "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/role"
	tag "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/tag"
	talk "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/talk"
	upload "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/upload"
	website "github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/handler/website"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"

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
			[]rest.Middleware{serverCtx.SignToken, serverCtx.JwtToken},
			[]rest.Route{
				{
					// 查询在线用户列表
					Method:  http.MethodPost,
					Path:    "/user/find_online_user_list",
					Handler: account.FindOnlineUserListHandler(serverCtx),
				},
				{
					// 查询用户列表
					Method:  http.MethodPost,
					Path:    "/user/find_user_list",
					Handler: account.FindUserListHandler(serverCtx),
				},
				{
					// 查询用户登录历史
					Method:  http.MethodPost,
					Path:    "/user/find_user_login_history_list",
					Handler: account.FindUserLoginHistoryListHandler(serverCtx),
				},
				{
					// 获取用户接口权限
					Method:  http.MethodGet,
					Path:    "/user/get_user_apis",
					Handler: account.GetUserApisHandler(serverCtx),
				},
				{
					// 获取用户分布地区
					Method:  http.MethodPost,
					Path:    "/user/get_user_area_analysis",
					Handler: account.GetUserAreaAnalysisHandler(serverCtx),
				},
				{
					// 获取用户信息
					Method:  http.MethodGet,
					Path:    "/user/get_user_info",
					Handler: account.GetUserInfoHandler(serverCtx),
				},
				{
					// 获取用户菜单权限
					Method:  http.MethodGet,
					Path:    "/user/get_user_menus",
					Handler: account.GetUserMenusHandler(serverCtx),
				},
				{
					// 获取用户角色
					Method:  http.MethodGet,
					Path:    "/user/get_user_roles",
					Handler: account.GetUserRolesHandler(serverCtx),
				},
				{
					// 修改用户信息
					Method:  http.MethodPost,
					Path:    "/user/update_user_info",
					Handler: account.UpdateUserInfoHandler(serverCtx),
				},
				{
					// 修改用户角色
					Method:  http.MethodPost,
					Path:    "/user/update_user_roles",
					Handler: account.UpdateUserRolesHandler(serverCtx),
				},
				{
					// 修改用户状态
					Method:  http.MethodPost,
					Path:    "/user/update_user_status",
					Handler: account.UpdateUserStatusHandler(serverCtx),
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
					// 分页获取相册列表
					Method:  http.MethodPost,
					Path:    "/album/find_album_list",
					Handler: album.FindAlbumListHandler(serverCtx),
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
					// 创建相册
					Method:  http.MethodPost,
					Path:    "/album/add_album",
					Handler: album.AddAlbumHandler(serverCtx),
				},
				{
					// 批量删除相册
					Method:  http.MethodDelete,
					Path:    "/album/batch_delete_album",
					Handler: album.BatchDeleteAlbumHandler(serverCtx),
				},
				{
					// 删除相册
					Method:  http.MethodDelete,
					Path:    "/album/delete_album",
					Handler: album.DeleteAlbumHandler(serverCtx),
				},
				{
					// 查询相册
					Method:  http.MethodPost,
					Path:    "/album/get_album",
					Handler: album.GetAlbumHandler(serverCtx),
				},
				{
					// 更新相册
					Method:  http.MethodPut,
					Path:    "/album/update_album",
					Handler: album.UpdateAlbumHandler(serverCtx),
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
					// 分页获取api路由列表
					Method:  http.MethodPost,
					Path:    "/api/find_api_list",
					Handler: api.FindApiListHandler(serverCtx),
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
					// 创建api路由
					Method:  http.MethodPost,
					Path:    "/api/add_api",
					Handler: api.AddApiHandler(serverCtx),
				},
				{
					// 批量删除api路由
					Method:  http.MethodDelete,
					Path:    "/api/batch_delete_api",
					Handler: api.BatchDeleteApiHandler(serverCtx),
				},
				{
					// 清空接口列表
					Method:  http.MethodPost,
					Path:    "/api/clean_api_list",
					Handler: api.CleanApiListHandler(serverCtx),
				},
				{
					// 删除api路由
					Method:  http.MethodDelete,
					Path:    "/api/delete_api",
					Handler: api.DeleteApiHandler(serverCtx),
				},
				{
					// 同步api列表
					Method:  http.MethodPost,
					Path:    "/api/sync_api_list",
					Handler: api.SyncApiListHandler(serverCtx),
				},
				{
					// 更新api路由
					Method:  http.MethodPut,
					Path:    "/api/update_api",
					Handler: api.UpdateApiHandler(serverCtx),
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
					// 添加文章
					Method:  http.MethodPost,
					Path:    "/admin/article/add_article",
					Handler: article.AddArticleHandler(serverCtx),
				},
				{
					// 删除文章
					Method:  http.MethodPost,
					Path:    "/admin/article/delete_article",
					Handler: article.DeleteArticleHandler(serverCtx),
				},
				{
					// 导出文章列表
					Method:  http.MethodPost,
					Path:    "/admin/article/export_article_list",
					Handler: article.ExportArticleListHandler(serverCtx),
				},
				{
					// 查询文章列表
					Method:  http.MethodPost,
					Path:    "/admin/article/find_article_list",
					Handler: article.FindArticleListHandler(serverCtx),
				},
				{
					// 查询文章
					Method:  http.MethodPost,
					Path:    "/admin/article/get_article",
					Handler: article.GetArticleHandler(serverCtx),
				},
				{
					// 回收文章
					Method:  http.MethodPost,
					Path:    "/admin/article/recycle_article",
					Handler: article.RecycleArticleHandler(serverCtx),
				},
				{
					// 置顶文章
					Method:  http.MethodPost,
					Path:    "/admin/article/top_article",
					Handler: article.TopArticleHandler(serverCtx),
				},
				{
					// 保存文章
					Method:  http.MethodPost,
					Path:    "/admin/article/update_article",
					Handler: article.UpdateArticleHandler(serverCtx),
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
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtToken},
			[]rest.Route{
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
			[]rest.Middleware{serverCtx.SignToken, serverCtx.JwtToken},
			[]rest.Route{
				{
					// 创建页面
					Method:  http.MethodPost,
					Path:    "/banner/add_banner",
					Handler: banner.AddBannerHandler(serverCtx),
				},
				{
					// 删除页面
					Method:  http.MethodDelete,
					Path:    "/banner/delete_banner",
					Handler: banner.DeleteBannerHandler(serverCtx),
				},
				{
					// 分页获取页面列表
					Method:  http.MethodPost,
					Path:    "/banner/find_banner_list",
					Handler: banner.FindBannerListHandler(serverCtx),
				},
				{
					// 更新页面
					Method:  http.MethodPut,
					Path:    "/banner/update_banner",
					Handler: banner.UpdateBannerHandler(serverCtx),
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
			[]rest.Middleware{serverCtx.SignToken, serverCtx.JwtToken},
			[]rest.Route{
				{
					// 创建文章分类
					Method:  http.MethodPost,
					Path:    "/category/add_category",
					Handler: category.AddCategoryHandler(serverCtx),
				},
				{
					// 批量删除文章分类
					Method:  http.MethodDelete,
					Path:    "/category/batch_delete_category",
					Handler: category.BatchDeleteCategoryHandler(serverCtx),
				},
				{
					// 删除文章分类
					Method:  http.MethodDelete,
					Path:    "/category/delete_category",
					Handler: category.DeleteCategoryHandler(serverCtx),
				},
				{
					// 更新文章分类
					Method:  http.MethodPut,
					Path:    "/category/update_category",
					Handler: category.UpdateCategoryHandler(serverCtx),
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
					// 批量删除评论
					Method:  http.MethodDelete,
					Path:    "/comment/batch_delete_comment",
					Handler: comment.BatchDeleteCommentHandler(serverCtx),
				},
				{
					// 删除评论
					Method:  http.MethodDelete,
					Path:    "/comment/delete_comment",
					Handler: comment.DeleteCommentHandler(serverCtx),
				},
				{
					// 查询评论列表(后台)
					Method:  http.MethodPost,
					Path:    "/comment/find_comment_back_list",
					Handler: comment.FindCommentBackListHandler(serverCtx),
				},
				{
					// 更新评论审核状态
					Method:  http.MethodPut,
					Path:    "/comment/update_comment_review",
					Handler: comment.UpdateCommentReviewHandler(serverCtx),
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
					Path:    "/friend/find_friend_list",
					Handler: friend.FindFriendListHandler(serverCtx),
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
					// 创建友链
					Method:  http.MethodPost,
					Path:    "/friend/add_friend",
					Handler: friend.AddFriendHandler(serverCtx),
				},
				{
					// 批量删除友链
					Method:  http.MethodDelete,
					Path:    "/friend/batch_delete_friend",
					Handler: friend.BatchDeleteFriendHandler(serverCtx),
				},
				{
					// 删除友链
					Method:  http.MethodDelete,
					Path:    "/friend/delete_friend",
					Handler: friend.DeleteFriendHandler(serverCtx),
				},
				{
					// 更新友链
					Method:  http.MethodPut,
					Path:    "/friend/update_friend",
					Handler: friend.UpdateFriendHandler(serverCtx),
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
					// 创建菜单
					Method:  http.MethodPost,
					Path:    "/menu/add_menu",
					Handler: menu.AddMenuHandler(serverCtx),
				},
				{
					// 批量删除菜单
					Method:  http.MethodDelete,
					Path:    "/menu/batch_delete_menu",
					Handler: menu.BatchDeleteMenuHandler(serverCtx),
				},
				{
					// 清空菜单列表
					Method:  http.MethodPost,
					Path:    "/menu/clean_menu_list",
					Handler: menu.CleanMenuListHandler(serverCtx),
				},
				{
					// 删除菜单
					Method:  http.MethodDelete,
					Path:    "/menu/delete_menu",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
				{
					// 分页获取菜单列表
					Method:  http.MethodPost,
					Path:    "/menu/find_menu_list",
					Handler: menu.FindMenuListHandler(serverCtx),
				},
				{
					// 同步菜单列表
					Method:  http.MethodPost,
					Path:    "/menu/sync_menu_list",
					Handler: menu.SyncMenuListHandler(serverCtx),
				},
				{
					// 更新菜单
					Method:  http.MethodPut,
					Path:    "/menu/update_menu",
					Handler: menu.UpdateMenuHandler(serverCtx),
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
					// 批量删除操作记录
					Method:  http.MethodDelete,
					Path:    "/operation_log/batch_delete_operation_log",
					Handler: operation_log.BatchDeleteOperationLogHandler(serverCtx),
				},
				{
					// 删除操作记录
					Method:  http.MethodDelete,
					Path:    "/operation_log/delete_operation_log",
					Handler: operation_log.DeleteOperationLogHandler(serverCtx),
				},
				{
					// 分页获取操作记录列表
					Method:  http.MethodPost,
					Path:    "/operation_log/find_operation_log_list",
					Handler: operation_log.FindOperationLogListHandler(serverCtx),
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
					// 分页获取照片列表
					Method:  http.MethodPost,
					Path:    "/photo/find_photo_list",
					Handler: photo.FindPhotoListHandler(serverCtx),
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
					// 批量删除照片
					Method:  http.MethodDelete,
					Path:    "/album/batch_delete_photo",
					Handler: photo.BatchDeletePhotoHandler(serverCtx),
				},
				{
					// 创建照片
					Method:  http.MethodPost,
					Path:    "/photo/add_photo",
					Handler: photo.AddPhotoHandler(serverCtx),
				},
				{
					// 删除照片
					Method:  http.MethodDelete,
					Path:    "/photo/delete_photo",
					Handler: photo.DeletePhotoHandler(serverCtx),
				},
				{
					// 更新照片
					Method:  http.MethodPut,
					Path:    "/photo/update_photo",
					Handler: photo.UpdatePhotoHandler(serverCtx),
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
					// 批量删除留言
					Method:  http.MethodDelete,
					Path:    "/remark/batch_delete_remark",
					Handler: remark.BatchDeleteRemarkHandler(serverCtx),
				},
				{
					// 删除留言
					Method:  http.MethodDelete,
					Path:    "/remark/delete_remark",
					Handler: remark.DeleteRemarkHandler(serverCtx),
				},
				{
					// 更新留言
					Method:  http.MethodPut,
					Path:    "/remark/update_remark",
					Handler: remark.UpdateRemarkHandler(serverCtx),
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
					// 创建角色
					Method:  http.MethodPost,
					Path:    "/role/add_role",
					Handler: role.AddRoleHandler(serverCtx),
				},
				{
					// 批量删除角色
					Method:  http.MethodPost,
					Path:    "/role/batch_delete_role",
					Handler: role.BatchDeleteRoleHandler(serverCtx),
				},
				{
					// 删除角色
					Method:  http.MethodDelete,
					Path:    "/role/delete_role",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					// 分页获取角色列表
					Method:  http.MethodPost,
					Path:    "/role/find_role_list",
					Handler: role.FindRoleListHandler(serverCtx),
				},
				{
					// 获取角色资源列表
					Method:  http.MethodPost,
					Path:    "/role/find_role_resources",
					Handler: role.FindRoleResourcesHandler(serverCtx),
				},
				{
					// 更新角色
					Method:  http.MethodPut,
					Path:    "/role/update_role",
					Handler: role.UpdateRoleHandler(serverCtx),
				},
				{
					// 更新角色接口权限
					Method:  http.MethodPost,
					Path:    "/role/update_role_apis",
					Handler: role.UpdateRoleApisHandler(serverCtx),
				},
				{
					// 更新角色菜单权限
					Method:  http.MethodPost,
					Path:    "/role/update_role_menus",
					Handler: role.UpdateRoleMenusHandler(serverCtx),
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
			[]rest.Middleware{serverCtx.SignToken, serverCtx.JwtToken},
			[]rest.Route{
				{
					// 创建标签
					Method:  http.MethodPost,
					Path:    "/tag/add_tag",
					Handler: tag.AddTagHandler(serverCtx),
				},
				{
					// 批量删除标签
					Method:  http.MethodDelete,
					Path:    "/tag/batch_delete_tag",
					Handler: tag.BatchDeleteTagHandler(serverCtx),
				},
				{
					// 删除标签
					Method:  http.MethodDelete,
					Path:    "/tag/delete_tag",
					Handler: tag.DeleteTagHandler(serverCtx),
				},
				{
					// 更新标签
					Method:  http.MethodPut,
					Path:    "/tag/update_tag",
					Handler: tag.UpdateTagHandler(serverCtx),
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
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SignToken, serverCtx.JwtToken},
			[]rest.Route{
				{
					// 创建说说
					Method:  http.MethodPost,
					Path:    "/talk/add_talk",
					Handler: talk.AddTalkHandler(serverCtx),
				},
				{
					// 删除说说
					Method:  http.MethodDelete,
					Path:    "/talk/delete_talk",
					Handler: talk.DeleteTalkHandler(serverCtx),
				},
				{
					// 查询说说
					Method:  http.MethodPost,
					Path:    "/talk/get_talk",
					Handler: talk.GetTalkHandler(serverCtx),
				},
				{
					// 更新说说
					Method:  http.MethodPut,
					Path:    "/talk/update_talk",
					Handler: talk.UpdateTalkHandler(serverCtx),
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
					// 获取后台首页信息
					Method:  http.MethodGet,
					Path:    "/admin",
					Handler: website.GetAdminHomeInfoHandler(serverCtx),
				},
				{
					// 获取关于我的信息
					Method:  http.MethodGet,
					Path:    "/admin/about_me",
					Handler: website.GetAboutMeHandler(serverCtx),
				},
				{
					// 更新关于我的信息
					Method:  http.MethodPut,
					Path:    "/admin/about_me",
					Handler: website.UpdateAboutMeHandler(serverCtx),
				},
				{
					// 获取网站配置
					Method:  http.MethodGet,
					Path:    "/admin/get_website_config",
					Handler: website.GetWebsiteConfigHandler(serverCtx),
				},
				{
					// 获取服务器信息
					Method:  http.MethodGet,
					Path:    "/admin/system_state",
					Handler: website.GetSystemStateHandler(serverCtx),
				},
				{
					// 更新网站配置
					Method:  http.MethodPut,
					Path:    "/admin/update_website_config",
					Handler: website.UpdateWebsiteConfigHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)
}
