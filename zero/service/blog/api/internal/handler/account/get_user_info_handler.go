<<<<<<<< HEAD:zero/service/api/blog/internal/handler/user/get_user_info_handler.go
package user
========
package account
>>>>>>>> 09fef341 (v2.0.0 修改项目结构 (#17)):zero/service/blog/api/internal/handler/account/get_user_info_handler.go

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
<<<<<<<< HEAD:zero/service/api/blog/internal/handler/user/get_user_info_handler.go
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/logic/user"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
========
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/account"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
>>>>>>>> 09fef341 (v2.0.0 修改项目结构 (#17)):zero/service/blog/api/internal/handler/account/get_user_info_handler.go
)

// 获取用户信息
func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

<<<<<<<< HEAD:zero/service/api/blog/internal/handler/user/get_user_info_handler.go
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
========
		l := account.NewGetUserInfoLogic(r.Context(), svcCtx)
>>>>>>>> 09fef341 (v2.0.0 修改项目结构 (#17)):zero/service/blog/api/internal/handler/account/get_user_info_handler.go
		resp, err := l.GetUserInfo(&req)
		responsex.Response(r, w, resp, err)
	}
}
