package apidocs

const ApiTypeScript = `
{{- range .ImportPkgPaths}}
{{.}}
{{- end}}
{{range .ApiDeclares}}
/** {{ .Summary }} */
export function {{ .FunctionName }}(` + PathTpl + BodyTpl + `): Promise<{{.Response}}> {
	return http.request<{{.Response}}>({
    	url: ` + "`/api/v1/{{.Url}}`" + `,
		method: "{{ .Method }}",
		{{ if .Body }}data: {{ .Body.Name }},{{ end }}
	})
}
{{end}}
`

const ModelTypeScript = `
{{range $model := .}}
interface {{ $model.Name }} {
    {{- range $field := $model.Fields }}
    {{ $field.Name }}: {{ $field.Type }} {{ if $field.Comment }}// {{ $field.Comment }}{{ end }}
    {{- end }}
}
{{end}}
`

const PathTpl = `{{- if .Path -}}
	{{- range .Path -}}
	{{ .Name }}?: {{ .Type }}
	{{- end -}}
{{- end -}}`

const BodyTpl = `{{- if .Body -}}
	{{ .Body.Name }}?: {{ .Body.Type }}
{{- end -}}`
