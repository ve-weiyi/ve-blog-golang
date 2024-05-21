package authrpclogic

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/temputil"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterEmailLogic {
	return &RegisterEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送注册邮件
func (l *RegisterEmailLogic) RegisterEmail(in *blog.UserEmailReq) (*blog.EmptyResp, error) {
	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, in.Username)
	if account != nil {
		return nil, apierr.ErrorUserAlreadyExist
	}

	// 发送验证码邮件
	key := fmt.Sprintf("%s:%s", constant.Register, in.Username)
	code := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key)
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

	return &blog.EmptyResp{}, nil
}
