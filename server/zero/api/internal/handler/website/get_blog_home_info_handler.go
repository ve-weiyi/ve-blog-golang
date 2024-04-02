package website

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/website"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func GetBlogHomeInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := website.NewGetBlogHomeInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetBlogHomeInfo()
		responsex.Response(r, w, resp, err)
	}
}
