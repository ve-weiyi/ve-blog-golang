package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type WebsocketController struct {
	controller.BaseController
	svcCtx *svc.ServiceContext
}

func NewWebsocketController(svcCtx *svc.ServiceContext) *WebsocketController {
	return &WebsocketController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(),
	}
}

// @Tags		WebSocket
// @Summary		WebSocket消息
// @Router		/ws [get]
func (s *WebsocketController) WebSocket(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	// 接收消息
	receive := func(msg []byte) (tx []byte, err error) {
		glog.Println(string(msg))

		var chat entity.ChatRecord
		err = jsonconv.JsonToObject(string(msg), &chat)
		if err != nil {
			return nil, err
		}

		if chat.Content == "" {
			return nil, fmt.Errorf("content is empty")
		}
		if reqCtx.Uid != 0 {
			chat.UserId = reqCtx.Uid
		}

		_, err = service.NewChatRecordService(s.svcCtx).CreateChatRecord(reqCtx, &chat)
		if err != nil {
			return nil, err
		}

		return msg, nil
	}

	ws.HandleWebSocket(c.Writer, c.Request, receive)
}
