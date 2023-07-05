package codes

// 错误码定义规则
// 400-999 	为系统类型错误与具体业务无关
// 1xxx 	用户账号相关错误码
// 2xxx		会议室和设备相关错误码
// 3xxx		组织管理错误码
const (
	CodeInvalidParam       = 400
	CodeNoPerPermission    = 401 // 权限不足
	CodeInvalidRequest     = 404 // 非法请求
	CodeInvalidToken       = 406 // 无效的token
	CodeInternalError      = 500 // 内部错误
	CodeRequestToManyTimes = 601 // 操作频繁

	CodeMissingParameter = 404 //缺少参数
	CodeInvalidParameter = 404 //无效参数

	CodeRoleNoPerPermission = 403 //角色权限不足
	CodeForbiddenOperation  = 403 // 禁止操作
	CodeUserNotExist        = 404 // 用户不存在
	CodeUserAlreadyExist    = 405 // 用户已存在
	CodeSqlQuery            = 500 //数据库查询错误
	CodeCaptchaVerifyError  = 501 // 验证码错误
)
