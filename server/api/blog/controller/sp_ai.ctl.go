package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type AIController struct {
	controller.BaseController
	svcCtx *svc.ServiceContext
}

func NewAIController(svcCtx *svc.ServiceContext) *AIController {
	return &AIController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(),
	}
}

// @Tags		AI
// @Summary		和Chatgpt聊天
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ChatMessage				true	"请求body"
// @Success		200		{object}	response.Response{data=chatgpt.ChatResponse}	"返回信息"
// @Router		/ai/chat [post]
func (s *AIController) ChatAI(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.ChatMessage
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatAI(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data.Choices[0].Message.Content)
}

// @Tags		AI
// @Summary		Chatgpt扮演角色
// @accept		application/json
// @Produce		application/json
// @Param		data	body		string				true	"请求body"
// @Param		data	body		request.ChatMessage				true	"请求body"
// @Success		200		{object}	response.Response{data=chatgpt.ChatResponse}	"返回信息"
// @Router		/ai/cos [post]
func (s *AIController) ChatCos(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.ChatMessage
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatCos(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data.Choices[0].Message.Content)
}

// @Tags		AI
// @Summary		和Chatgpt聊天
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ChatStream				true	"请求body"
// @Success		200		{object}	response.Response{data=chatgpt.ChatResponse}	"返回信息"
// @Router		/ai/chat/stream [post]
func (s *AIController) ChatStream(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.ChatStream
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatStream(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.StreamResponse(c, data.Choices[0].Message.Content)
}

// @Tags		AI
// @Summary		和Chatgpt聊天
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ChatMessage				true	"请求body"
// @Success		200		{object}	response.Response{data=chatgpt.ChatResponse}	"返回信息"
// @Router		/ai/assistant [post]
func (s *AIController) ChatAssistant(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.ChatMessage
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatAssistant(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		AI
// @Summary		和Chatgpt聊天
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ChatHistory				true	"请求body"
// @Success		200		{object}	response.Response{data=[]*entity.ChatMessage}	"返回信息"
// @Router		/ai/assistant/history [post]
func (s *AIController) ChatAssistantHistory(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.ChatHistory
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatAssistantHistory(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
