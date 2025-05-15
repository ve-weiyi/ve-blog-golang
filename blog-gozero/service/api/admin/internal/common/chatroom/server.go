package chatroom

import (
	"github.com/ve-weiyi/ve-blog-golang/stompws"
)

func Init(server *stompws.StompWebsocketServer) {
	server.RegisterTopicHooks("/topic/online-count", stompws.NewOnlineHook(server.Log))
	server.RegisterTopicHooks("/topic/online-users", stompws.NewOnlineHook(server.Log))
}
