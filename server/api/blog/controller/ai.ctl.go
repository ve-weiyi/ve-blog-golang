package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type AIController struct {
	svcCtx *svctx.ServiceContext
}

func NewAIController(svcCtx *svctx.ServiceContext) *AIController {
	return &AIController{
		svcCtx: svcCtx,
	}
}

// @Tags		AI
// @Summary		和Chatgpt聊天
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ChatMessage				true	"请求body"
// @Success		200		{object}	response.Response{data=entity.ChatMessage}	"返回信息"
// @Router		/ai/chat [post]
func (s *AIController) ChatAI(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ChatMessage
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatAI(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data.Choices[0].Message.Content)
}

// @Tags		AI
// @Summary		Chatgpt扮演角色
// @accept		application/json
// @Produce		application/json
// @Param		data	body		string				true	"请求body"
// @Param		data	body		request.ChatMessage				true	"请求body"
// @Success		200		{object}	response.Response{data=entity.ChatMessage}	"返回信息"
// @Router		/ai/cos [post]
func (s *AIController) ChatCos(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ChatMessage
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatCos(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data.Choices[0].Message.Content)
}

// @Tags		AI
// @Summary		和Chatgpt聊天
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ChatStream				true	"请求body"
// @Success		200		{object}	response.Response{data=entity.ChatMessage}	"返回信息"
// @Router		/ai/chat/stream [post]
func (s *AIController) ChatStream(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ChatStream
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatStream(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseStream(c, data.Choices[0].Message.Content)
}

// @Tags		AI
// @Summary		和Chatgpt聊天历史记录
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.ChatHistory				true	"请求body"
// @Success		200		{object}	response.Response{data=[]entity.ChatMessage}	"返回信息"
// @Router		/ai/assistant/history [post]
func (s *AIController) ChatAssistantHistory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ChatHistory
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAIService(s.svcCtx).ChatAssistantHistory(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
