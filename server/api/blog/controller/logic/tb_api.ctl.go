package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type ApiController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewApiController(svcCtx *svc.ControllerContext) *ApiController {
	return &ApiController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Api
// @Summary		创建接口
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Api							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Api}	"返回信息"
// @Router		/api/create [post]
func (s *ApiController) CreateApi(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var api entity.Api
	err = s.ShouldBind(c, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.CreateApi(reqCtx, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Api
// @Summary		删除接口
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.Api 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/api/delete [delete]
func (s *ApiController) DeleteApi(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var api entity.Api
	err = s.ShouldBind(c, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.DeleteApi(reqCtx, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Api
// @Summary		更新接口
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.Api							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Api}	"返回信息"
// @Router 		/api/update [put]
func (s *ApiController) UpdateApi(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var api entity.Api
	err = s.ShouldBind(c, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.UpdateApi(reqCtx, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Api
// @Summary		查询接口
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		body		entity.Api							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.Api}	"返回信息"
// @Router 		/api/find [get]
func (s *ApiController) FindApi(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var api entity.Api
	err = s.ShouldBind(c, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.FindApi(reqCtx, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Api
// @Summary		批量删除接口
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/api/deleteByIds [delete]
func (s *ApiController) DeleteApiByIds(c *gin.Context) {
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

	data, err := s.svcCtx.ApiService.DeleteApiByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Api
// @Summary		分页获取接口列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Api}}	"返回信息"
// @Router		/api/list [post]
func (s *ApiController) FindApiList(c *gin.Context) {
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

	list, total, err := s.svcCtx.ApiService.FindApiList(reqCtx, &page)
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
