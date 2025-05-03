package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type WebsocketService struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsocketService(svcCtx *svctx.ServiceContext) *WebsocketService {
	return &WebsocketService{
		svcCtx: svcCtx,
	}
}

// WebSocket消息
func (s *WebsocketService) Websocket(reqCtx *request.Context) (err error) {
	// todo

	return
}
