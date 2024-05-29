package apierr

import "fmt"

// ApiError 是一个API错误的结构体
type ApiError struct {
	Code    int
	Message string
	Data    interface{}
}

// Error 返回错误的消息
func (e *ApiError) Error() string {
	return e.Message
}

func (e *ApiError) Details() string {
	return fmt.Sprintf("code:%d, message:'%s'", e.Code, e.Message)
}

func (e *ApiError) WrapError(err error) *ApiError {
	ne := &ApiError{
		Code:    e.Code,
		Message: fmt.Sprintf("%v, %v", e.Message, err.Error()),
		Data:    nil,
	}
	return ne
}

func (e *ApiError) WrapMessage(err string) *ApiError {
	ne := &ApiError{
		Code:    e.Code,
		Message: fmt.Sprintf("%v, %v", e.Message, err),
		Data:    nil,
	}
	return ne
}

// NewApiError 创建一个新的API错误
func NewApiError(code int, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
