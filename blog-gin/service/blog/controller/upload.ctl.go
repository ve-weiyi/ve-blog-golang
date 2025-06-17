package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
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
// @Summary		"删除文件列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.DeletesUploadFileReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/blog-api/v1/upload/deletes_upload_file [DELETE]
func (s *UploadController) DeletesUploadFile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.DeletesUploadFileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUploadService(s.svcCtx).DeletesUploadFile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Upload
// @Summary		"获取文件列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.ListUploadFileReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/blog-api/v1/upload/list_upload_file [POST]
func (s *UploadController) ListUploadFile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.ListUploadFileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUploadService(s.svcCtx).ListUploadFile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Upload
// @Summary		"上传文件列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.MultiUploadFileReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=[]dto.FileInfoVO}	"返回信息"
// @Router		/blog-api/v1/upload/multi_upload_file [POST]
func (s *UploadController) MultiUploadFile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.MultiUploadFileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUploadService(s.svcCtx).MultiUploadFile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Upload
// @Summary		"上传文件"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UploadFileReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.FileInfoVO}	"返回信息"
// @Router		/blog-api/v1/upload/upload_file [POST]
func (s *UploadController) UploadFile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UploadFileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUploadService(s.svcCtx).UploadFile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
