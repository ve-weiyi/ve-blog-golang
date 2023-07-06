package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type UserController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewUserController(svcCtx *svc.ControllerContext) *UserController {
	return &UserController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

//	@Tags		User
//	@Summary	获取用户信息
//	@Security	ApiKeyUser
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.User			true	"请求数据"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/user/info [get]
func (m *UserController) GetUserinfo(c *gin.Context) {
	reqCtx, err := m.GetRequestContext(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.svcCtx.UserService.GetUserinfo(reqCtx, reqCtx.UID)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

//	@Tags		User
//	@Summary	获取用户菜单
//	@Security	ApiKeyUser
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.User			true	"请求数据"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/user/menus [post]
func (m *UserController) GetUserMenus(c *gin.Context) {
	reqCtx, err := m.GetRequestContext(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.svcCtx.MenuService.GetUserMenus(reqCtx, nil)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

// GetUserResources
//
//	@Tags		User
//	@Summary	获取用户资源
//	@Security	ApiKeyUser
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.User			true	"请求数据"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/user/apis [post]
func (m *UserController) GetUserResources(c *gin.Context) {
	reqCtx, err := m.GetRequestContext(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.svcCtx.ApiService.GetUserApis(reqCtx, nil)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

//	@Tags		User
//	@Summary	获取用户登录历史
//	@Security	ApiKeyUser
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.User			true	"请求数据"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/user/login_history [get]
func (m *UserController) GetLoginHistory(c *gin.Context) {
	reqCtx, err := m.GetRequestContext(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = m.ShouldBindQuery(c, &page)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	list, total, err := m.svcCtx.UserService.GetLoginHistory(reqCtx, &page)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

//	@Tags		User
//	@Summary	获取用户列表
//	@Security	ApiKeyUser
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.User			true	"请求数据"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/admin/user/list [post]
func (m *UserController) GetUserList(c *gin.Context) {
	reqCtx, err := m.GetRequestContext(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = m.ShouldBindQuery(c, &page)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	list, total, err := m.svcCtx.UserService.GetUserList(reqCtx, &page)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

//	@Tags		Role
//	@Summary	修改用户角色
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.UpdateUserRoles				true	"请求数据"
//	@Success	200		{object}	response.Response{data=entity.Role}	"返回信息"
//	@Router		/admin/user/update_roles [post]
func (m *UserController) UpdateUserRoles(c *gin.Context) {
	reqCtx, err := m.GetRequestContext(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var req request.UpdateUserRoles
	err = m.ShouldBindJSON(c, &req)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.svcCtx.UserService.UpdateUserRoles(reqCtx, &req)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

//	@Tags		Role
//	@Summary	修改用户状态
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.UserAccount	true	"请求数据"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/admin/user/update_status [post]
func (m *UserController) UpdateUserStatus(c *gin.Context) {
	reqCtx, err := m.GetRequestContext(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var req entity.UserAccount
	err = m.ShouldBindJSON(c, &req)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.svcCtx.UserService.UpdateUserStatus(reqCtx, &req)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}
