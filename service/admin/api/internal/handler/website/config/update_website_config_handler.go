package config

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/logic/website/config"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
)

// 更新网站配置
func UpdateWebsiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsiteConfigVO
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := config.NewUpdateWebsiteConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateWebsiteConfig(&req)
		responsex.Response(r, w, resp, err)
	}
}
