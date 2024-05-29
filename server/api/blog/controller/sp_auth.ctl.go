package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type AuthController struct {
	controller.BaseController
	svcCtx *svc.ServiceContext
}

func NewAuthController(svcCtx *svc.ServiceContext) *AuthController {
	return &AuthController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(),
	}
}

// @Tags		Auth
// @Summary		登录
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.LoginReq			true	"请求body"
// @Success		200		{object}	response.Response{data=response.LoginResp}	"返回信息"
// @Router		/login [post]
func (s *AuthController) Login(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var user request.LoginReq
	err = s.ShouldBind(c, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Login(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		登出
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/logout [post]
func (s *AuthController) Logout(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Logout(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		注销
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/logoff [post]
func (s *AuthController) Logoff(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Logoff(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		注册
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.LoginReq			true	"请求body"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/register [post]
func (s *AuthController) Register(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var user request.LoginReq
	err = s.ShouldBind(c, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Register(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		发送注册邮件
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.UserEmailReq		true	"请求body"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/register/email [post]
func (s *AuthController) RegisterEmail(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UserEmailReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).SendRegisterEmail(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		发送忘记密码邮件
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.UserEmailReq		true	"请求参数"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/forget/email [post]
func (s *UserController) ForgetPasswordEmail(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var user request.UserEmailReq
	err = s.ShouldBind(c, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).SendForgetPwdEmail(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		重置密码
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.ResetPasswordReq	true	"请求参数"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/forget/reset_password [post]
func (s *UserController) ResetPassword(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var user request.ResetPasswordReq
	err = s.ShouldBind(c, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).ResetPassword(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		第三方登录
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string											false	"token"
// @Param		uid		header		string											false	"uid"
// @Param		data	body		request.OauthLoginReq							true	"请求body"
// @Success		200		{object}	response.Response{data=response.LoginResp}	"返回信息"
// @Router		/oauth/login [post]
func (s *AuthController) OauthLogin(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.OauthLoginReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).OauthLogin(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		获取授权地址
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string											false	"token"
// @Param		uid		header		string											false	"uid"
// @Param		data	body		request.OauthLoginReq							true	"请求body"
// @Success		200		{object}	response.Response{data=response.OauthLoginUrl}	"返回信息"
// @Router		/oauth/authorize_url [post]
func (s *AuthController) GetOauthAuthorizeUrl(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.OauthLoginReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).GetAuthorizeUrl(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
