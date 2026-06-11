package page

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/logic/website/page"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
)

// 创建页面
func CreatePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePageReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := page.NewCreatePageLogic(r.Context(), svcCtx)
		resp, err := l.CreatePage(&req)
		responsex.Response(r, w, resp, err)
	}
}
