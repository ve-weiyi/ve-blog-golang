package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/account"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
)

// 修改用户状态
func UpdateUserStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqCtx types.RestHeader
		if err := httpx.ParseHeaders(r, &reqCtx); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		var req types.UpdateUserStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := account.NewUpdateUserStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserStatus(&reqCtx, &req)
		responsex.Response(r, w, resp, err)
	}
}
