package photo_album

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/logic/photo_album"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
)

// 获取相册详情列表
func FindPhotoAlbumDetailsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqCtx types.RestHeader
		if err := httpx.ParseHeaders(r, &reqCtx); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var req types.PageQuery
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := photo_album.NewFindPhotoAlbumDetailsListLogic(r.Context(), svcCtx)
		resp, err := l.FindPhotoAlbumDetailsList(&reqCtx, &req)
		responsex.Response(r, w, resp, err)
	}
}
