package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type MenuController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewMenuController(svcCtx *svc.ControllerContext) *MenuController {
	return &MenuController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Menu
// @Summary		创建菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Menu							true	"请求body"
// @Success		200		{object}	response.Response{data=entity.Menu}	"返回信息"
// @Router		/menu/create [post]
func (s *MenuController) CreateMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var menu entity.Menu
	err = s.ShouldBind(c, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.MenuService.CreateMenu(reqCtx, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Menu
// @Summary		删除菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Menu			true	"请求body"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/menu/delete [delete]
func (s *MenuController) DeleteMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var menu entity.Menu
	err = s.ShouldBind(c, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.MenuService.DeleteMenu(reqCtx, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Menu
// @Summary		更新菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Menu							true	"请求body"
// @Success		200		{object}	response.Response{data=entity.Menu}	"返回信息"
// @Router		/menu/update [put]
func (s *MenuController) UpdateMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var menu entity.Menu
	err = s.ShouldBind(c, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.MenuService.UpdateMenu(reqCtx, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Menu
// @Summary		查询菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Menu							true	"请求body"
// @Success		200		{object}	response.Response{data=entity.Menu}	"返回信息"
// @Router		/menu/query [get]
func (s *MenuController) GetMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var menu entity.Menu
	err = s.ShouldBind(c, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.MenuService.GetMenu(reqCtx, &menu)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Menu
// @Summary		批量删除菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		[]int				true	"删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/menu/deleteByIds [delete]
func (s *MenuController) DeleteMenuByIds(c *gin.Context) {
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

	data, err := s.svcCtx.MenuService.DeleteMenuByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Menu
// @Summary		分页获取菜单列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo												true	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Menu}}	"返回信息"
// @Router		/menu/list [get]
func (s *MenuController) FindMenuList(c *gin.Context) {
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

	list, total, err := s.svcCtx.MenuService.FindMenuList(reqCtx, &page)
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
