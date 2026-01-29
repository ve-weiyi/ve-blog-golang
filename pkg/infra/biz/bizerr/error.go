package bizerr

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/pkg/infra/biz/bizcode"
)

// BizError 是一个业务错误的结构体
type BizError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"` // 原始错误，不序列化到JSON
}

// Error 返回错误的消息
func (e *BizError) Error() string {
	return e.Message
}

// Details 返回详细的错误信息
func (e *BizError) Details() string {
	if e.Err != nil {
		return fmt.Sprintf("code:%d, message:'%s', err:%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("code:%d, message:'%s'", e.Code, e.Message)
}

// Unwrap 返回原始错误，支持 errors.Is 和 errors.As
func (e *BizError) Unwrap() error {
	return e.Err
}

// NewBizError 创建一个新的业务错误
func NewBizError(code int64, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

// WrapBizError 包装一个错误为业务错误
func WrapBizError(code int64, message string, err error) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// 预定义的常用业务错误
var (
	// 成功
	ErrSuccess = NewBizError(bizcode.CodeSuccess, "操作成功")

	// 3xx 请求频繁
	ErrTooManyRequests = NewBizError(bizcode.CodeTooManyRequests, "操作频繁，请稍后再试")

	// 4xx 请求参数错误
	ErrInvalidParam      = NewBizError(bizcode.CodeInvalidParam, "请求参数错误")
	ErrUserUnLogin       = NewBizError(bizcode.CodeUserUnLogin, "用户未登录")
	ErrUserLoginExpired  = NewBizError(bizcode.CodeUserLoginExpired, "用户登录已过期")
	ErrUserNotPermission = NewBizError(bizcode.CodeUserNotPermission, "用户无权限")
	ErrUserNotExist      = NewBizError(bizcode.CodeUserNotExist, "用户不存在")
	ErrUserAlreadyExist  = NewBizError(bizcode.CodeUserAlreadyExist, "用户已存在")
	ErrUserPasswordError = NewBizError(bizcode.CodeUserPasswordError, "用户密码错误")
	ErrUserDisabled      = NewBizError(bizcode.CodeUserDisabled, "用户已被禁用")

	// 5xx 服务器错误
	ErrInternalServerError = NewBizError(bizcode.CodeInternalServerError, "服务器内部错误")
	ErrSqlQueryError       = NewBizError(bizcode.CodeSqlQueryError, "数据库查询错误")
	ErrCaptchaVerify       = NewBizError(bizcode.CodeCaptchaVerify, "验证码错误")
)
