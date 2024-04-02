syntax = "v1"

{{range .ImportPkgPaths -}}
import "{{.}}"
{{- end}}

type (
    {{.UpperStartCamelName}}Req {
        {{- range .Fields}}
        {{.Name}} {{.Type}} `{{.Tags}}` {{if .ColumnComment}}// {{.ColumnComment}}{{end}}
    	{{- end}}
    }

    {{.UpperStartCamelName}}Resp {
        {{- range .Fields}}
        {{.Name}} {{.Type}} `{{.Tags}}` {{if .ColumnComment}}// {{.ColumnComment}}{{end}}
    	{{- end}}
    }

)

@server(
    prefix: /api/v1
    group: {{.LowerStartCamelName}}
)

service blog-api {
    @doc "创建{{.CommentName}}"
    @handler Create{{.UpperStartCamelName}}
    post /{{.LowerStartCamelName}}/create ({{.UpperStartCamelName}}Req) returns ({{.UpperStartCamelName}}Resp)

    @doc "更新{{.CommentName}}"
    @handler Update{{.UpperStartCamelName}}
    put /{{.LowerStartCamelName}}/update ({{.UpperStartCamelName}}Req) returns ({{.UpperStartCamelName}}Resp)

    @doc "删除{{.CommentName}}"
    @handler Delete{{.UpperStartCamelName}}
    delete /{{.LowerStartCamelName}}/delete (IdReq) returns (EmptyResp)

    @doc "查找{{.CommentName}}"
    @handler Find{{.UpperStartCamelName}}
    post /{{.LowerStartCamelName}}/find (IdReq) returns ({{.UpperStartCamelName}}Resp)

    @doc "删除{{.CommentName}}列表"
    @handler Delete{{.UpperStartCamelName}}List
    delete /{{.LowerStartCamelName}}/batch_delete (IdsReq) returns (BatchResult)

    @doc "查找{{.CommentName}}列表"
    @handler Find{{.UpperStartCamelName}}List
    post /{{.LowerStartCamelName}}/list (PageQuery) returns (PageResult)
}
