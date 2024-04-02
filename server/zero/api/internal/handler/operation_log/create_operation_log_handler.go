package operation_log

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/operation_log"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"
)

func CreateOperationLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OperationLog
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := operation_log.NewCreateOperationLogLogic(r.Context(), svcCtx)
		resp, err := l.CreateOperationLog(&req)
		responsex.Response(r, w, resp, err)
	}
}
