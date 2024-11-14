package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/gin/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
)

type WebsocketController struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsocketController(svcCtx *svctx.ServiceContext) *WebsocketController {
	return &WebsocketController{
		svcCtx: svcCtx,
	}
}

// @Tags		Websocket
// @Summary		"WebSocket消息"
// @accept		application/json
// @Produce		application/json
// @Router		/api/v1/websocket [GET]
func (s *WebsocketController) WebSocket(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	err = service.NewWebsocketService(s.svcCtx).WebSocket(reqCtx)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, nil)
}
