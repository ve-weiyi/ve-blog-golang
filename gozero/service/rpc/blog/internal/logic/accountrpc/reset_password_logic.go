package accountrpclogic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/constant"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置密码
func (l *ResetPasswordLogic) ResetPassword(in *accountrpc.ResetPasswordReq) (*accountrpc.EmptyResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, apierr.NewApiError(apierr.CodeInvalidParam, "邮箱格式不正确")
	}

	// 验证用户是否存在
	user, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, apierr.NewApiError(apierr.CodeUserNotExist, err.Error())
	}

	// 验证code是否正确
	key := fmt.Sprintf("%s:%s", constant.ResetPwd, in.Username)
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, in.VerifyCode) {
		return nil, apierr.NewApiError(apierr.CodeCaptchaVerify, "验证码错误")
	}

	// 更新密码
	user.Password = crypto.BcryptHash(in.Password)

	_, err = l.svcCtx.TUserModel.Save(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &accountrpc.EmptyResp{}, nil
}
