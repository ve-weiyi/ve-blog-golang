package types

import (

)

{{- range .}}
{{range .Comments}}
     {{.}}
{{- end}}
type {{.RawName}} struct {
    {{range .Members}}
         {{.Name}} {{.Type.RawName}} {{.Tag}} {{if .Comment}}{{.Comment}}{{end}}
    {{- end}}
}
{{- end}}

