package service

import (

)

type {{Case2Camel .Name}}Service struct {
	svcCtx *svc.ServiceContext
}

func New{{Case2Camel .Name}}Service(svcCtx *svc.ServiceContext) *{{Case2Camel .Name}}Service {
	return &{{Case2Camel .Name}}Service{
		svcCtx: svcCtx,
	}
}

{{- range .Routes}}

// {{.Doc}}
{{ if .Request }}
func (s *{{Case2Camel $.Name}}Service) {{.Handler}}(reqCtx *request.Context, in *types.{{.Request}}) (out *types.{{.Response}}, err error) {
    // todo

    return
}
{{ else }}
func (s *{{Case2Camel $.Name}}Service) {{.Handler}}(reqCtx *request.Context) (err error) {
    // todo

    return
}
{{ end }}

{{- end}}
