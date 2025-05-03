package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type VisitLogController struct {
	svcCtx *svctx.ServiceContext
}

func NewVisitLogController(svcCtx *svctx.ServiceContext) *VisitLogController {
	return &VisitLogController{
		svcCtx: svcCtx,
	}
}

// @Tags		VisitLog
// @Summary		"删除操作记录"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdsReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/visit_log/deletes_visit_log [DELETE]
func (s *VisitLogController) DeletesVisitLog(c *gin.Context) {
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

	data, err := service.NewVisitLogService(s.svcCtx).DeletesVisitLog(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		VisitLog
// @Summary		"分页获取操作记录列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.VisitLogQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/visit_log/find_visit_log_list [POST]
func (s *VisitLogController) FindVisitLogList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.VisitLogQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewVisitLogService(s.svcCtx).FindVisitLogList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
