{{- range . -}}
{{- range .Docs -}}
{{ . }}
{{ end -}}
export interface {{ .Name }} {{convertExtends .Members}}{
  {{- range .Members }}
{{- if .Name }}
  {{ convertJson .Name }}?: {{ convertTsType .Type.RawName }}; {{ .Comment }}
{{- end -}}
  {{- end }}
}

{{ end -}}
