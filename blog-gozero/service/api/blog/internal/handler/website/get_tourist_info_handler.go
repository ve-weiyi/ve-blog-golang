package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/logic/auth"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
)

// 获取游客身份信息
func GetTouristInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

<<<<<<<< HEAD:blog-gozero/service/api/blog/internal/handler/auth/get_tourist_info_handler.go
		l := auth.NewGetTouristInfoLogic(r.Context(), svcCtx)
========
		l := website.NewGetTouristInfoLogic(r.Context(), svcCtx)
>>>>>>>> 97aca835 (v3.4.1 hotfix (#29)):blog-gozero/service/api/blog/internal/handler/website/get_tourist_info_handler.go
		resp, err := l.GetTouristInfo(&req)
		responsex.Response(r, w, resp, err)
	}
}
