package {{.Package}}

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	{{- range .Imports}}
    {{.}}
    {{- end}}
)

type {{Case2Camel .Group}}Service struct {
	svcCtx *svctx.ServiceContext
}

func New{{Case2Camel .Group}}Service(svcCtx *svctx.ServiceContext) *{{Case2Camel .Group}}Service {
	return &{{Case2Camel .Group}}Service{
		svcCtx: svcCtx,
	}
}

{{- range .GroupRoutes}}
{{- range .Routes}}
// {{.Doc}}
{{- if .Request }}
func (s *{{Case2Camel $.Group}}Service) {{.Handler}}(reqCtx *request.Context, in *dto.{{.Request}}) (out *dto.{{.Response}}, err error) {
    // todo

    return
}
{{- else }}
func (s *{{Case2Camel $.Group}}Service) {{.Handler}}(reqCtx *request.Context) (err error) {
    // todo

    return
}
{{- end }}

{{- end}}
{{- end}}
