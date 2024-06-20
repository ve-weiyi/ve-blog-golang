package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
)

// @Tags		Api
// @Summary		获取接口列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.ApiDetailsDTO}}	"返回信息"
// @Router		/api/find_api_details_list [post]
func (s *ApiController) FindApiDetailsList(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewApiService(s.svcCtx).FindApiDetailsList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     0,
		PageSize: int64(len(list)),
	})
}

// @Tags		Api
// @Summary		同步接口列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		data	body		request.EmptyReq						true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResult}	"返回信息"
// @Router		/api/sync_api_list [post]
func (s *ApiController) SyncApiList(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewApiService(s.svcCtx).SyncApiList(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.BatchResult{
		SuccessCount: data,
	})
}

// @Tags		Api
// @Summary		清空接口列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		data	body		request.EmptyReq						true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}				"返回信息"
// @Router		/api/clean_api_list [post]
func (s *ApiController) CleanApiList(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewApiService(s.svcCtx).CreateApi(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
