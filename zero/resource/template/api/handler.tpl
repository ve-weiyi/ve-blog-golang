package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        var reqCtx types.RestHeader
        if err := httpx.ParseHeaders(r, &reqCtx); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}(&reqCtx{{if .HasRequest}}, &req{{end}})
		{{if .HasResp}}responsex.Response(r, w, resp, err){{else}}responsex.Response(r, w, nil, err){{end}}
	}
}