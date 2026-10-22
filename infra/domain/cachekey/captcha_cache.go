package cachekey

import "fmt"

// 验证码 redis类型：string
func GetCaptchaKey(module string, username string) string {
	return fmt.Sprintf("blog:captcha:%s:%s", module, username)
}

const (
	// 验证码存储前缀
	CaptchaStorePrefixApp   = "blog:app:captcha:"
	CaptchaStorePrefixAdmin = "blog:admin:captcha:"
)
