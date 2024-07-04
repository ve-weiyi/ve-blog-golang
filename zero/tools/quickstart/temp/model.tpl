package entity

import (

)

// TableName{{.UpperStartCamelName}} return the table name of <{{.TableName}}>
const TableName{{.UpperStartCamelName}} = "{{.TableName}}"

// {{.UpperStartCamelName}} mapped from table <{{.TableName}}>
type {{.UpperStartCamelName}} struct {
    {{range .Fields}}
         {{.Name}} {{.Type}} `{{.Tags}}` {{if .Comment}}// {{.Comment}}{{end}}
    {{- end}}
}

{{ if .TableName }}
// TableName {{.UpperStartCamelName}} 's table name
func (*{{.UpperStartCamelName}}) TableName() string {
  return TableName{{.UpperStartCamelName}}
}
{{ end }}

