package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AccountController struct {
	svcCtx *svctx.ServiceContext
}

func NewAccountController(svcCtx *svctx.ServiceContext) *AccountController {
	return &AccountController{
		svcCtx: svcCtx,
	}
}

// @Tags		Account
// @Summary		"查询用户列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.AccountQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/account/find_account_list [POST]
func (s *AccountController) FindAccountList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.AccountQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAccountService(s.svcCtx).FindAccountList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		"查询在线用户列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.AccountQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/account/find_account_online_list [POST]
func (s *AccountController) FindAccountOnlineList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.AccountQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAccountService(s.svcCtx).FindAccountOnlineList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		"修改用户密码"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateAccountPasswordReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/account/update_account_password [POST]
func (s *AccountController) UpdateAccountPassword(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateAccountPasswordReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAccountService(s.svcCtx).UpdateAccountPassword(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		"修改用户角色"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateAccountRolesReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/account/update_account_roles [POST]
func (s *AccountController) UpdateAccountRoles(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateAccountRolesReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAccountService(s.svcCtx).UpdateAccountRoles(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		"修改用户状态"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateAccountStatusReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/account/update_account_status [POST]
func (s *AccountController) UpdateAccountStatus(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateAccountStatusReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAccountService(s.svcCtx).UpdateAccountStatus(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
