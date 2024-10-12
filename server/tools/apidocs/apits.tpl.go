package apidocs

const ApiTypeScript = `
{{- range .ImportPkgPaths -}}
{{.}}
{{ end -}}
` + ModelImportTpl + `
{{ range .ApiDeclares }}
/** {{ .Summary }} */
export function {{ .FunctionName }}(` + ParamsTpl + `): Promise<{{.Response}}> {
  return http.request<{{.Response}}>({
    url: ` + "`{{.Route}}`" + `,
    method: "{{ .Method }}",
    {{- if .Body }}
    data: {{ .Body.Group }},
    {{- end }}
  })
}
{{ end -}}
`

const ModelTypeScript = `
{{- range . -}}
export interface {{ .Group }} {{ if .Extends }}extends {{ joinArray .Extends }} {{ end }}{
  {{- range .Fields }}
  {{ .Group }}?: {{ .Type }}{{ if .Comment }} // {{ .Comment }}{{ end }}
  {{- end }}
}

{{ end -}}
`
const ParamsTpl = `{{- .Request -}}`

const PathTpl = `{{- if .Path -}}
  {{- range .Path -}}
  {{ .Group }}?: {{ .Type }}
  {{- end -}}
{{- end -}}`

const BodyTpl = `{{- if .Body -}}
  {{ .Body.Group }}?: {{ .Body.Type }}
{{- end -}}`

const ModelImportTpl = `{{ if .ImportModelTypes }}import { {{ joinArray .ImportModelTypes }} } from "./types"{{ end }}`
