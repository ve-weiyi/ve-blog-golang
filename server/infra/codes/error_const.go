package codes

var (
	ErrorSqlQuery         = NewApiError(CodeSqlQuery, "数据库查询错误")
	ErrorUserUnLogin      = NewApiError(CodeForbiddenOperation, "用户未登录")
	ErrorCaptchaVerify    = NewApiError(CodeCaptchaVerifyError, "验证码错误")
	ErrorUserNotExist     = NewApiError(CodeUserNotExist, "用户不存在")
	ErrorUserAlreadyExist = NewApiError(CodeUserAlreadyExist, "用户已存在")
)
