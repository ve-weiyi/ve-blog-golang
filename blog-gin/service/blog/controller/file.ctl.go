package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type FileController struct {
	svcCtx *svctx.ServiceContext
}

func NewFileController(svcCtx *svctx.ServiceContext) *FileController {
	return &FileController{
		svcCtx: svcCtx,
	}
}

// @Tags		File
// @Summary		"上传文件列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.MultiUploadFileReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=[]dto.FileBackVO}	"返回信息"
// @Router		/blog-api/v1/file/multi_upload_file [POST]
func (s *FileController) MultiUploadFile(c *gin.Context) {
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

	data, err := service.NewFileService(s.svcCtx).MultiUploadFile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		File
// @Summary		"上传文件"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UploadFileReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.FileBackVO}	"返回信息"
// @Router		/blog-api/v1/file/upload_file [POST]
func (s *FileController) UploadFile(c *gin.Context) {
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

	data, err := service.NewFileService(s.svcCtx).UploadFile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
