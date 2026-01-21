package notice

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/logic/notice"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
)

// 查询用户通知列表
func FindUserNoticeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryUserNoticeReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := notice.NewFindUserNoticeListLogic(r.Context(), svcCtx)
		resp, err := l.FindUserNoticeList(&req)
		responsex.Response(r, w, resp, err)
	}
}
