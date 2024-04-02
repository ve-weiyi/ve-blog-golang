package upload

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/upload"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upload.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile()
		responsex.Response(r, w, resp, err)
	}
}
