package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/logic/account"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/types"
)

// 修改用户角色
func UpdateAccountRolesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAccountRolesReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := account.NewUpdateAccountRolesLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAccountRoles(&req)
		responsex.Response(r, w, resp, err)
	}
}
