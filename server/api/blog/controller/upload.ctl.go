package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type UploadController struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadController(svcCtx *svctx.ServiceContext) *UploadController {
	return &UploadController{
		svcCtx: svcCtx,
	}
}

// @Tags		Upload
// @Summary		"上传文件"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UploadFileReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.UploadFileResp}	"返回信息"
// @Router		/api/v1/upload/upload_file [POST]
func (s *UploadController) UploadFile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req dto.UploadFileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUploadService(s.svcCtx).UploadFile(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
