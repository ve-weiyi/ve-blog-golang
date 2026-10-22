package websocket

import (
	"net/http"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/logic/common/websocket"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
)

// WebSocket消息
func WebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := websocket.NewWebsocketLogic(r.Context(), svcCtx)
		// 连接升级成功后不能再写 HTTP 响应；失败信息由升级逻辑自行写回
		_ = l.Websocket(w, r)
	}
}
