package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type BaseController struct {
	svcCtx *svctx.ServiceContext
}

func NewBaseController(svcCtx *svctx.ServiceContext) *BaseController {
	return &BaseController{
		svcCtx: svcCtx,
	}
}

// @Tags		Base
// @Summary		"ping"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.PingReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PingResp}	"返回信息"
// @Router		/admin_api/v1/ping [GET]
func (s *BaseController) Ping(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req dto.PingReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewBaseService(s.svcCtx).Ping(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
