package message

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/logic/message"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
)

// 分页获取留言列表
func FindMessageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := message.NewFindMessageListLogic(r.Context(), svcCtx)
		resp, err := l.FindMessageList(&req)
		responsex.Response(r, w, resp, err)
	}
}
