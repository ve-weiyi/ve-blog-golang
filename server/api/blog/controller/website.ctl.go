package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type WebsiteController struct {
	svcCtx *svc.ServiceContext
}

func NewWebsiteController(svcCtx *svc.ServiceContext) *WebsiteController {
	return &WebsiteController{
		svcCtx: svcCtx,
	}
}

// @Tags		Website
// @Summary		查询聊天记录
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string										false	"token"
// @Param		uid		header		string										false	"uid"
// @Param		page	body		request.PageQuery							true	"分页信息"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]entity.ChatRecord}}	"返回信息"
// @Router		/chat/records [post]
func (s *WebsiteController) FindChatRecords(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewChatRecordService(s.svcCtx).FindChatRecordList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: total,
	})
}

// @Tags		Website
// @Summary		获取博客前台首页信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Body{data=dto.BlogHomeInfo}	"返回信息"
// @Router		/ [get]
func (s *WebsiteController) GetBlogHomeInfo(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetBlogHomeInfo(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		获取后台首页信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Body{data=dto.AdminHomeInfo}	"返回信息"
// @Router		/admin [get]
func (s *WebsiteController) GetAdminHomeInfo(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetAdminHomeInfo(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		获取网站配置
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Success		200		{object}	response.Body{data=dto.WebsiteConfigDTO}	"返回信息"
// @Router		/website/config [get]
func (s *WebsiteController) GetWebsiteConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.WebsiteConfigReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetWebsiteConfig(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		更新配置
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.WebsiteConfigDTO		true	"请求信息"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin/config [put]
func (s *WebsiteController) UpdateWebsiteConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.WebsiteConfigDTO
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).UpdateWebsiteConfig(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		获取配置
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.WebsiteConfigDTO		true	"请求信息"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin/config [post]
func (s *WebsiteController) GetConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req dto.WebsiteConfigReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetConfig(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		更新配置
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.WebsiteConfigReq		true	"请求信息"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin/config [put]
func (s *WebsiteController) UpdateConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.WebsiteConfigReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).UpdateConfig(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		关于我
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Success		200		{object}	response.Body{data=dto.AboutMeResp}	"返回信息"
// @Router		/about/me [get]
func (s *WebsiteController) GetAboutMe(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetAboutMe(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		更新我的信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.AboutMeReq			true	"请求信息"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin/about/me [post]
func (s *WebsiteController) UpdateAboutMe(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.AboutMeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).UpdateAboutMe(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		获取服务器信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin/system/state [get]
func (s *WebsiteController) GetSystemState(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewWebsiteService(s.svcCtx).GetSystemState(reqCtx, nil)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
