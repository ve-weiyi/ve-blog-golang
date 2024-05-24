package friend_link

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/friend_link"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
)

// 创建友链
func CreateFriendLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqCtx types.RestHeader
		if err := httpx.ParseHeaders(r, &reqCtx); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		var req types.FriendLink
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := friend_link.NewCreateFriendLinkLogic(r.Context(), svcCtx)
		resp, err := l.CreateFriendLink(&reqCtx, &req)
		responsex.Response(r, w, resp, err)
	}
}
