import request from "@/utils/request";

{{ range .Types }}
{{- .Comment }}
export interface {{ .Name }} {{ if .Extends }}extends {{Join .Extends}} {{ end }}{
  {{- range .Fields }}
  {{ .Name }}{{ if .Nullable }}?{{ end }}: {{ .Type }}; {{ .Comment }}
  {{- end }}
}
{{ end }}

export const {{ .Name }}API = {
{{- range .TsApiGroups }}
{{- $prefix := .Prefix -}}
{{- range .Routes}}
  /** {{ .Summery }} */
  {{ .Handler }}(
  {{- if .Request }}data?: {{ .Request }}{{ end -}}
  ): Promise<IApiResponse<{{.Response}}>> {
    {{- if .PathFields }}
    let url = "{{$prefix}}{{.Path}}";
    {{- range .PathFields }}
    if (data?.{{ . }}) {
      url = url.replace(":{{ . }}", String(data.{{ . }}));
    }
    {{- end }}
    {{- end }}

    {{- if .QueryFields }}
    const params = new URLSearchParams();
    {{- range .QueryFields }}
    if (data?.{{ . }} !== undefined) {
      params.append("{{ . }}", String(data.{{ . }}));
    }
    {{- end }}
    {{- end }}

    {{- if .FormFields }}
    const formData = new FormData();
    {{- range .FormFields }}
    if (data?.{{ . }} !== undefined) {
      formData.append("{{ . }}", data.{{ . }});
    }
    {{- end }}
    {{- end }}
    return request({
      url: {{ if .PathFields }}url{{ else }}"{{$prefix}}{{.Path}}"{{ end }}{{ if .QueryFields }} + "?" + params.toString(){{ end }},
      method: "{{.Method}}",
      {{- if .FormFields }}
      data: formData,
      headers: { "Content-Type": "multipart/form-data" },
      {{- else if .Request }}
      data: data,
      {{- end }}
    });
  },
{{ end -}}
{{ end -}}
};
