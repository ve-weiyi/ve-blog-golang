package hub

import (
	"context"
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
)

var Hub *ws.Hub
var l *ChatRoomLogic

func Init(svcCtx *svc.ServiceContext) {
	// http升级成ws后，request.Context 会失效。
	l = NewChatRoomLogic(context.Background(), svcCtx)

	Hub = ws.NewHub()
	go Hub.Run()
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) error {
	client, err := ws.Upgrade(w, r)
	if err != nil {
		return err
	}

	client.SetHandleConnect(func() error {
		return l.HandleConnect(client)
	})

	client.SetHandleDisconnect(func() error {
		return l.HandleDisconnect(client)
	})

	client.SetHandleMessage(func(message []byte) error {
		return l.HandleMessage(client, message)
	})

	return client.Connect(l.svcCtx.Hub)
}
