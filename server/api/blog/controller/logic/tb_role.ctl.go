package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
)

// @Tags		Admin
// @Summary		获取角色列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Role}	"返回信息"
// @Router		/admin/roles [post]
func (s *AdminController) GetRoles(c *gin.Context) {
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

	list, total, err := s.svcCtx.RoleService.GetRoles(reqCtx, &page)
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
