package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
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
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Upload							true	"请求body"
// @Success		200		{object}	response.Response{data=entity.Upload}	"返回信息"
// @Router		/file/upload/:label [post]
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
