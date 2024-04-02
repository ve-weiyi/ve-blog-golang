syntax = "v1"

import "types.api"

@server(
    prefix: /api/v1
    group: {{.Tag}}
)

service blog-api {
    {{- range .ApiDeclares}}
    @doc "{{.Summary}}"
    @handler {{.FunctionName}}
    {{.Method}} {{.Route}} ({{.Request}}) returns ({{.Response}})
    {{end}}
}
