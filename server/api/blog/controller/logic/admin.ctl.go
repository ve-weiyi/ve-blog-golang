package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type AdminController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewAdminController(svcCtx *svc.ControllerContext) *AdminController {
	return &AdminController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Admin
// @Summary		获取用户地区
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/home [post]
func (s *AdminController) GetHomeInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.BlogService.GetAdminHomeInfo(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Admin
// @Summary		获取用户地区
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo	true	"分页参数"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/user/areas [post]
func (s *AdminController) GetUserAreas(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBindQuery(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.UserService.GetUserAreas(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

// @Tags		Admin
// @Summary		获取用户列表
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo	true	"分页参数"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/users [post]
func (s *AdminController) GetUserList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBindQuery(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.UserService.GetUserList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

// @Tags		Admin
// @Summary		获取用户列表
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo	true	"分页参数"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/comments [post]
func (s *AdminController) GetAdminComments(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBindQuery(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.Log.JsonIndent(page)
	list, total, err := s.svcCtx.CommentService.FindCommonBackList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

// @Tags		Admin
// @Summary		获取角色列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Role}	"返回信息"
// @Router		/admin/roles [post]
func (s *AdminController) GetRoleTreeList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.RoleService.GetRoleTreeList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: int(total),
	})
}

// @Tags		Admin
// @Summary		获取菜单列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Menu}	"返回信息"
// @Router		/admin/menus [post]
func (s *AdminController) GetMenuTreeList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.MenuService.GetAllMenusList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: int(total),
	})
}

// @Tags		Admin
// @Summary		获取api列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Api}	"返回信息"
// @Router		/admin/apis [post]
func (s *AdminController) GetApiTreeList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ApiService.GetAllApiList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: int(total),
	})
}

// @Tags		Role
// @Summary		更新角色菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Role							true	"创建角色"
// @Success		200		{object}	response.Response{data=entity.Role}	"返回信息"
// @Router		/admin/role/update_menus [post]
func (s *AdminController) UpdateRoleMenus(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UpdateRoleMenus
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RoleService.UpdateRoleMenus(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Admin
// @Summary		更新角色资源
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Role							true	"创建角色"
// @Success		200		{object}	response.Response{data=entity.Role}	"返回信息"
// @Router		/admin/role/update_resources [post]
func (s *AdminController) UpdateRoleResources(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UpdateRoleResources
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RoleService.UpdateRoleResources(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Admin
// @Summary		修改用户角色
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.UpdateUserRoles				true	"请求数据"
// @Success		200		{object}	response.Response{data=entity.Role}	"返回信息"
// @Router		/admin/update_roles [post]
func (s *AdminController) UpdateUserRoles(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UpdateUserRoles
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserRoles(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Admin
// @Summary		修改用户状态
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.UserAccount	true	"请求数据"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/update_status [post]
func (s *AdminController) UpdateUserStatus(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req entity.UserAccount
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserStatus(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
