package enums

// 用户状态 (t_user.status)
const (
	UserStatusDeleted  = -1 // 已删除
	UserStatusNormal   = 0  // 正常
	UserStatusDisabled = 1  // 禁用
	UserStatusLogoff   = 2  // 注销
)

// IsValidUserStatus 判断状态是否可由管理端设置（已删除仅内部使用）。
func IsValidUserStatus(status int64) bool {
	switch status {
	case UserStatusNormal, UserStatusDisabled, UserStatusLogoff:
		return true
	default:
		return false
	}
}

// 登录类型 (t_login_log.login_type)
const (
	LoginTypeRegister = "register" // 用户注册
	LoginTypeUsername = "username" // 用户名登录
	LoginTypeEmail    = "email"    // 邮箱登录
	LoginTypeMobile   = "mobile"   // 手机登录
	LoginTypeOauth    = "oauth"    // 第三方登录
)
