package website

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/logic/website"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
)

// 访客上报
func ReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := website.NewReportLogic(r.Context(), svcCtx)
		resp, err := l.Report(&req)
		responsex.Response(r, w, resp, err)
	}
}
