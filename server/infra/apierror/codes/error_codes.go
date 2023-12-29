package codes

// 错误码定义规则
const (
	CodeInvalidParam       = 400
	CodeTokenInvalid       = 401 // 未授权,一般是token过期 http.StatusUnauthorized
	CodeInvalidRequest     = 404 // 非法请求
	CodeInvalidToken       = 406 // 无效的token
	CodeInternalError      = 500 // 内部错误
	CodeRequestToManyTimes = 601 // 操作频繁

	CodeMissingParameter = 400 // 缺少参数
	CodeInvalidParameter = 400 // 无效参数

	CodeUserNotExist       = 404 // 用户不存在
	CodeUserAlreadyExist   = 405 // 用户已存在
	CodeSqlQuery           = 500 //数据库查询错误
	CodeCaptchaVerifyError = 501 // 验证码错误
)
