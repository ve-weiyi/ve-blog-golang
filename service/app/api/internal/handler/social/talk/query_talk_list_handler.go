package talk

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/logic/social/talk"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/types"
)

// 获取说说列表
func QueryTalkListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTalkListReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := talk.NewQueryTalkListLogic(r.Context(), svcCtx)
		resp, err := l.QueryTalkList(&req)
		responsex.Response(r, w, resp, err)
	}
}
