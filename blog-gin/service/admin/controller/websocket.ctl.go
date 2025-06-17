package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
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
// @Router		/admin-api/v1/websocket [GET]
func (s *WebsocketController) Websocket(c *gin.Context) {
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
