package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type MenuController struct {
	controller.BaseController
	svcCtx *svc.ServiceContext
}

func NewMenuController(svcCtx *svc.ServiceContext) *MenuController {
	return &MenuController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(),
	}
}

// @Tags		Menu
// @Summary		创建菜单
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		entity.Menu		true	"请求参数"
// @Success		200		{object}	response.Response{data=entity.Menu}	"返回信息"
// @Router		/menu/create_menu [post]
func (s *MenuController) CreateMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req entity.Menu
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewMenuService(s.svcCtx).CreateMenu(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Menu
// @Summary		更新菜单
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	data	body 	 	entity.Menu		true	"请求参数"
// @Success		200		{object}	response.Response{data=entity.Menu}	"返回信息"
// @Router 		/menu/update_menu [put]
func (s *MenuController) UpdateMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req entity.Menu
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewMenuService(s.svcCtx).UpdateMenu(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Menu
// @Summary		删除菜单
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdReq				true	"request"
// @Success		200		{object}	response.Response{data=response.BatchResult}	"返回信息"
// @Router		/menu/delete_menu [delete]
func (s *MenuController) DeleteMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewMenuService(s.svcCtx).DeleteMenu(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Menu
// @Summary		批量删除菜单
// @Accept 	 	application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdsReq				true	"删除id列表"
// @Success		200		{object}	response.Response{data=response.BatchResult}	"返回信息"
// @Router		/menu/delete_menu_list [delete]
func (s *MenuController) DeleteMenuList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.IdsReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewMenuService(s.svcCtx).DeleteMenuList(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.BatchResult{
		SuccessCount: data,
	})
}

// @Tags 	 	Menu
// @Summary		查询菜单
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdReq				true	"request"
// @Success		200		{object}	response.Response{data=entity.Menu}	"返回信息"
// @Router 		/menu/find_menu [post]
func (s *MenuController) FindMenu(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewMenuService(s.svcCtx).FindMenu(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Menu
// @Summary		分页获取菜单列表
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	page 	body		request.PageQuery 			true 	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Menu}}	"返回信息"
// @Router		/menu/find_menu_list [post]
func (s *MenuController) FindMenuList(c *gin.Context) {
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

	list, total, err := service.NewMenuService(s.svcCtx).FindMenuList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}
