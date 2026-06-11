package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/logic/content/article"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/types"
)

// 点赞文章
func LikeArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := article.NewLikeArticleLogic(r.Context(), svcCtx)
		resp, err := l.LikeArticle(&req)
		responsex.Response(r, w, resp, err)
	}
}
