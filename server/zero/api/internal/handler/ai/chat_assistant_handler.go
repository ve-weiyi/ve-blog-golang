package ai

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/ai"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"
)

func ChatAssistantHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatMessage
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ai.NewChatAssistantLogic(r.Context(), svcCtx)
		resp, err := l.ChatAssistant(&req)
		responsex.Response(r, w, resp, err)
	}
}
