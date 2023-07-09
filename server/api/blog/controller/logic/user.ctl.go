package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type UserController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewUserController(svcCtx *svc.ControllerContext) *UserController {
	return &UserController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Auth
// @Summary		重置密码
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ResetPasswordReq			true	"请求参数"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/user/password/reset [post]
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

	data, err := s.svcCtx.UserService.ResetPassword(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Auth
// @Summary		发送忘记密码邮件
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.UserEmail					true	"请求参数"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/user/password/forget [post]
func (s *UserController) ForgetPasswordEmail(c *gin.Context) {
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

	data, err := s.svcCtx.UserService.SendForgetPwdEmail(reqCtx, &user)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		获取用户信息
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/user/info [get]
func (s *UserController) GetUserinfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.GetUserinfo(reqCtx, reqCtx.UID)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		获取用户菜单权限
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/user/menus [post]
func (s *UserController) GetUserMenus(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.MenuService.GetUserMenus(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		获取用户接口权限
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/user/apis [post]
func (s *UserController) GetUserApis(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.GetUserApis(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		获取用户登录历史
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo	true	"分页参数"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/user/login_history [get]
func (s *UserController) GetLoginHistory(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBindQuery(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.UserService.GetLoginHistory(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}
