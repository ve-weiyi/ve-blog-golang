package dto

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
)

type CaptchaDTO struct {
	Id         string `json:"id"`
	EncodeData string `json:"encode_data"` // 验证码内容，base64编码
	Length     int64  `json:"length"`
}

// 验证码生成
type CaptchaReq struct {
	CaptchaType string `json:"captcha_type"`
	Height      int64  `json:"height"` // Height png height in pixel.
	Width       int64  `json:"width"`  // Width CaptchaReq png width in pixel.
	Length      int64  `json:"length"` // DefaultLen Default number of digits in captcha solution.
}

// 验证码验证请求
type CaptchaVerifyReq struct {
	Id   string `json:"id"`
	Code string `json:"code"`
}

type CaptchaEmailReq struct {
	Email   string `json:"email"`   // 目标邮箱
	Service string `json:"service"` // 服务
	Check   bool   `json:"check"`   // 是否检查邮箱是否存在
}

func (req *CaptchaEmailReq) IsValid() error {
	// 参数校验
	if !valid.IsEmailValid(req.Email) {
		return fmt.Errorf("邮箱格式不正确")
	}
	if req.Service == "" {
		return fmt.Errorf("服务不能为空")
	}

	return nil
}
