package entity

import (

)

// register your handlers here
func RegisterHandlers(r *gin.RouterGroup, serverCtx *svc.ServiceContext) {
    {{- range .}}
    // {{.Name}}
    {
        handler := controller.New{{Case2Camel .Name}}Controller(serverCtx)
        {{- range .Routes}}
            {{- range .Doc}}
            // {{.}}
            {{- end}}
            r.{{.Method}}("{{.Path}}", handler.{{.Handler}})
        {{- end}}
    }
    {{- end}}
}

