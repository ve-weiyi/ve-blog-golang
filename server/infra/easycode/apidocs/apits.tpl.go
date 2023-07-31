package apidocs

const ApiTypeScript = `
{{- range .ImportPkgPaths -}}
{{.}}
{{ end }}

{{- range .ModelDeclares }}
interface {{ .Name }} {{ if .Extend }}extends {{ .Extend.Name }} {{ end }}{
  {{- range .Fields }}
  {{ .Name }}: {{ .Type }}{{ if .Comment }} // {{ .Comment }}{{ end }}
  {{- end }}
}
{{ end }}

{{- range .ApiDeclares }}
/** {{ .Summary }} */
export function {{ .FunctionName }}(` + ParamsTpl + `): Promise<{{.Response}}> {
  return http.request<{{.Response}}>({
    url: ` + "`/api/v1/{{.Url}}`" + `,
    method: "{{ .Method }}",
    {{- if .Body }}
    data: {{ .Body.Name }},
    {{- end }}
  })
}
{{ end -}}
`

const ModelTypeScript = `
{{- range .}}
interface {{ .Name }} {{ if .Extend }}extends {{ .Extend.Name }} {{ end }}{
  {{- range .Fields }}
    {{ .Name }}: {{ .Type }}{{ if .Comment }} // {{ .Comment }}{{ end }}
    {{- end }}
}
{{ end }}
`
const ParamsTpl = `{{- .Request -}}`

const PathTpl = `{{- if .Path -}}
  {{- range .Path -}}
  {{ .Name }}?: {{ .Type }}
  {{- end -}}
{{- end -}}`

const BodyTpl = `{{- if .Body -}}
  {{ .Body.Name }}?: {{ .Body.Type }}
{{- end -}}`
