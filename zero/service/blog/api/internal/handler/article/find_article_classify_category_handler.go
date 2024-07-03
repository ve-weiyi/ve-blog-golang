package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/article"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
)

// 通过标签或者id获取文章列表
func FindArticleClassifyCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleClassifyReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := article.NewFindArticleClassifyCategoryLogic(r.Context(), svcCtx)
		resp, err := l.FindArticleClassifyCategory(&req)
		responsex.Response(r, w, resp, err)
	}
}
