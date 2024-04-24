package website

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/logic/website"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
)

func UpdateWebsiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsiteConfig
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := website.NewUpdateWebsiteConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateWebsiteConfig(&req)
		responsex.Response(r, w, resp, err)
	}
}
