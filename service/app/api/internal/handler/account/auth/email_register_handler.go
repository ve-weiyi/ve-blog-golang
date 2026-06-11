package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/logic/account/auth"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/types"
)

// 邮箱注册
func EmailRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := auth.NewEmailRegisterLogic(r.Context(), svcCtx)
		resp, err := l.EmailRegister(&req)
		responsex.Response(r, w, resp, err)
	}
}
