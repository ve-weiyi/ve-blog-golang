package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
)

// @Tags		Role
// @Summary		获取角色列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.RoleDetailsDTO}}	"返回信息"
// @Router		/role/find_role_details_list [post]
func (s *RoleController) FindRoleDetailsList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := service.NewRoleService(s.svcCtx).FindRoleDetailsList(reqCtx, &page)
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
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.UpdateRoleMenusReq			true	"创建角色"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/role/update_role_menus [post]
func (s *RoleController) UpdateRoleMenus(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UpdateRoleMenusReq
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewRoleService(s.svcCtx).UpdateRoleMenus(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Role
// @Summary		更新角色资源
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.UpdateRoleApisReq			true	"创建角色"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/role/update_role_resources [post]
func (s *RoleController) UpdateRoleResources(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UpdateRoleApisReq
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewRoleService(s.svcCtx).UpdateRoleResources(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
