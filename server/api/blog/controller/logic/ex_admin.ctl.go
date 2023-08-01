package logic

import (
	"github.com/gin-gonic/gin"

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

// @Tags		Admin
// @Summary		更新我的信息
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/about [post]
func (s *AdminController) UpdateAboutMe(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req string
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteConfigService.UpdateAboutMe(reqCtx, req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Admin
// @Summary		获取用户地区
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/home [post]
func (s *AdminController) GetHomeInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.BlogService.GetAdminHomeInfo(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
