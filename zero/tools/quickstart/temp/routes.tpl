package {{ .Package }}

import (
{{- range .Imports }}
    {{ . }}
{{- end }}
)

// register your handlers here
func RegisterHandlers(r *gin.RouterGroup, serverCtx *svc.ServiceContext) {
    {{- range .Groups}}
    // {{.Name}}
    {
        handler := controller.New{{Case2Camel .Name}}Controller(serverCtx)
        {{- range .Routes}}
            // {{.Doc}}
            r.{{.Method}}("{{.Path}}", handler.{{.Handler}})
        {{- end}}
    }
    {{- end}}
}

