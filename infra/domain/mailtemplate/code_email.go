package mailtemplate

import (
	_ "embed"

	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

//go:embed code_email.tpl
var codeEmailTemplate string

// CodeSceneText 返回验证码邮件的标题与引导文案。
func CodeSceneText(scene string) (title, intro string) {
	switch scene {
	case enums.CodeSceneRegister:
		return "注册验证码", "感谢你注册博客，请使用以下验证码完成邮箱验证。"
	case enums.CodeSceneEmailLogin:
		return "登录验证码", "你正在使用邮箱验证码登录博客，请使用以下验证码完成登录。"
	case enums.CodeSceneResetPassword:
		return "重置密码验证码", "你正在重置博客账号密码，请使用以下验证码完成操作。"
	case enums.CodeSceneBind:
		return "绑定验证码", "你正在绑定邮箱到博客账号，请使用以下验证码完成绑定。"
	default:
		return "验证码", "你正在进行身份验证，请使用以下验证码完成操作。"
	}
}

// CodeEmail 渲染验证码邮件的 HTML 内容。
func CodeEmail(scene, code, minutes string) (string, error) {
	title, intro := CodeSceneText(scene)

	return Render(codeEmailTemplate, map[string]string{
		"title": title,
		"intro": intro,
		"code":  code,
		"time":  minutes,
	})
}
