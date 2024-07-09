syntax = "v1"

type (
    {{- range $key, $value := .}}
    {{$value.Name}} {
        {{- range $value.Fields}}
        {{.Name}} {{messageType .Type}} `json:"{{.Json}},optional"` {{if .Comment}}// {{.Comment}}{{end}}
    	{{- end}}
    }
    {{end}}
)

