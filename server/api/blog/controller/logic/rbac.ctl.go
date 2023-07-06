package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type RBACController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewRBACController(svcCtx *svc.ControllerContext) *RBACController {
	return &RBACController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}
