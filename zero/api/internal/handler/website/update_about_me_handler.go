package website

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/logic/website"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
)

func UpdateAboutMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AboutMe
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := website.NewUpdateAboutMeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAboutMe(&req)
		responsex.Response(r, w, resp, err)
	}
}
