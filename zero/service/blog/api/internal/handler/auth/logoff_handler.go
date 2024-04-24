package auth

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/auth"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
)

// 注销
func LogoffHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewLogoffLogic(r.Context(), svcCtx)
		resp, err := l.Logoff()
		responsex.Response(r, w, resp, err)
	}
}
