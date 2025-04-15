package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/global/mailtemplate"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/tempx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendRegisterEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendRegisterEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendRegisterEmailLogic {
	return &SendRegisterEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送注册邮件
func (l *SendRegisterEmailLogic) SendRegisterEmail(in *accountrpc.UserEmailReq) (*accountrpc.EmptyResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, bizerr.NewBizError(bizerr.CodeInvalidParam, "邮箱格式不正确")
	}

	// 验证用户是否存在
	exist, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if exist != nil {
		return nil, bizerr.NewBizError(bizerr.CodeUserAlreadyExist, "用户已存在")
	}

	// 发送验证码邮件
	key := rediskey.GetCaptchaKey(constant.Register, in.Username)
	code, _ := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key)
	data := mailtemplate.CaptchaEmail{
		Username: in.Username,
		Content:  "欢迎注册我的博客平台。",
		Code:     code,
	}

	// 组装邮件内容
	content, err := tempx.TempParseString(mailtemplate.TempCaptchaCode, data)
	if err != nil {
		return nil, err
	}

	msg := &mail.EmailMessage{
		To:      []string{in.Username},
		Subject: "注册邮件提醒",
		Content: content,
		CC:      false,
	}
	// 发送邮件
	err = l.svcCtx.EmailDeliver.DeliveryEmail(msg)
	if err != nil {
		return nil, err
	}

	return &accountrpc.EmptyResp{}, nil
}
