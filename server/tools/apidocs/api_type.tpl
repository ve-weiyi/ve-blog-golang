syntax = "v1"

type (
    {{- range $key, $value := .}}
    {{$value.Name}} {
        {{- range $value.Fields}}
        {{.Name}} {{.Type}} `json:"{{.Json}}"` {{if .Comment}}// {{.Comment}}{{end}}
    	{{- end}}
    }
    {{end}}
)

