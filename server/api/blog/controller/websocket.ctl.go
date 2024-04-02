package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type WebsocketController struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsocketController(svcCtx *svctx.ServiceContext) *WebsocketController {
	return &WebsocketController{
		svcCtx: svcCtx,
	}
}

// @Tags		WebSocket
// @Summary		WebSocket消息
// @Router		/ws [get]
func (s *WebsocketController) WebSocket(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	// 接收消息
	receive := func(msg []byte) (err error) {

		var chat entity.ChatRecord
		err = json.Unmarshal(msg, &chat)
		if err != nil {
			return err
		}

		if chat.Content == "" {
			return fmt.Errorf("content is empty")
		}
		if reqCtx.Uid != 0 {
			chat.UserId = reqCtx.Uid
		}

		_, err = service.NewChatRecordService(s.svcCtx).CreateChatRecord(reqCtx, &chat)
		if err != nil {
			return err
		}

		return nil
	}

	ws.HandleWebSocket(c.Writer, c.Request, receive)
}
