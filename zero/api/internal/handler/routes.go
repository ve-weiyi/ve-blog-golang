// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "github.com/ve-weiyi/ve-blog-golang/zero/api/internal/handler/api"
	auth "github.com/ve-weiyi/ve-blog-golang/zero/api/internal/handler/auth"
	menu "github.com/ve-weiyi/ve-blog-golang/zero/api/internal/handler/menu"
	role "github.com/ve-weiyi/ve-blog-golang/zero/api/internal/handler/role"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
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
				Method:  http.MethodPost,
				Path:    "/api/create_api",
				Handler: api.CreateApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/update_api",
				Handler: api.UpdateApiHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/delete_api",
				Handler: api.DeleteApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/find_api",
				Handler: api.FindApiHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/delete_api_list",
				Handler: api.DeleteApiListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/find_api_list",
				Handler: api.FindApiListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sync_api_list",
				Handler: api.SyncApiListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/clean_api_list",
				Handler: api.CleanApiListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: auth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logout",
				Handler: auth.LogoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logoff",
				Handler: auth.LogoffHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: auth.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register/email",
				Handler: auth.RegisterEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/forget/email",
				Handler: auth.ForgetPasswordEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/forget/reset_password",
				Handler: auth.ResetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: auth.OauthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oauth/authorize_url",
				Handler: auth.GetOauthAuthorizeUrlHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/menu/create_menu",
				Handler: menu.CreateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/menu/update_menu",
				Handler: menu.UpdateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/menu/delete_menu",
				Handler: menu.DeleteMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/menu/delete_menu_list",
				Handler: menu.DeleteMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/find_menu",
				Handler: menu.FindMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/find_menu_list",
				Handler: menu.FindMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/sync_menu_list",
				Handler: menu.SyncMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/clean_menu_list",
				Handler: menu.CleanMenuListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/role/create_role",
				Handler: role.CreateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/role/update_role",
				Handler: role.UpdateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/role/delete_role",
				Handler: role.DeleteRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/role/delete_role_list",
				Handler: role.DeleteRoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/find_role",
				Handler: role.FindRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/find_role_list",
				Handler: role.FindRoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/find_role_resources",
				Handler: role.FindRoleResourcesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/update_menus",
				Handler: role.UpdateRoleMenusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/update_apis",
				Handler: role.UpdateRoleApisHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
