package request

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/fmtplus"
)

// 验证码生成
type CaptchaReq struct {
	CaptchaType string `json:"captcha_type"`
	Height      int    `json:"height"` // Height png height in pixel.
	Width       int    `json:"width"`  // Width CaptchaReq png width in pixel.
	Length      int    `json:"length"` // DefaultLen Default number of digits in captcha solution.
}

// 验证码验证请求
type CaptchaVerifyReq struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

type CaptchaEmailReq struct {
	Email   string `json:"email"`   // 目标邮箱
	Service string `json:"service"` // 服务
	Check   bool   `json:"check"`   // 是否检查邮箱是否存在
}

func (req *CaptchaEmailReq) IsValid() error {
	// 参数校验
	if !fmtplus.IsEmailValid(req.Email) {
		return codes.NewApiError(codes.CodeInvalidParameter, "邮箱格式不正确")
	}
	if req.Service == "" {
		return codes.NewApiError(codes.CodeInvalidParameter, "服务名不能为空")
	}

	return nil
}
