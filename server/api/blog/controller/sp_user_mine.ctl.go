package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
)

// @Tags		Account
// @Summary		查询用户登录历史
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		page	body		request.PageQuery			true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.LoginHistory}}	"返回信息"
// @Router		/user/login_history [post]
func (s *UserController) FindUserLoginHistoryList(c *gin.Context) {
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

	list, total, err := service.NewUserService(s.svcCtx).FindUserLoginHistoryList(reqCtx, &page)
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
// @Summary		批量删除登录历史
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdsReq				true	"删除id列表"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/user/login_history/delete_login_history_list [delete]
func (s *UserController) DeleteUserLoginHistoryList(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).DeleteUserLoginHistoryList(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.BatchResult{
		SuccessCount: data,
	})
}

// @Tags		Account
// @Summary		获取用户菜单权限
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Body{data=[]dto.MenuDetailsDTO}	"返回信息"
// @Router		/user/menus [get]
func (s *UserController) GetUserMenus(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).GetUserMenus(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		获取用户接口权限
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Body{data=[]dto.ApiDetailsDTO}	"返回信息"
// @Router		/user/apis [get]
func (s *UserController) GetUserApis(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).GetUserApis(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		获取用户信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Body{data=dto.UserInfo}	"返回信息"
// @Router		/user/get_user_info [get]
func (s *UserController) GetUserInfo(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).GetUserInfo(reqCtx, reqCtx.Uid)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		修改用户信息
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		data	body		request.UserInfoReq					true	"请求body"
// @Success		200		{object}	response.Body{data=entity.UserInformation}	"返回信息"
// @Router		/user/update_user_info [post]
func (s *UserController) UpdateUserInfo(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.UserInfoReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserInfo(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		更换用户头像
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		file	formData	file									true	"文件"
// @Success		200		{object}	response.Body{data=entity.UserInformation}	"返回信息"
// @Router		/user/update_user_avatar [post]
func (s *UserController) UpdateUserAvatar(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserAvatar(reqCtx, file)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
