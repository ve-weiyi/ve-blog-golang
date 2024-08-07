package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type RemarkController struct {
	svcCtx *svctx.ServiceContext
}

func NewRemarkController(svcCtx *svctx.ServiceContext) *RemarkController {
	return &RemarkController{
		svcCtx: svcCtx,
	}
}

// @Tags		Remark
// @Summary		创建留言
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		entity.Remark		true	"请求参数"
// @Success		200		{object}	response.Body{data=entity.Remark}	"返回信息"
// @Router		/remark/create_remark [post]
func (s *RemarkController) CreateRemark(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req entity.Remark
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewRemarkService(s.svcCtx).CreateRemark(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Remark
// @Summary		更新留言
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	data	body 	 	entity.Remark		true	"请求参数"
// @Success		200		{object}	response.Body{data=entity.Remark}	"返回信息"
// @Router 		/remark/update_remark [put]
func (s *RemarkController) UpdateRemark(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req entity.Remark
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewRemarkService(s.svcCtx).UpdateRemark(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Remark
// @Summary		删除留言
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdReq				true	"request"
// @Success		200		{object}	response.Body{data=dto.BatchResult}	"返回信息"
// @Router		/remark/delete_remark [delete]
func (s *RemarkController) DeleteRemark(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewRemarkService(s.svcCtx).DeleteRemark(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Remark
// @Summary		批量删除留言
// @Accept 	 	application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdsReq				true	"删除id列表"
// @Success		200		{object}	response.Body{data=dto.BatchResult}	"返回信息"
// @Router		/remark/delete_remark_list [delete]
func (s *RemarkController) DeleteRemarkList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewRemarkService(s.svcCtx).DeleteRemarkList(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.BatchResult{
		SuccessCount: data,
	})
}

// @Tags 	 	Remark
// @Summary		查询留言
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdReq				true	"request"
// @Success		200		{object}	response.Body{data=entity.Remark}	"返回信息"
// @Router 		/remark/find_remark [post]
func (s *RemarkController) FindRemark(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewRemarkService(s.svcCtx).FindRemark(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Remark
// @Summary		分页获取留言列表
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	page 	body		request.PageQuery 			true 	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]entity.Remark}}	"返回信息"
// @Router		/remark/find_remark_list [post]
func (s *RemarkController) FindRemarkList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
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

	list, total, err := service.NewRemarkService(s.svcCtx).FindRemarkList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}
