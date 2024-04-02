package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// @Tags		Api
// @Summary		获取接口列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.ApiDetailsDTO}}	"返回信息"
// @Router		/api/find_api_details_list [post]
func (s *ApiController) FindApiDetailsList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ApiService.FindApiDetailsList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     0,
		PageSize: len(list),
	})
}

// @Tags		Api
// @Summary		同步接口列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		data	body		request.EmptyReq						true	"请求参数"
// @Success		200		{object}	response.Response{data=response.BatchResult}	"返回信息"
// @Router		/api/sync_api_list [post]
func (s *ApiController) SyncApiList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.SyncApiList(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.BatchResult{
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
// @Success		200		{object}	response.Response{data=response.EmptyResp}				"返回信息"
// @Router		/api/clean_api_list [post]
func (s *ApiController) CleanApiList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.CreateApi(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
