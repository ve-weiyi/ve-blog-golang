package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/logic/auth"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
)

// 绑定邮箱
func BindUserEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BindUserEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := auth.NewBindUserEmailLogic(r.Context(), svcCtx)
		resp, err := l.BindUserEmail(&req)
		responsex.Response(r, w, resp, err)
	}
}
