// Package redactx 对日志中的敏感字段做脱敏。
package redactx

import (
	"encoding/json"
	"strings"
)

// credentialKeys 是日志中必须脱敏的凭据字段名（小写比较）。
var credentialKeys = map[string]struct{}{
	"password":         {},
	"old_password":     {},
	"new_password":     {},
	"confirm_password": {},
	"token":            {},
	"access_token":     {},
	"refresh_token":    {},
	"secret":           {},
	"authorization":    {},
}

// verificationKeys 是验证码字段。只用于 RPC 日志：HTTP 统一响应体顶层的
// code 是业务状态（SUCCESS/错误标识），按字段名脱敏会把审计信息一起抹掉。
var verificationKeys = map[string]struct{}{
	"verify_code": {},
	"code":        {},
}

const maskedValue = "***"

// Marshal 序列化任意值并脱敏凭据与验证码字段，用于 RPC 日志。
func Marshal(v any) []byte {
	if v == nil {
		return []byte("null")
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return []byte(`"<marshal error>"`)
	}

	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return raw
	}

	out, err := json.Marshal(redactValue(decoded, true))
	if err != nil {
		return raw
	}
	return out
}

// JSON 脱敏一段 JSON 文本中的凭据字段；非法 JSON 原样返回。
func JSON(raw string) string {
	if raw == "" {
		return raw
	}

	var value any
	if err := json.Unmarshal([]byte(raw), &value); err != nil {
		return raw
	}

	out, err := json.Marshal(redactValue(value, false))
	if err != nil {
		return raw
	}
	return string(out)
}

func redactValue(v any, withVerification bool) any {
	switch value := v.(type) {
	case map[string]any:
		for key, item := range value {
			lower := strings.ToLower(key)
			if _, ok := credentialKeys[lower]; ok {
				value[key] = maskedValue
				continue
			}
			if withVerification {
				if _, ok := verificationKeys[lower]; ok {
					value[key] = maskedValue
					continue
				}
			}
			value[key] = redactValue(item, withVerification)
		}
		return value
	case []any:
		for i, item := range value {
			value[i] = redactValue(item, withVerification)
		}
		return value
	default:
		return v
	}
}
