package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type ChatController struct {
	svcCtx *svctx.ServiceContext
}

func NewChatController(svcCtx *svctx.ServiceContext) *ChatController {
	return &ChatController{
		svcCtx: svcCtx,
	}
}

// @Tags		Chat
// @Summary		"查询聊天记录"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.ChatQueryReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/api/v1/chat/records [POST]
func (s *ChatController) GetChatRecords(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req dto.ChatQueryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewChatService(s.svcCtx).GetChatRecords(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
