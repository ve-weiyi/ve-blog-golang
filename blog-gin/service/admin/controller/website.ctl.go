package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type WebsiteController struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsiteController(svcCtx *svctx.ServiceContext) *WebsiteController {
	return &WebsiteController{
		svcCtx: svcCtx,
	}
}

// @Tags		Website
// @Summary		"获取用户分布地区"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.GetUserAreaStatsReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.GetUserAreaStatsResp}	"返回信息"
// @Router		/admin-api/v1/account/get_user_area_stats [POST]
func (s *WebsiteController) GetUserAreaStats(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.GetUserAreaStatsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetUserAreaStats(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"获取后台首页信息"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.AdminHomeInfo}	"返回信息"
// @Router		/admin-api/v1/admin [GET]
func (s *WebsiteController) GetAdminHomeInfo(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetAdminHomeInfo(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"获取关于我的信息"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.AboutMeVO}	"返回信息"
// @Router		/admin-api/v1/admin/get_about_me [GET]
func (s *WebsiteController) GetAboutMe(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetAboutMe(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"获取访客数据分析"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.GetVisitStatsResp}	"返回信息"
// @Router		/admin-api/v1/admin/get_visit_stats [GET]
func (s *WebsiteController) GetVisitStats(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetVisitStats(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"获取访客数据趋势"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.GetVisitTrendReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.GetVisitTrendResp}	"返回信息"
// @Router		/admin-api/v1/admin/get_visit_trend [POST]
func (s *WebsiteController) GetVisitTrend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.GetVisitTrendReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetVisitTrend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"获取网站配置"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.WebsiteConfigVO}	"返回信息"
// @Router		/admin-api/v1/admin/get_website_config [GET]
func (s *WebsiteController) GetWebsiteConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetWebsiteConfig(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"获取服务器信息"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.EmptyReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.Server}	"返回信息"
// @Router		/admin-api/v1/admin/system_state [GET]
func (s *WebsiteController) GetSystemState(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetSystemState(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"更新关于我的信息"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.AboutMeVO		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/admin/update_about_me [PUT]
func (s *WebsiteController) UpdateAboutMe(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.AboutMeVO
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).UpdateAboutMe(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		"更新网站配置"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.WebsiteConfigVO		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin-api/v1/admin/update_website_config [PUT]
func (s *WebsiteController) UpdateWebsiteConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.WebsiteConfigVO
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).UpdateWebsiteConfig(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
