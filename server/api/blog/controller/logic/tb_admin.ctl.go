package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type AdminController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewAdminController(svcCtx *svc.ControllerContext) *AdminController {
	return &AdminController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}
