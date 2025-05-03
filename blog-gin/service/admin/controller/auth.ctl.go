package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AuthController struct {
	svcCtx *svctx.ServiceContext
}

func NewAuthController(svcCtx *svctx.ServiceContext) *AuthController {
	return &AuthController{
		svcCtx: svcCtx,
	}
}

// @Tags		Auth
// @Summary		"获取游客身份信息"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.GetTouristInfoResp}	"返回信息"
// @Router		/admin-api/v1/get_tourist_info [GET]
func (s *AuthController) GetTouristInfo(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).GetTouristInfo(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"邮箱登录"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmailLoginReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.LoginResp}	"返回信息"
// @Router		/admin-api/v1/email_login [POST]
func (s *AuthController) EmailLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmailLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).EmailLogin(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"获取验证码"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.GetCaptchaCodeReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.GetCaptchaCodeResp}	"返回信息"
// @Router		/admin-api/v1/get_captcha_code [POST]
func (s *AuthController) GetCaptchaCode(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.GetCaptchaCodeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).GetCaptchaCode(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"第三方登录授权地址"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.GetOauthAuthorizeUrlReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.GetOauthAuthorizeUrlResp}	"返回信息"
// @Router		/admin-api/v1/get_oauth_authorize_url [POST]
func (s *AuthController) GetOauthAuthorizeUrl(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.GetOauthAuthorizeUrlReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).GetOauthAuthorizeUrl(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"登录"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.LoginReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.LoginResp}	"返回信息"
// @Router		/admin-api/v1/login [POST]
func (s *AuthController) Login(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.LoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Login(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"手机登录"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.PhoneLoginReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.LoginResp}	"返回信息"
// @Router		/admin-api/v1/phone_login [POST]
func (s *AuthController) PhoneLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.PhoneLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).PhoneLogin(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"注册"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.RegisterReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/register [POST]
func (s *AuthController) Register(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.RegisterReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Register(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"重置密码"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.ResetPasswordReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/reset_password [POST]
func (s *AuthController) ResetPassword(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.ResetPasswordReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).ResetPassword(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"发送邮件验证码"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.SendEmailVerifyCodeReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/send_email_verify_code [POST]
func (s *AuthController) SendEmailVerifyCode(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.SendEmailVerifyCodeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).SendEmailVerifyCode(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"发送手机验证码"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.SendPhoneVerifyCodeReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/send_phone_verify_code [POST]
func (s *AuthController) SendPhoneVerifyCode(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.SendPhoneVerifyCodeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).SendPhoneVerifyCode(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"第三方登录"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.ThirdLoginReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.LoginResp}	"返回信息"
// @Router		/admin-api/v1/third_login [POST]
func (s *AuthController) ThirdLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.ThirdLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).ThirdLogin(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"注销"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/logoff [POST]
func (s *AuthController) Logoff(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Logoff(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		"登出"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/logout [POST]
func (s *AuthController) Logout(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAuthService(s.svcCtx).Logout(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
