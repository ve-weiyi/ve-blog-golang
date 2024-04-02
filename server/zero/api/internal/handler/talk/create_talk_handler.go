package talk

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/talk"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"
)

func CreateTalkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Talk
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := talk.NewCreateTalkLogic(r.Context(), svcCtx)
		resp, err := l.CreateTalk(&req)
		responsex.Response(r, w, resp, err)
	}
}
