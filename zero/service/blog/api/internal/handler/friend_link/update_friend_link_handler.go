package friend_link

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/friend_link"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
)

// 更新友链
func UpdateFriendLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqCtx types.RestHeader
		if err := httpx.ParseHeaders(r, &reqCtx); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var req types.FriendLink
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := friend_link.NewUpdateFriendLinkLogic(r.Context(), svcCtx)
		resp, err := l.UpdateFriendLink(&reqCtx, &req)
		responsex.Response(r, w, resp, err)
	}
}
