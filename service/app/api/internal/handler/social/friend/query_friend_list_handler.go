package friend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/infra/responsex"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/logic/social/friend"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/types"
)

// 获取友链列表
func QueryFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryFriendListReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := friend.NewQueryFriendListLogic(r.Context(), svcCtx)
		resp, err := l.QueryFriendList(&req)
		responsex.Response(r, w, resp, err)
	}
}
