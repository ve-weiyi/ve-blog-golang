package apierror

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/apierror/codes"
)

var (
	ErrorSqlQuery         = NewApiError(codes.CodeSqlQuery, "数据库查询错误")
	ErrorUserUnLogin      = NewApiError(codes.CodeUnauthorized, "用户未登录")
	ErrorCaptchaVerify    = NewApiError(codes.CodeCaptchaVerifyError, "验证码错误")
	ErrorUserNotExist     = NewApiError(codes.CodeUserNotExist, "用户不存在")
	ErrorUserAlreadyExist = NewApiError(codes.CodeUserAlreadyExist, "用户已存在")
)
