package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// @Tags		Account
// @Summary		查询用户登录历史
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		page	body		request.PageQuery			true	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.LoginHistory}}	"返回信息"
// @Router		/user/login_history [post]
func (s *UserController) FindUserLoginHistoryList(c *gin.Context) {
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

	list, total, err := s.svcCtx.UserService.FindUserLoginHistoryList(reqCtx, &page)
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

// @Tags		Account
// @Summary		批量删除登录历史
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdsReq				true	"删除id列表"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/user/login_history/batch_delete [delete]
func (s *UserController) DeleteUserLoginHistoryList(c *gin.Context) {
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

	data, err := s.svcCtx.UserService.DeleteUserLoginHistoryList(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.BatchResult{
		SuccessCount: data,
	})
}

// @Tags		Account
// @Summary		获取用户菜单权限
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=[]response.MenuDetailsDTO}	"返回信息"
// @Router		/user/menus [get]
func (s *UserController) GetUserMenus(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.GetUserMenus(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		获取用户接口权限
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=[]response.ApiDetailsDTO}	"返回信息"
// @Router		/user/apis [get]
func (s *UserController) GetUserApis(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.GetUserApis(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		获取用户信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=response.UserInfo}	"返回信息"
// @Router		/user/info [get]
func (s *UserController) GetUserInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.GetUserInfo(reqCtx, reqCtx.UID)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		修改用户信息
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		data	body		request.UserInfoReq					true	"请求body"
// @Success		200		{object}	response.Response{data=entity.UserInformation}	"返回信息"
// @Router		/user/info [post]
func (s *UserController) UpdateUserInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UserInfoReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserInfo(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Account
// @Summary		更换用户头像
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		file	formData	file									true	"文件"
// @Success		200		{object}	response.Response{data=entity.UserInformation}	"返回信息"
// @Router		/user/avatar [post]
func (s *UserController) UpdateUserAvatar(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserAvatar(reqCtx, file)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
