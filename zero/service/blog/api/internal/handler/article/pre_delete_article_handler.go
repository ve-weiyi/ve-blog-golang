package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/article"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
)

// 删除文章-逻辑删除
func PreDeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewPreDeleteArticleLogic(r.Context(), svcCtx)
		resp, err := l.PreDeleteArticle(&req)
		responsex.Response(r, w, resp, err)
	}
}
