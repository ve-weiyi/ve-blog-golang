package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/logic/discussion/comment"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/infra/http/responsex"
)

// 批量审核评论
func BatchReviewCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchReviewCommentsReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := comment.NewBatchReviewCommentsLogic(r.Context(), svcCtx)
		resp, err := l.BatchReviewComments(&req)
		responsex.Response(r, w, resp, err)
	}
}
