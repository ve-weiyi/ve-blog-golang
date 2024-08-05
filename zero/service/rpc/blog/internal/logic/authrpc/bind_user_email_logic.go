package authrpclogic

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindUserEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUserEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserEmailLogic {
	return &BindUserEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户邮箱
func (l *BindUserEmailLogic) BindUserEmail(in *blog.BindUserEmailReq) (*blog.EmptyResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Email) {
		return nil, apierr.ErrorInvalidParam
	}

	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	// 验证code是否正确
	key := fmt.Sprintf("%s:%s", constant.BindEmail, in.Email)
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, in.VerifyCode) {
		return nil, apierr.ErrorCaptchaVerify
	}

	// 更新密码
	account.Username = in.Email
	account.Email = in.Email

	_, err = l.svcCtx.UserAccountModel.Update(l.ctx, account)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
