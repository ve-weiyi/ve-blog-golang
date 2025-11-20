package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type FriendController struct {
	svcCtx *svctx.ServiceContext
}

func NewFriendController(svcCtx *svctx.ServiceContext) *FriendController {
	return &FriendController{
		svcCtx: svcCtx,
	}
}

// @Tags		Friend
// @Summary		"创建友链"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.FriendNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.FriendBackVO}	"返回信息"
// @Router		/admin-api/v1/friend/add_friend [POST]
func (s *FriendController) AddFriend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.FriendNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewFriendService(s.svcCtx).AddFriend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Friend
// @Summary		"删除友链"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdsReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/friend/deletes_friend [DELETE]
func (s *FriendController) DeletesFriend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.IdsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewFriendService(s.svcCtx).DeletesFriend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Friend
// @Summary		"分页获取友链列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.FriendQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/friend/find_friend_list [POST]
func (s *FriendController) FindFriendList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.FriendQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewFriendService(s.svcCtx).FindFriendList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Friend
// @Summary		"更新友链"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.FriendNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.FriendBackVO}	"返回信息"
// @Router		/admin-api/v1/friend/update_friend [PUT]
func (s *FriendController) UpdateFriend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.FriendNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewFriendService(s.svcCtx).UpdateFriend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
