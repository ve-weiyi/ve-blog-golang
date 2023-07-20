package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type OperationLogController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewOperationLogController(svcCtx *svc.ControllerContext) *OperationLogController {
	return &OperationLogController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		OperationLog
// @Summary		创建操作记录
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.OperationLog							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.OperationLog}	"返回信息"
// @Router		/operationLog/create [post]
func (s *OperationLogController) CreateOperationLog(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var operationLog entity.OperationLog
	err = s.ShouldBind(c, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.OperationLogService.CreateOperationLog(reqCtx, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		OperationLog
// @Summary		删除操作记录
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.OperationLog 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/operationLog/delete [delete]
func (s *OperationLogController) DeleteOperationLog(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var operationLog entity.OperationLog
	err = s.ShouldBind(c, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.OperationLogService.DeleteOperationLog(reqCtx, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	OperationLog
// @Summary		更新操作记录
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.OperationLog							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.OperationLog}	"返回信息"
// @Router 		/operationLog/update [put]
func (s *OperationLogController) UpdateOperationLog(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var operationLog entity.OperationLog
	err = s.ShouldBind(c, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.OperationLogService.UpdateOperationLog(reqCtx, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	OperationLog
// @Summary		查询操作记录
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		body		entity.OperationLog							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.OperationLog}	"返回信息"
// @Router 		/operationLog/find [get]
func (s *OperationLogController) FindOperationLog(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var operationLog entity.OperationLog
	err = s.ShouldBind(c, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.OperationLogService.FindOperationLog(reqCtx, &operationLog)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	OperationLog
// @Summary		批量删除操作记录
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/operationLog/deleteByIds [delete]
func (s *OperationLogController) DeleteOperationLogByIds(c *gin.Context) {
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

	data, err := s.svcCtx.OperationLogService.DeleteOperationLogByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	OperationLog
// @Summary		分页获取操作记录列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.OperationLog}}	"返回信息"
// @Router		/operationLog/list [post]
func (s *OperationLogController) FindOperationLogList(c *gin.Context) {
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

	list, total, err := s.svcCtx.OperationLogService.FindOperationLogList(reqCtx, &page)
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
