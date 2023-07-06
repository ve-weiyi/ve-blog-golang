package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type AuthController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewAuthController(svcCtx *svc.ControllerContext) *AuthController {
	return &AuthController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Auth
// @Summary	登录
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Auth							true	"创建权限认证"
// @Success	200		{object}	response.Response{data=entity.Auth}	"返回信息"
// @Router		/user/login [post]
func (s *AuthController) Login(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var user request.User
	err = s.ShouldBind(c, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.AuthService.Login(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary	登出
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Auth							true	"创建权限认证"
// @Success	200		{object}	response.Response{data=entity.Auth}	"返回信息"
// @Router		/user/logout [get]
func (s *AuthController) Logout(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.Log.Println("Logout")
	data, err := s.svcCtx.AuthService.Logout(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary	注销
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Auth			true	"请求body"
// @Success	200		{object}	response.Response{}	"返回信息"
// @Router		/user/logoff [post]
func (s *AuthController) Logoff(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.AuthService.Logoff(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary	注册
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Auth			true	"删除权限认证"
// @Success	200		{object}	response.Response{}	"返回信息"
// @Router		/user/register [post]
func (s *AuthController) Register(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var user request.User
	err = s.ShouldBind(c, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.AuthService.Register(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary	重置密码
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Auth							true	"创建权限认证"
// @Success	200		{object}	response.Response{data=entity.Auth}	"返回信息"
// @Router		/user/password/reset [post]
func (s *AuthController) ResetPassword(c *gin.Context) {
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

	data, err := s.svcCtx.AuthService.ResetPassword(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// Register
//
//	@Tags		Auth
//	@Summary	发送注册邮件
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.Auth			true	"删除权限认证"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/user/register/email [post]
func (s *AuthController) RegisterEmail(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.UserEmail
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.AuthService.SendRegisterEmail(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary	发送忘记密码邮件
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Auth							true	"创建权限认证"
// @Success	200		{object}	response.Response{data=entity.Auth}	"返回信息"
// @Router		/user/password/forget [post]
func (s *AuthController) ForgetPasswordEmail(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var user request.UserEmail
	err = s.ShouldBind(c, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.AuthService.SendForgetPwdEmail(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary	获取授权地址
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		request.OauthLoginReq							true	"请求body信息"
// @Success	200		{object}	response.Response{data=response.OauthLoginUrl}	"返回信息"
// @Router		/user/oauth/login [post]
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

	data, err := s.svcCtx.AuthService.OauthLogin(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary	获取授权地址
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		request.OauthLoginReq							true	"请求body信息"
// @Success	200		{object}	response.Response{data=response.OauthLoginUrl}	"返回信息"
// @Router		/user/oauth/url [post]
func (s *AuthController) GetAuthorizeUrl(c *gin.Context) {
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

	data, err := s.svcCtx.AuthService.GetAuthorizeUrl(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
