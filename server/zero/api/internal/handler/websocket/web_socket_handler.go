package websocket

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/websocket"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func WebSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := websocket.NewWebSocketLogic(r.Context(), svcCtx)
		err := l.WebSocket()
		responsex.Response(r, w, nil, err)
	}
}
