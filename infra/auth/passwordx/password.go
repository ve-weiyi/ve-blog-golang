package passwordx

import (
	"strings"

	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

const (
	MinPasswordLength = 6
	MaxPasswordLength = 64
)

// ValidatePassword 校验密码的基本强度要求。
func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return bizerr.NewBizError(bizcode.CodeInvalidParam, "密码不能为空")
	}
	if len(password) < MinPasswordLength {
		return bizerr.NewBizError(bizcode.CodeInvalidParam, "密码不能少于6位")
	}
	if len(password) > MaxPasswordLength {
		return bizerr.NewBizError(bizcode.CodeInvalidParam, "密码不能超过64位")
	}
	return nil
}

// ValidatePasswordConfirm 校验两次输入的密码是否一致。
func ValidatePasswordConfirm(password, confirm string) error {
	if password != confirm {
		return bizerr.NewBizError(bizcode.CodeInvalidParam, "两次密码输入不一致")
	}
	return nil
}
