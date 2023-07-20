package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type FriendLinkController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewFriendLinkController(svcCtx *svc.ControllerContext) *FriendLinkController {
	return &FriendLinkController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		FriendLink
// @Summary		创建友链
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.FriendLink							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.FriendLink}	"返回信息"
// @Router		/friendLink/create [post]
func (s *FriendLinkController) CreateFriendLink(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var friendLink entity.FriendLink
	err = s.ShouldBind(c, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.FriendLinkService.CreateFriendLink(reqCtx, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		FriendLink
// @Summary		删除友链
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.FriendLink 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/friendLink/delete [delete]
func (s *FriendLinkController) DeleteFriendLink(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var friendLink entity.FriendLink
	err = s.ShouldBind(c, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.FriendLinkService.DeleteFriendLink(reqCtx, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	FriendLink
// @Summary		更新友链
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.FriendLink							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.FriendLink}	"返回信息"
// @Router 		/friendLink/update [put]
func (s *FriendLinkController) UpdateFriendLink(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var friendLink entity.FriendLink
	err = s.ShouldBind(c, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.FriendLinkService.UpdateFriendLink(reqCtx, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	FriendLink
// @Summary		查询友链
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		query		entity.FriendLink							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.FriendLink}	"返回信息"
// @Router 		/friendLink/find [get]
func (s *FriendLinkController) FindFriendLink(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var friendLink entity.FriendLink
	err = s.ShouldBind(c, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.FriendLinkService.FindFriendLink(reqCtx, &friendLink)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	FriendLink
// @Summary		批量删除友链
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/friendLink/deleteByIds [delete]
func (s *FriendLinkController) DeleteFriendLinkByIds(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var IDS []int
	err = s.ShouldBind(c, &IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.FriendLinkService.DeleteFriendLinkByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	FriendLink
// @Summary		分页获取友链列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.FriendLink}}	"返回信息"
// @Router		/friendLink/list [post]
func (s *FriendLinkController) FindFriendLinkList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.FriendLinkService.FindFriendLinkList(reqCtx, &page)
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