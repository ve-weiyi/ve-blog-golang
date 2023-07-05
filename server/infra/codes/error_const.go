package codes

var (
	ErrorSqlQuery         = NewError(CodeSqlQuery, "数据库查询错误")
	ErrorUserUnLogin      = NewError(CodeForbiddenOperation, "用户未登录")
	ErrorCaptchaVerify    = NewError(CodeCaptchaVerifyError, "验证码错误")
	ErrorUserNotExist     = NewError(CodeUserNotExist, "用户不存在")
	ErrorUserAlreadyExist = NewError(CodeUserAlreadyExist, "用户已存在")
)
