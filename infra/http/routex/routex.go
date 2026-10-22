// Package routex 提供路由路径的规范化与匹配。
//
// 规则：:param 与 {param} 匹配任意单段；* 匹配单段；** 匹配零到多段。
// 参数两种写法在 Normalize 后统一为 :param，供路由、Swagger 与权限策略共用。
package routex

import "strings"

// Normalize 去掉 query 与尾随斜杠、补齐前导斜杠，并把 {param} 统一成 :param。
func Normalize(path string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return ""
	}
	if idx := strings.IndexByte(path, '?'); idx >= 0 {
		path = path[:idx]
	}
	if path != "/" {
		path = strings.TrimSuffix(path, "/")
	}
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	segments := split(path)
	for i, segment := range segments {
		if isBraceParam(segment) {
			segments[i] = ":" + segment[1:len(segment)-1]
		}
	}
	return "/" + strings.Join(segments, "/")
}

// Match 判断具体路径是否匹配模式。
func Match(pattern, path string) bool {
	patterns := split(Normalize(pattern))
	segments := split(Normalize(path))

	var match func(patternIndex, segmentIndex int) bool
	match = func(patternIndex, segmentIndex int) bool {
		if patternIndex == len(patterns) {
			return segmentIndex == len(segments)
		}

		if patterns[patternIndex] == "**" {
			if patternIndex == len(patterns)-1 {
				return true
			}
			for i := segmentIndex; i <= len(segments); i++ {
				if match(patternIndex+1, i) {
					return true
				}
			}
			return false
		}

		if segmentIndex == len(segments) {
			return false
		}

		segment := patterns[patternIndex]
		if segment == "*" || isParam(segment) || segment == segments[segmentIndex] {
			return match(patternIndex+1, segmentIndex+1)
		}
		return false
	}

	return match(0, 0)
}

func split(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return nil
	}
	return strings.Split(path, "/")
}

func isParam(segment string) bool {
	return strings.HasPrefix(segment, ":") && len(segment) > 1
}

func isBraceParam(segment string) bool {
	return strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") && len(segment) > 2
}
