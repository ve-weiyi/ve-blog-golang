package enums

// 验证码场景
const (
	CodeSceneMobileLogin   = "mobile_login"
	CodeSceneEmailLogin    = "email_login"
	CodeSceneRegister      = "register"
	CodeSceneResetPassword = "reset_password"
	CodeSceneBind          = "bind"
)

// IsEmailCodeScene 判断场景是否可用于邮箱验证码。
func IsEmailCodeScene(scene string) bool {
	switch scene {
	case CodeSceneRegister, CodeSceneEmailLogin, CodeSceneResetPassword, CodeSceneBind:
		return true
	default:
		return false
	}
}

// IsMobileCodeScene 判断场景是否可用于手机验证码。
func IsMobileCodeScene(scene string) bool {
	switch scene {
	case CodeSceneMobileLogin, CodeSceneResetPassword, CodeSceneBind:
		return true
	default:
		return false
	}
}
