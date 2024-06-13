{{- range . -}}
{{ .Comment }}
export interface {{ .Name }} {{ if .Extends }}extends {{Join .Extends}} {{ end }}{
  {{- range .Fields }}
  {{ .Name }}?: {{ .Type }}; {{ .Comment }}
  {{- end }}
}
{{ end -}}
