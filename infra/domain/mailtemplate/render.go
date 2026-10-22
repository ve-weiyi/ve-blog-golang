package mailtemplate

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var placeholderPattern = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)

// Render 用 params 替换 content 中的 {key} 占位符。
// 有未提供的占位符时返回错误，避免把字面量发给用户。
func Render(content string, params map[string]string) (string, error) {
	missing := make(map[string]struct{})

	out := placeholderPattern.ReplaceAllStringFunc(content, func(placeholder string) string {
		key := placeholder[1 : len(placeholder)-1]
		value, ok := params[key]
		if !ok {
			missing[key] = struct{}{}
			return placeholder
		}
		return value
	})

	if len(missing) == 0 {
		return out, nil
	}

	keys := make([]string, 0, len(missing))
	for key := range missing {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return "", fmt.Errorf("missing template params: %s", strings.Join(keys, ", "))
}
