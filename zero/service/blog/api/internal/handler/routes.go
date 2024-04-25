// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	account "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/account"
	api "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/api"
	article "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/article"
	auth "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/auth"
	category "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/category"
	friend_link "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/friend_link"
	menu "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/menu"
	mine "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/mine"
	remark "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/remark"
	role "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/role"
	tag "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/tag"
	website "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler/website"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"

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
		[]rest.Route{
			{
				// 查询在线用户列表
				Method:  http.MethodPost,
				Path:    "/user/find_online_user_list",
				Handler: account.FindOnlineUserListHandler(serverCtx),
			},
			{
				// 获取用户分布地区
				Method:  http.MethodPost,
				Path:    "/user/find_user_areas",
				Handler: account.FindUserAreasHandler(serverCtx),
			},
			{
				// 查询用户列表
				Method:  http.MethodPost,
				Path:    "/user/find_user_list",
				Handler: account.FindUserListHandler(serverCtx),
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
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 清空接口列表
				Method:  http.MethodPost,
				Path:    "/api/clean_api_list",
				Handler: api.CleanApiListHandler(serverCtx),
			},
			{
				// 创建api路由
				Method:  http.MethodPost,
				Path:    "/api/create_api",
				Handler: api.CreateApiHandler(serverCtx),
			},
			{
				// 删除api路由
				Method:  http.MethodDelete,
				Path:    "/api/delete_api",
				Handler: api.DeleteApiHandler(serverCtx),
			},
			{
				// 批量删除api路由
				Method:  http.MethodDelete,
				Path:    "/api/delete_api_list",
				Handler: api.DeleteApiListHandler(serverCtx),
			},
			{
				// 查询api路由
				Method:  http.MethodPost,
				Path:    "/api/find_api",
				Handler: api.FindApiHandler(serverCtx),
			},
			{
				// 分页获取api路由列表
				Method:  http.MethodPost,
				Path:    "/api/find_api_list",
				Handler: api.FindApiListHandler(serverCtx),
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
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 删除文章-物理删除
				Method:  http.MethodPost,
				Path:    "/admin/article/delete_article",
				Handler: article.DeleteArticleHandler(serverCtx),
			},
			{
				// 查询文章
				Method:  http.MethodPost,
				Path:    "/admin/article/find_article",
				Handler: article.FindArticleHandler(serverCtx),
			},
			{
				// 分页获取文章列表
				Method:  http.MethodPost,
				Path:    "/admin/article/find_article_list",
				Handler: article.FindArticleListHandler(serverCtx),
			},
			{
				// 删除文章-逻辑删除
				Method:  http.MethodPost,
				Path:    "/admin/article/pre_delete_article",
				Handler: article.PreDeleteArticleHandler(serverCtx),
			},
			{
				// 保存文章
				Method:  http.MethodPost,
				Path:    "/admin/article/save_article",
				Handler: article.SaveArticleHandler(serverCtx),
			},
			{
				// 置顶文章
				Method:  http.MethodPost,
				Path:    "/admin/article/top_article",
				Handler: article.TopArticleHandler(serverCtx),
			},
			{
				// 文章归档(时间轴)
				Method:  http.MethodPost,
				Path:    "/article/article_archives",
				Handler: article.FindArticleArchivesHandler(serverCtx),
			},
			{
				// 通过标签或者id获取文章列表
				Method:  http.MethodPost,
				Path:    "/article/article_classify_category",
				Handler: article.FindArticleClassifyCategoryHandler(serverCtx),
			},
			{
				// 通过标签或者id获取文章列表
				Method:  http.MethodPost,
				Path:    "/article/article_classify_tag",
				Handler: article.FindArticleClassifyTagHandler(serverCtx),
			},
			{
				// 分页获取文章列表
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
			{
				// 点赞文章
				Method:  http.MethodPost,
				Path:    "/article/like_article",
				Handler: article.LikeArticleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 发送忘记密码邮件
				Method:  http.MethodPost,
				Path:    "/forget/email",
				Handler: auth.ForgetPasswordEmailHandler(serverCtx),
			},
			{
				// 重置密码
				Method:  http.MethodPost,
				Path:    "/forget/reset_password",
				Handler: auth.ResetPasswordHandler(serverCtx),
			},
			{
				// 登录
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: auth.LoginHandler(serverCtx),
			},
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
				// 获取授权地址
				Method:  http.MethodPost,
				Path:    "/oauth/authorize_url",
				Handler: auth.GetOauthAuthorizeUrlHandler(serverCtx),
			},
			{
				// 第三方登录
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: auth.OauthLoginHandler(serverCtx),
			},
			{
				// 注册
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: auth.RegisterHandler(serverCtx),
			},
			{
				// 发送注册邮件
				Method:  http.MethodPost,
				Path:    "/register/email",
				Handler: auth.RegisterEmailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建文章分类
				Method:  http.MethodPost,
				Path:    "/category/create_category",
				Handler: category.CreateCategoryHandler(serverCtx),
			},
			{
				// 删除文章分类
				Method:  http.MethodDelete,
				Path:    "/category/delete_category",
				Handler: category.DeleteCategoryHandler(serverCtx),
			},
			{
				// 批量删除文章分类
				Method:  http.MethodDelete,
				Path:    "/category/delete_category_list",
				Handler: category.DeleteCategoryListHandler(serverCtx),
			},
			{
				// 查询文章分类
				Method:  http.MethodPost,
				Path:    "/category/find_category",
				Handler: category.FindCategoryHandler(serverCtx),
			},
			{
				// 分页获取文章分类列表
				Method:  http.MethodPost,
				Path:    "/category/find_category_list",
				Handler: category.FindCategoryListHandler(serverCtx),
			},
			{
				// 更新文章分类
				Method:  http.MethodPut,
				Path:    "/category/update_category",
				Handler: category.UpdateCategoryHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建友链
				Method:  http.MethodPost,
				Path:    "/friend_link/create_friend_link",
				Handler: friend_link.CreateFriendLinkHandler(serverCtx),
			},
			{
				// 删除友链
				Method:  http.MethodDelete,
				Path:    "/friend_link/delete_friend_link",
				Handler: friend_link.DeleteFriendLinkHandler(serverCtx),
			},
			{
				// 批量删除友链
				Method:  http.MethodDelete,
				Path:    "/friend_link/delete_friend_link_list",
				Handler: friend_link.DeleteFriendLinkListHandler(serverCtx),
			},
			{
				// 查询友链
				Method:  http.MethodPost,
				Path:    "/friend_link/find_friend_link",
				Handler: friend_link.FindFriendLinkHandler(serverCtx),
			},
			{
				// 分页获取友链列表
				Method:  http.MethodPost,
				Path:    "/friend_link/find_friend_link_list",
				Handler: friend_link.FindFriendLinkListHandler(serverCtx),
			},
			{
				// 更新友链
				Method:  http.MethodPut,
				Path:    "/friend_link/update_friend_link",
				Handler: friend_link.UpdateFriendLinkHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 清空菜单列表
				Method:  http.MethodPost,
				Path:    "/menu/clean_menu_list",
				Handler: menu.CleanMenuListHandler(serverCtx),
			},
			{
				// 创建菜单
				Method:  http.MethodPost,
				Path:    "/menu/create_menu",
				Handler: menu.CreateMenuHandler(serverCtx),
			},
			{
				// 删除菜单
				Method:  http.MethodDelete,
				Path:    "/menu/delete_menu",
				Handler: menu.DeleteMenuHandler(serverCtx),
			},
			{
				// 批量删除菜单
				Method:  http.MethodDelete,
				Path:    "/menu/delete_menu_list",
				Handler: menu.DeleteMenuListHandler(serverCtx),
			},
			{
				// 查询菜单
				Method:  http.MethodPost,
				Path:    "/menu/find_menu",
				Handler: menu.FindMenuHandler(serverCtx),
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
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取用户接口权限
				Method:  http.MethodGet,
				Path:    "/user/apis",
				Handler: mine.GetUserApisHandler(serverCtx),
			},
			{
				// 更换用户头像
				Method:  http.MethodPost,
				Path:    "/user/avatar",
				Handler: mine.UpdateUserAvatarHandler(serverCtx),
			},
			{
				// 批量删除登录历史
				Method:  http.MethodDelete,
				Path:    "/user/delete_login_history_list",
				Handler: mine.DeleteUserLoginHistoryListHandler(serverCtx),
			},
			{
				// 获取用户信息
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: mine.GetUserInfoHandler(serverCtx),
			},
			{
				// 修改用户信息
				Method:  http.MethodPost,
				Path:    "/user/info",
				Handler: mine.UpdateUserInfoHandler(serverCtx),
			},
			{
				// 查询用户登录历史
				Method:  http.MethodPost,
				Path:    "/user/login_history",
				Handler: mine.FindUserLoginHistoryListHandler(serverCtx),
			},
			{
				// 获取用户菜单权限
				Method:  http.MethodGet,
				Path:    "/user/menus",
				Handler: mine.GetUserMenusHandler(serverCtx),
			},
			{
				// 获取用户角色
				Method:  http.MethodGet,
				Path:    "/user/roles",
				Handler: mine.GetUserRoleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建留言
				Method:  http.MethodPost,
				Path:    "/remark/create_remark",
				Handler: remark.CreateRemarkHandler(serverCtx),
			},
			{
				// 删除留言
				Method:  http.MethodDelete,
				Path:    "/remark/delete_remark",
				Handler: remark.DeleteRemarkHandler(serverCtx),
			},
			{
				// 批量删除留言
				Method:  http.MethodDelete,
				Path:    "/remark/delete_remark_list",
				Handler: remark.DeleteRemarkListHandler(serverCtx),
			},
			{
				// 查询留言
				Method:  http.MethodPost,
				Path:    "/remark/find_remark",
				Handler: remark.FindRemarkHandler(serverCtx),
			},
			{
				// 分页获取留言列表
				Method:  http.MethodPost,
				Path:    "/remark/find_remark_list",
				Handler: remark.FindRemarkListHandler(serverCtx),
			},
			{
				// 更新留言
				Method:  http.MethodPut,
				Path:    "/remark/update_remark",
				Handler: remark.UpdateRemarkHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建角色
				Method:  http.MethodPost,
				Path:    "/role/create_role",
				Handler: role.CreateRoleHandler(serverCtx),
			},
			{
				// 删除角色
				Method:  http.MethodDelete,
				Path:    "/role/delete_role",
				Handler: role.DeleteRoleHandler(serverCtx),
			},
			{
				// 批量删除角色
				Method:  http.MethodDelete,
				Path:    "/role/delete_role_list",
				Handler: role.DeleteRoleListHandler(serverCtx),
			},
			{
				// 查询角色
				Method:  http.MethodPost,
				Path:    "/role/find_role",
				Handler: role.FindRoleHandler(serverCtx),
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
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建标签
				Method:  http.MethodPost,
				Path:    "/tag/create_tag",
				Handler: tag.CreateTagHandler(serverCtx),
			},
			{
				// 删除标签
				Method:  http.MethodDelete,
				Path:    "/tag/delete_tag",
				Handler: tag.DeleteTagHandler(serverCtx),
			},
			{
				// 批量删除标签
				Method:  http.MethodDelete,
				Path:    "/tag/delete_tag_list",
				Handler: tag.DeleteTagListHandler(serverCtx),
			},
			{
				// 查询标签
				Method:  http.MethodPost,
				Path:    "/tag/find_tag",
				Handler: tag.FindTagHandler(serverCtx),
			},
			{
				// 分页获取标签列表
				Method:  http.MethodPost,
				Path:    "/tag/find_tag_list",
				Handler: tag.FindTagListHandler(serverCtx),
			},
			{
				// 更新标签
				Method:  http.MethodPut,
				Path:    "/tag/update_tag",
				Handler: tag.UpdateTagHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取后台首页信息
				Method:  http.MethodGet,
				Path:    "/admin",
				Handler: website.GetAdminHomeInfoHandler(serverCtx),
			},
			{
				// 更新关于我的信息
				Method:  http.MethodPut,
				Path:    "/admin/about_me",
				Handler: website.UpdateAboutMeHandler(serverCtx),
			},
			{
				// 获取服务器信息
				Method:  http.MethodGet,
				Path:    "/admin/system_state",
				Handler: website.GetSystemStateHandler(serverCtx),
			},
			{
				// 更新配置
				Method:  http.MethodPut,
				Path:    "/admin/websit_config",
				Handler: website.UpdateWebsiteConfigHandler(serverCtx),
			},
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
			{
				// 获取网站前台配置
				Method:  http.MethodGet,
				Path:    "/blog/websit_config",
				Handler: website.GetWebsiteConfigHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
