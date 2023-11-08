package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type UploadController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewUploadController(svcCtx *svc.ControllerContext) *UploadController {
	return &UploadController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Upload
// @Summary		文件上传
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		label	path		string									true	"标签"
// @Param		file	formData	file									true	"文件"
// @Success		200		{object}	response.Response{data=entity.UploadRecord}	"返回信息"
// @Router		/upload/{label} [post]
func (s *UploadController) UploadFile(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	// 获取上传的文件标签
	label := c.Param("label")

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UploadService.CreateUpload(reqCtx, label, file)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
