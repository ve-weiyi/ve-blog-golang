{{- range .ImportPkgPaths -}}
{{.}}
{{ end -}}

{{ if .ImportTypes }}import { {{ Join .ImportTypes }} } from "./types"{{ end }}

{{ range .Routes -}}
/** {{ .Summery }} */
export function {{ .Handler }}(

{{- if .Request }}data?: {{ .Request }}{{ end -}}

): Promise<IApiResponseData<{{.Response}}>> {
  return request({
    url: "{{.Path}}",
    method: "{{.Method}}",
    {{ if .Request }}data: data,{{ end }}
  })
}
{{ end -}}
