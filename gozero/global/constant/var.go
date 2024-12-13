package constant

// 登录类型
const (
	LoginTypeEmail  = "email" // 邮箱登录
	LoginTypeMobile = "phone" // 手机登录
	LoginTypeOauth  = "oauth" // 第三方登录
)

// 用户状态
const (
	UserStatusDeleted  = -1
	UserStatusNormal   = 0
	UserStatusDisabled = 1
)
