package api

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/logic/api"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
)

func CleanApiListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := api.NewCleanApiListLogic(r.Context(), svcCtx)
		resp, err := l.CleanApiList()
		responsex.Response(r, w, resp, err)
	}
}
