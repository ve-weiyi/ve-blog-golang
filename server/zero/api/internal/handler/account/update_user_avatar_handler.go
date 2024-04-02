package account

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/account"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func UpdateUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewUpdateUserAvatarLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserAvatar()
		responsex.Response(r, w, resp, err)
	}
}
