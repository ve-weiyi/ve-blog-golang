package websocket

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

type WebSocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebSocketLogic {
	return &WebSocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebSocketLogic) WebSocket() error {
	// todo: add your logic here and delete this line

	return nil
}
