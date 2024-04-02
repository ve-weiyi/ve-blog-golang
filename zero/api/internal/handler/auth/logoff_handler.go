package auth

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/logic/auth"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
)

func LogoffHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewLogoffLogic(r.Context(), svcCtx)
		resp, err := l.Logoff()
		responsex.Response(r, w, resp, err)
	}
}
