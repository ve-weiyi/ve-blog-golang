package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type ApiController struct {
	svcCtx *svctx.ServiceContext
}

func NewApiController(svcCtx *svctx.ServiceContext) *ApiController {
	return &ApiController{
		svcCtx: svcCtx,
	}
}

// @Tags		Api
// @Summary		"创建api路由"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.ApiNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.ApiBackVO}	"返回信息"
// @Router		/admin-api/v1/api/add_api [POST]
func (s *ApiController) AddApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.ApiNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewApiService(s.svcCtx).AddApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Api
// @Summary		"清空接口列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/api/clean_api_list [POST]
func (s *ApiController) CleanApiList(c *gin.Context) {
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

	data, err := service.NewApiService(s.svcCtx).CleanApiList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Api
// @Summary		"删除api路由"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdsReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/api/deletes_api [DELETE]
func (s *ApiController) DeletesApi(c *gin.Context) {
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

	data, err := service.NewApiService(s.svcCtx).DeletesApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Api
// @Summary		"分页获取api路由列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.ApiQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/api/find_api_list [POST]
func (s *ApiController) FindApiList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.ApiQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewApiService(s.svcCtx).FindApiList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Api
// @Summary		"同步api列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.SyncApiReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/api/sync_api_list [POST]
func (s *ApiController) SyncApiList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.SyncApiReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewApiService(s.svcCtx).SyncApiList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Api
// @Summary		"更新api路由"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.ApiNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.ApiBackVO}	"返回信息"
// @Router		/admin-api/v1/api/update_api [PUT]
func (s *ApiController) UpdateApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.ApiNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewApiService(s.svcCtx).UpdateApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
