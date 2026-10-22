package upload

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/logic/common/upload"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/types"
	"github.com/ve-weiyi/blog-cloud/infra/http/responsex"
)

// 上传文件（服务端上传）
func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.RequestInvalid(r, w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			responsex.RequestInvalid(r, w, err)
			return
		}
		defer file.Close()

		l := upload.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(&req, file, header.Filename, header.Size)
		responsex.Response(r, w, resp, err)
	}
}
