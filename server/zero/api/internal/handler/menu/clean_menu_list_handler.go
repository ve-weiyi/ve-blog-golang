package menu

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/menu"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func CleanMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewCleanMenuListLogic(r.Context(), svcCtx)
		resp, err := l.CleanMenuList()
		responsex.Response(r, w, resp, err)
	}
}
