package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type UserController struct {
	svcCtx *svctx.ServiceContext
}

func NewUserController(svcCtx *svctx.ServiceContext) *UserController {
	return &UserController{
		svcCtx: svcCtx,
	}
}

// @Tags		User
// @Summary		"删除用户绑定第三方平台账号"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.DeleteUserBindThirdPartyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/user/delete_user_bind_third_party [POST]
func (s *UserController) DeleteUserBindThirdParty(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.DeleteUserBindThirdPartyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).DeleteUserBindThirdParty(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"获取用户信息"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.UserInfoResp}	"返回信息"
// @Router		/blog-api/v1/user/get_user_info [GET]
func (s *UserController) GetUserInfo(c *gin.Context) {
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

	data, err := service.NewUserService(s.svcCtx).GetUserInfo(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"获取用户点赞列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.UserLikeResp}	"返回信息"
// @Router		/blog-api/v1/user/get_user_like [GET]
func (s *UserController) GetUserLike(c *gin.Context) {
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

	data, err := service.NewUserService(s.svcCtx).GetUserLike(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"修改用户头像"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateUserAvatarReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/user/update_user_avatar [POST]
func (s *UserController) UpdateUserAvatar(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateUserAvatarReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserAvatar(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"修改用户绑定邮箱"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateUserBindEmailReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/user/update_user_bind_email [POST]
func (s *UserController) UpdateUserBindEmail(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateUserBindEmailReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserBindEmail(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"修改用户绑定手机号"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateUserBindPhoneReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/user/update_user_bind_phone [POST]
func (s *UserController) UpdateUserBindPhone(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateUserBindPhoneReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserBindPhone(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"修改用户绑定第三方平台账号"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateUserBindThirdPartyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/user/update_user_bind_third_party [POST]
func (s *UserController) UpdateUserBindThirdParty(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateUserBindThirdPartyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserBindThirdParty(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"修改用户信息"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateUserInfoReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/user/update_user_info [POST]
func (s *UserController) UpdateUserInfo(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateUserInfoReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserInfo(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		User
// @Summary		"修改用户密码"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateUserPasswordReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/user/update_user_password [POST]
func (s *UserController) UpdateUserPassword(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateUserPasswordReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUserService(s.svcCtx).UpdateUserPassword(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
