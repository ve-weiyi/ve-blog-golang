package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/chatgpt"
)

type AIController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewAIController(svcCtx *svc.ControllerContext) *AIController {
	return &AIController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
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

	data, err := s.svcCtx.AIChatService.ChatAssistant(reqCtx, &req)
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

	data, err := s.svcCtx.AIChatService.ChatAssistantHistory(reqCtx, &req)
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
// @Param		data	body		[]chatgpt.ChatMessage				true	"请求body"
// @Success		200		{object}	response.Response{data=chatgpt.ChatResponse}	"返回信息"
// @Router		/ai/chat [post]
func (s *AIController) ChatAI(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req []*chatgpt.ChatMessage
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.AIChatService.ChatAI(reqCtx, req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		AI
// @Summary		Chatgpt扮演角色
// @accept		application/json
// @Produce		application/json
// @Param		data	body		string				true	"请求body"
// @Success		200		{object}	response.Response{data=chatgpt.ChatResponse}	"返回信息"
// @Router		/ai/cos [post]
func (s *AIController) ChatCos(c *gin.Context) {
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

	data, err := s.svcCtx.AIChatService.ChatCos(reqCtx, req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
