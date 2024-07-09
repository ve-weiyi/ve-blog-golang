// 声明 proto 语法版本，固定值
syntax = "proto3";

// proto 包名
package {{.LowerStartCamelName}};

// 生成 golang 代码后的包名
option go_package = "{{.LowerStartCamelName}}";

// {{.CommentName}}结构体
message {{.UpperStartCamelName}} {
{{- range $index, $field := .Fields}}
  {{$field.Type}} {{$field.ColumnName}} = {{add $index 1}};{{if $field.ColumnComment}}// {{$field.ColumnComment}}{{end}}
{{- end}}
}

// {{.CommentName}}服务
service {{.UpperStartCamelName}}Rpc {
  rpc Create{{.UpperStartCamelName}}({{.UpperStartCamelName}}) returns({{.UpperStartCamelName}});

  rpc Update{{.UpperStartCamelName}}({{.UpperStartCamelName}}) returns({{.UpperStartCamelName}});

  rpc Delete{{.UpperStartCamelName}}(IdReq) returns(EmptyResp);

  rpc Find{{.UpperStartCamelName}}(IdReq) returns({{.UpperStartCamelName}});

  rpc Delete{{.UpperStartCamelName}}List(IdsReq) returns(BatchResult);

  rpc Find{{.UpperStartCamelName}}List(PageQuery) returns(PageResult);
}
