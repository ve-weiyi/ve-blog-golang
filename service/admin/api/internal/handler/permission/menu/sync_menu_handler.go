package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/logic/permission/menu"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
)

// 同步菜单列表
func SyncMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SyncMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := menu.NewSyncMenuLogic(r.Context(), svcCtx)
		resp, err := l.SyncMenu(&req)
		responsex.Response(r, w, resp, err)
	}
}
