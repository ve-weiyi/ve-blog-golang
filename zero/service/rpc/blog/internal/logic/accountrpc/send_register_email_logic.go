package accountrpclogic

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr/codex"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/temputil"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

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
		return nil, apierr.NewApiError(codex.CodeInvalidParam, "邮箱格式不正确")
	}

	// 验证用户是否存在
	exist, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, in.Username)
	if exist != nil {
		return nil, apierr.NewApiError(codex.CodeUserAlreadyExist, "用户已存在")
	}

	// 发送验证码邮件
	key := fmt.Sprintf("%s:%s", constant.Register, in.Username)
	code, _ := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key)
	data := mail.CaptchaEmail{
		Username: in.Username,
		Code:     code,
	}

	// 组装邮件内容
	content, err := temputil.TempParseString(mail.TempRegister, data)
	if err != nil {
		return nil, err
	}

	msg := &mail.EmailMessage{
		To:      []string{in.Username},
		Subject: "注册邮件提醒",
		Content: content,
		Type:    0,
	}
	// 发送邮件
	err = l.svcCtx.EmailMQ.PublishMessage([]byte(jsonconv.ObjectToJson(msg)))
	if err != nil {
		return nil, err
	}

	return &accountrpc.EmptyResp{}, nil
}
