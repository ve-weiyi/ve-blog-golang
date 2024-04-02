package website

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/website"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func GetAboutMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := website.NewGetAboutMeLogic(r.Context(), svcCtx)
		resp, err := l.GetAboutMe()
		responsex.Response(r, w, resp, err)
	}
}
