package tag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/logic/tag"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/types"
)

// 删除标签
func DeleteTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := tag.NewDeleteTagLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTag(&req)
		responsex.Response(r, w, resp, err)
	}
}
