package website

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/website"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
)

// 更新配置
func UpdateWebsiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqCtx types.RestHeader
		if err := httpx.ParseHeaders(r, &reqCtx); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var req types.WebsiteConfig
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := website.NewUpdateWebsiteConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateWebsiteConfig(&reqCtx, &req)
		responsex.Response(r, w, resp, err)
	}
}
