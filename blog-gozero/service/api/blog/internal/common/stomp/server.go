package stomp

import (
	"context"
	"net/http"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/stompws"
	"github.com/ve-weiyi/ve-blog-golang/stompws/logws"
)

var stompServer *stompws.StompWebsocketServer

func Init(svcCtx *svc.ServiceContext) {
	l := logws.NewZapLogger(glog.Default().Logger())
	stompServer = stompws.NewWebsocketServer(
		stompws.Config{
			Authenticator: nil,
			HeartBeatTime: 5 * time.Millisecond,
			Log:           l,
		},
	)

	stompServer.RegisterTopicHooks("/topic/chatroom", NewChatRoomTopicHook(context.Background(), svcCtx))

	go stompServer.Run()
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) error {
	return stompServer.HandleWebSocket(w, r)
}
