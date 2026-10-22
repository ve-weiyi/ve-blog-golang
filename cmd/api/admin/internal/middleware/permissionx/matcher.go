package permissionx

import (
	"strings"

	"github.com/ve-weiyi/blog-cloud/infra/http/routex"
)

// 规范化 HTTP 方法（大小写/通配符）
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

// 规范化请求路径
func normalizePath(path string) string {
	return routex.Normalize(path)
}

// matchPermission 支持通配符路径匹配（策略中的 pattern 匹配请求中的 target）
func matchPermission(pattern, target string) bool {
	pm, pp, ok1 := splitPermID(pattern)
	tm, tp, ok2 := splitPermID(target)
	if !ok1 || !ok2 {
		return false
	}
	if pm != "*" && pm != tm {
		return false
	}
	return pathMatch(pp, tp)
}

// splitPermID 将 "METHOD:PATH" 拆分为 (method, path, ok)
func splitPermID(id string) (method, path string, ok bool) {
	idx := strings.IndexByte(id, ':')
	if idx < 0 {
		return "", "", false
	}
	return id[:idx], id[idx+1:], true
}

// isRegistered 判断请求路径是否命中已登记的接口模式。
func isRegistered(registeredByMethod map[string][]string, method, path string) bool {
	for _, pattern := range registeredByMethod[method] {
		if pathMatch(pattern, path) {
			return true
		}
	}
	for _, pattern := range registeredByMethod["*"] {
		if pathMatch(pattern, path) {
			return true
		}
	}
	return false
}

// pathMatch 支持 :param/{param} 与通配符。权限策略里的 * 表示匹配剩余所有段。
func pathMatch(pattern, path string) bool {
	patterns := splitSegments(routex.Normalize(pattern))
	segments := splitSegments(routex.Normalize(path))

	i, j := 0, 0
	for i < len(patterns) && j < len(segments) {
		segment := patterns[i]
		if segment == "*" {
			return true
		}
		if isParam(segment) {
			i++
			j++
			continue
		}
		if segment != segments[j] {
			return false
		}
		i++
		j++
	}

	return i == len(patterns) && j == len(segments)
}

func splitSegments(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return nil
	}
	return strings.Split(path, "/")
}

func isParam(segment string) bool {
	return strings.HasPrefix(segment, ":") && len(segment) > 1
}
