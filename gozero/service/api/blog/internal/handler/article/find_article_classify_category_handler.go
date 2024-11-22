package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/logic/article"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
)

// 通过分类获取文章列表
func FindArticleClassifyCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleClassifyQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := article.NewFindArticleClassifyCategoryLogic(r.Context(), svcCtx)
		resp, err := l.FindArticleClassifyCategory(&req)
		responsex.Response(r, w, resp, err)
	}
}