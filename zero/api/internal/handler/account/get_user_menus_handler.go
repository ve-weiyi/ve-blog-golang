package account

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/logic/account"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
)

func GetUserMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewGetUserMenusLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMenus()
		responsex.Response(r, w, resp, err)
	}
}
