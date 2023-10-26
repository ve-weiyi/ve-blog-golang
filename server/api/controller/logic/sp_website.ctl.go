package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/ws"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

type WebsiteController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewWebsiteController(svcCtx *svc.ControllerContext) *WebsiteController {
	return &WebsiteController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Website
// @Summary		查询聊天记录
// @Router		/ws [get]
func (s *WebsiteController) WebSocket(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	// 接收消息
	receive := func(msg []byte) {
		global.LOG.Println(string(msg))

		var chat entity.ChatRecord
		err = jsonconv.JsonToObject(string(msg), &chat)
		if err != nil {
			global.LOG.Error(err)
		}

		if chat.Content == "" {
			return
		}
		if reqCtx.UID != 0 {
			chat.UserID = reqCtx.UID
		}

		_, err = s.svcCtx.ChatRecordService.CreateChatRecord(reqCtx, &chat)
		if err != nil {
			global.LOG.Error(err)
		}
	}

	ws.HandleWebSocket(c.Writer, c.Request, receive)
}

// @Tags		Website
// @Summary		关于我
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/about/me [get]
func (s *WebsiteController) GetAboutMe(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteConfigService.GetAboutMe(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		更新我的信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		string						true	"请求信息"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/about/me [post]
func (s *WebsiteController) UpdateAboutMe(c *gin.Context) {
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

// @Tags		Website
// @Summary		获取网站配置
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.WebsiteConfigRequest		true	"请求信息"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/website/config [post]
func (s *WebsiteController) GetWebsiteConfig(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}
	var req request.WebsiteConfigRequest
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteConfigService.GetConfig(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		更新网站配置
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string								false	"token"
// @Param		uid		header		string								false	"uid"
// @Param		data	body		request.WebsiteConfigRequest		true	"请求信息"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/website/config [put]
func (s *WebsiteController) UpdateWebsiteConfig(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.WebsiteConfigRequest
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteConfigService.UpdateConfig(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		获取后台首页信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=response.WebsiteAdminHomeInfo}	"返回信息"
// @Router		/admin/home [get]
func (s *WebsiteController) GetAdminHomeInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteService.GetWebsiteAdminHomeInfo(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Website
// @Summary		查询聊天记录
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string										false	"token"
// @Param		uid		header		string										false	"uid"
// @Param		page	body		request.PageQuery							true	"分页信息"
// @Success		200		{object}	response.Response{data=entity.ChatRecord}	"返回信息"
// @Router		/chat/records [post]
func (s *WebsiteController) FindChatRecords(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ChatRecordService.FindChatRecordList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: int(total),
	})
}

// @Tags		Website
// @Summary		获取服务器信息
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Success		200		{object}	response.Response{data=response.WebsiteAdminHomeInfo}	"返回信息"
// @Router		/admin/system/state [get]
func (s *WebsiteController) GetSystemState(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteConfigService.GetSystemState(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
