package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type CaptchaController struct {
	svcCtx *svctx.ServiceContext
}

func NewCaptchaController(svcCtx *svctx.ServiceContext) *CaptchaController {
	return &CaptchaController{
		svcCtx: svcCtx,
	}
}

// @Tags		Captcha
// @Summary		发送验证码
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.CaptchaEmailReq		true	"请求body"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"返回信息"
// @Router		/captcha/email [post]
func (s *CaptchaController) SendCaptchaEmail(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.CaptchaEmailReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCaptchaService(s.svcCtx).SendCaptchaEmail(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Captcha
// @Summary		生成验证码
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.CaptchaReq				true	"请求body"
// @Success		200		{object}	response.Response{data=dto.CaptchaDTO}	"生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router		/captcha/image [post]
func (s *CaptchaController) GetCaptchaImage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.CaptchaReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCaptchaService(s.svcCtx).GetCaptchaImage(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Captcha
// @Summary		检验验证码
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.CaptchaVerifyReq		true	"请求body"
// @Success		200		{object}	response.Response{data=dto.EmptyResp}	"生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router		/captcha/verify [post]
func (s *CaptchaController) VerifyCaptcha(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.CaptchaVerifyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCaptchaService(s.svcCtx).VerifyImageCaptcha(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
