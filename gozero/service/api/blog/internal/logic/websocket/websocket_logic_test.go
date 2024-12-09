package websocket

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
)

func TestNewWebsocketLogic(t *testing.T) {
	ctx := context.Background()

	svcCtx := &svc.ServiceContext{
		WebsocketManager: ws.NewDefaultClientManager(),
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("websocket")
		NewWebsocketLogic(ctx, svcCtx).Websocket(w, r)
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
