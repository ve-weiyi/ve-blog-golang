// 声明 proto 语法版本，固定值
syntax = "proto3";

// proto 包名
package blog;

// 生成 golang 代码后的包名
option go_package = "./blog";

import "any.proto";

{{- range $key, $value := .}}
message {{$value.Name}} {
    {{- range $index, $field := $value.Fields}}
      {{messageType $field.Type}} {{$field.Json}} = {{add $index 1}};{{if $field.Comment}}// {{$field.Comment}}{{end}}
    {{- end}}
}
{{end}}
