package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
)

// @Tags		Role
// @Summary		获取角色列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.RoleDetailsDTO}}	"返回信息"
// @Router		/role/find_role_details_list [post]
func (s *RoleController) FindRoleDetailsList(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewRoleService(s.svcCtx).FindRoleDetailsList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: total,
	})
}

// @Tags		Role
// @Summary		更新角色菜单
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.UpdateRoleMenusReq			true	"创建角色"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/role/update_role_menus [post]
func (s *RoleController) UpdateRoleMenus(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.UpdateRoleMenusReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewRoleService(s.svcCtx).UpdateRoleMenus(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Role
// @Summary		更新角色资源
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.UpdateRoleApisReq			true	"创建角色"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/role/update_role_resources [post]
func (s *RoleController) UpdateRoleResources(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.UpdateRoleApisReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewRoleService(s.svcCtx).UpdateRoleResources(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
