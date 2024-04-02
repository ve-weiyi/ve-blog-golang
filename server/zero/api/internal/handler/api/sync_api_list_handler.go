package api

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/api"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func SyncApiListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := api.NewSyncApiListLogic(r.Context(), svcCtx)
		resp, err := l.SyncApiList()
		responsex.Response(r, w, resp, err)
	}
}
