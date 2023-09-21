package codes

import "fmt"

type ApiError struct {
	code   int
	errMsg string
}

func NewApiError(code int, msg string) *ApiError {
	return &ApiError{
		code:   code,
		errMsg: msg,
	}
}

func (sel *ApiError) Wrap(err error) *ApiError {
	ne := &ApiError{errMsg: fmt.Sprintf("%v,err:%v", sel.errMsg, err), code: sel.code}
	return ne
}

func (sel *ApiError) IsErr(err error) bool {
	if e, ok := err.(*ApiError); ok {
		return e.code == sel.code
	}
	return false
}

func (sel *ApiError) IsCode(c int) bool {
	return sel.code == c
}

func (sel *ApiError) Code() int {
	return sel.code
}

func (sel *ApiError) Message() string {
	return string(sel.errMsg)
}

func (sel *ApiError) Error() string {
	return sel.errMsg
}
