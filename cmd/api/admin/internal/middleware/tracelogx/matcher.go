package tracelogx

import (
	"strings"

	"github.com/ve-weiyi/blog-cloud/infra/http/routex"
)

func normalizeMethod(method string) string {
	method = strings.TrimSpace(method)
	if method == "" {
		return ""
	}
	method = strings.ToUpper(method)
	if method == "*" || method == "ALL" {
		return "*"
	}
	return method
}

func normalizePath(path string) string {
	return routex.Normalize(path)
}
