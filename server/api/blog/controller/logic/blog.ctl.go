package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type BlogController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewBlogController(svcCtx *svc.ControllerContext) *BlogController {
	return &BlogController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// CreateApi 创建api路由 https://www.jianshu.com/p/4bb4283632e4
// @Tags 标签
// @Summary 标题
// @Description 描述,可以有多个
// @Security ApiKeyAuth
// @Param    file  formData  file   true  "上传文件"
// @Param    id	   path      int    true  "id"
// @Param    token header    string true  "token"
// @Param 	 data  body 	 entity.Api		 true  "创建api路由"
// @Success  200   {object}  response.Response{data=entity.Api}  	"返回信息"
// @Router /api/version [get]
func (s *BlogController) ApiVersion(c *gin.Context) {

	s.ResponseOk(c, nil)
}

// CreateApi 创建api路由
// @Tags	 Api
// @Summary  创建api路由
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.Api		true  "创建api路由"
// @Success  200   {object}  response.Response{data=entity.Api}  	"返回信息"
// @Router /about [get]
func (s *BlogController) GetAboutMe(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var api entity.Api
	err = s.ShouldBindJSON(c, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.CreateApi(reqCtx, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
