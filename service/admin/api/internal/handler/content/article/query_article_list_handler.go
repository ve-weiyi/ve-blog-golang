package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/logic/content/article"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
)

// 获取文章列表
func QueryArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryArticleListReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := article.NewQueryArticleListLogic(r.Context(), svcCtx)
		resp, err := l.QueryArticleList(&req)
		responsex.Response(r, w, resp, err)
	}
}
