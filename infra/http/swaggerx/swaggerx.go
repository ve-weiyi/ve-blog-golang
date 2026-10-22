package swaggerx

import (
	"sort"

	"github.com/go-openapi/spec"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodPatch  = "PATCH"
	MethodDelete = "DELETE"
)

// Operation 是 Swagger 中一条可登记为接口的记录。
type Operation struct {
	Path   string
	Method string
	Name   string
	Tags   []string
}

// apiOperations 声明需要展开的方法，顺序即同一路径下的产出顺序。
var apiOperations = []struct {
	method string
	pick   func(spec.PathItem) *spec.Operation
}{
	{MethodGet, func(item spec.PathItem) *spec.Operation { return item.Get }},
	{MethodPost, func(item spec.PathItem) *spec.Operation { return item.Post }},
	{MethodPut, func(item spec.PathItem) *spec.Operation { return item.Put }},
	{MethodPatch, func(item spec.PathItem) *spec.Operation { return item.Patch }},
	{MethodDelete, func(item spec.PathItem) *spec.Operation { return item.Delete }},
}

// Operations 展开 Swagger 操作，按路径排序保证重复同步的结果稳定。
func Operations(sp *spec.Swagger) []Operation {
	paths := make([]string, 0, len(sp.Paths.Paths))
	for path := range sp.Paths.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	var operations []Operation
	for _, path := range paths {
		item := sp.Paths.Paths[path]
		for _, operation := range apiOperations {
			if op := operation.pick(item); op != nil {
				operations = append(operations, Operation{
					Path:   path,
					Method: operation.method,
					Name:   op.Summary,
					Tags:   op.Tags,
				})
			}
		}
	}

	return operations
}
