package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
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
