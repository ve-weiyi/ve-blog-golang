package stompws

import (
	"log"
	"net/http"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/ve-weiyi/ve-blog-golang/stompws/logws"
)

func TestWebSocketServer(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	l := logws.NewZapLogger(zap.L())
	l.Infof("hello")

	stompServer := NewWebsocketServer(
		Config{
			Authenticator: nil,
			HeartBeatTime: 5 * time.Millisecond,
			Log:           l,
		},
	)

	go stompServer.Run()
	// 设置 WebSocket 路由
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("ws")
		stompServer.HandleWebSocket(writer, request)
	})

	log.Println("server run on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
