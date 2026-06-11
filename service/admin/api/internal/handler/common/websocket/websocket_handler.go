package websocket

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/logic/common/websocket"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
)

// WebSocket消息
func WebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := websocket.NewWebsocketLogic(r.Context(), svcCtx)
		err := l.Websocket(w, r)
		responsex.Response(r, w, nil, err)
	}
}
