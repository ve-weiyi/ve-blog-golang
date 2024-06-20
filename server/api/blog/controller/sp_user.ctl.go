package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type UserController struct {
	svcCtx *svc.ServiceContext
}

func NewUserController(svcCtx *svc.ServiceContext) *UserController {
	return &UserController{
		svcCtx: svcCtx,
	}
}

// @Tags		Account
// @Summary		查询用户列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		page	body		request.PageQuery			true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.UserDTO}}	"返回信息"
// @Router		/user/find_user_list [post]
func (s *UserController) FindUserList(c *gin.Context) {
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

	list, total, err := service.NewUserService(s.svcCtx).FindUserList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Account
// @Summary		查询在线用户列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		page	body		request.PageQuery			true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.UserDTO}}	"返回信息"
// @Router		/user/find_online_user_list [post]
func (s *UserController) FindOnlineUserList(c *gin.Context) {
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

	list, total, err := service.NewUserService(s.svcCtx).FindOnlineUserList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Account
// @Summary		获取用户分布地区
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		page	body		request.PageQuery			true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.UserAreaDTO}}	"返回信息"
// @Router		/user/find_user_areas [post]
func (s *UserController) FindUserAreas(c *gin.Context) {
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

	list, total, err := service.NewUserService(s.svcCtx).FindUserAreaList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Account
// @Summary		修改用户状态
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		entity.UserAccount			true	"请求数据"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/user/update_user_status [post]
func (s *UserController) UpdateUserStatus(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req entity.UserAccount
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserStatus(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		修改用户角色
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.UpdateUserRolesReq				true	"请求数据"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/user/update_user_roles [post]
func (s *UserController) UpdateUserRoles(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.UpdateUserRolesReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserRoles(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
