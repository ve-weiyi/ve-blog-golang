package upload

import (
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/logic/upload"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
)

func UploadVoiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upload.NewUploadVoiceLogic(r.Context(), svcCtx)
		resp, err := l.UploadVoice()
		responsex.Response(r, w, resp, err)
	}
}
