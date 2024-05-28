package service

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/temputil"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
)

type CaptchaService struct {
	svcCtx *svc.ServiceContext
}

func NewCaptchaService(svcCtx *svc.ServiceContext) *CaptchaService {
	return &CaptchaService{
		svcCtx: svcCtx,
	}
}

// 发送验证码
func (l *CaptchaService) SendCaptchaEmail(reqCtx *request.Context, req *request.CaptchaEmailReq) (result interface{}, err error) {
	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountRepository.LoadUserByUsername(reqCtx, req.Email)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	userinfo, err := l.svcCtx.UserAccountRepository.FindUserInfo(reqCtx, account.Id)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	// 设置key
	key := fmt.Sprintf("%s:%s", req.Service, req.Email)
	code := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key)
	data := mail.NewEmailContent()

	data.Title = fmt.Sprintf("重置密码邮件提醒")
	data.DearUser = fmt.Sprintf("你好，%s", userinfo.Nickname)
	data.Content = fmt.Sprintf("你的账号 %s 正在尝试重置密码。验证码为 %s，有效期15分钟！如果您没有尝试重置密码，请忽略此邮件。", req.Email, code)
	data.ButtonTips = fmt.Sprintf("点击重置密码")
	data.ButtonLink = "https://veweiyi.cn/blog/reset_password"

	var temp string
	switch req.Service {
	case constant.ForgetPassword:
		temp = mail.TempForgetPassword
	case constant.Register:
		temp = mail.TempRegister
	default:
		temp = mail.TempSimpleCode
	}

	content, err := temputil.TempParseString(temp, data)
	if err != nil {
		return nil, err
	}

	msg := &mail.EmailMessage{
		To:      []string{req.Email},
		Subject: data.Title,
		Content: content,
		Type:    0,
	}

	err = l.svcCtx.EmailPublisher.PublishMessage([]byte(jsonconv.ObjectToJson(msg)))
	if err != nil {
		return nil, err
	}
	return true, nil
}

// 获取图片验证码
func (l *CaptchaService) GetCaptchaImage(reqCtx *request.Context, req *request.CaptchaReq) (resp *response.CaptchaDTO, err error) {
	id, b64s, err := l.svcCtx.CaptchaHolder.GetImageCaptcha(req.CaptchaType, req.Height, req.Width, req.Length)
	if err != nil {
		return nil, err
	}

	resp = &response.CaptchaDTO{
		Id:         id,
		EncodeData: b64s,
		Length:     req.Length,
	}
	return resp, nil
}

func (l *CaptchaService) VerifyImageCaptcha(reqCtx *request.Context, req *request.CaptchaVerifyReq) (resp interface{}, err error) {
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(req.Id, req.Code) {
		return nil, apierr.ErrorCaptchaVerify
	}

	return resp, nil
}
