package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type AuthController struct {
	svcCtx *svc.ServiceContext
}

func NewAuthController(svcCtx *svc.ServiceContext) *AuthController {
	return &AuthController{
		svcCtx: svcCtx,
	}
}

// @Tags		Auth
// @Summary		登录
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.LoginReq			true	"请求body"
// @Success		200		{object}	response.Response{data=dto.LoginResp}	"返回信息"
// @Router		/login [post]
func (s *AuthController) Login(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var user dto.LoginReq
	err = request.ShouldBind(c, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Login(reqCtx, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		登出
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"返回信息"
// @Router		/logout [post]
func (s *AuthController) Logout(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Logout(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		注销
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"返回信息"
// @Router		/logoff [post]
func (s *AuthController) Logoff(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Logoff(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		注册
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.LoginReq			true	"请求body"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"返回信息"
// @Router		/register [post]
func (s *AuthController) Register(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var user dto.LoginReq
	err = request.ShouldBind(c, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Register(reqCtx, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		发送注册邮件
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.UserEmailReq		true	"请求body"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"返回信息"
// @Router		/register/email [post]
func (s *AuthController) SendRegisterEmail(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.UserEmailReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).SendRegisterEmail(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		发送忘记密码邮件
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.UserEmailReq		true	"请求参数"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"返回信息"
// @Router		/forget/email [post]
func (s *AuthController) SendForgetEmail(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var user dto.UserEmailReq
	err = request.ShouldBind(c, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).SendForgetPwdEmail(reqCtx, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		重置密码
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.ResetPasswordReq	true	"请求参数"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"返回信息"
// @Router		/forget/reset_password [post]
func (s *AuthController) ResetPassword(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var user dto.ResetPasswordReq
	err = request.ShouldBind(c, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).ResetPassword(reqCtx, &user)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		第三方登录
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string											false	"token"
// @Param		uid		header		string											false	"uid"
// @Param		data	body		request.OauthLoginReq							true	"请求body"
// @Success		200		{object}	response.Response{data=dto.LoginResp}	"返回信息"
// @Router		/oauth/login [post]
func (s *AuthController) OauthLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.OauthLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).OauthLogin(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		获取授权地址
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string											false	"token"
// @Param		uid		header		string											false	"uid"
// @Param		data	body		request.OauthLoginReq							true	"请求body"
// @Success		200		{object}	response.Response{data=dto.OauthLoginUrl}	"返回信息"
// @Router		/oauth/authorize_url [post]
func (s *AuthController) OauthAuthorizeUrl(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.OauthLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).OauthAuthorizeUrl(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
