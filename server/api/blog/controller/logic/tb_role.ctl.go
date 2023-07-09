package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type RoleController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewRoleController(svcCtx *svc.ControllerContext) *RoleController {
	return &RoleController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Role
// @Summary		创建角色
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Role							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Role}	"返回信息"
// @Router		/role/create [post]
func (s *RoleController) CreateRole(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var role entity.Role
	err = s.ShouldBind(c, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RoleService.CreateRole(reqCtx, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Role
// @Summary		删除角色
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.Role 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/role/delete [delete]
func (s *RoleController) DeleteRole(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var role entity.Role
	err = s.ShouldBind(c, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RoleService.DeleteRole(reqCtx, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Role
// @Summary		更新角色
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.Role							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Role}	"返回信息"
// @Router 		/role/update [put]
func (s *RoleController) UpdateRole(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var role entity.Role
	err = s.ShouldBind(c, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RoleService.UpdateRole(reqCtx, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Role
// @Summary		查询角色
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		body		entity.Role							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.Role}	"返回信息"
// @Router 		/role/find [get]
func (s *RoleController) FindRole(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var role entity.Role
	err = s.ShouldBind(c, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RoleService.FindRole(reqCtx, &role)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Role
// @Summary		批量删除角色
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/role/deleteByIds [delete]
func (s *RoleController) DeleteRoleByIds(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var IDS []int
	err = s.ShouldBind(c, &IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RoleService.DeleteRoleByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Role
// @Summary		分页获取角色列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Role}}	"返回信息"
// @Router		/role/list [post]
func (s *RoleController) FindRoleList(c *gin.Context) {
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

	list, total, err := s.svcCtx.RoleService.FindRoleList(reqCtx, &page)
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
