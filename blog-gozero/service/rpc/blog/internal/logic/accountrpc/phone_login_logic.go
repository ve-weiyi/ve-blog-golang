package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhoneLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneLoginLogic {
	return &PhoneLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 手机号登录
func (l *PhoneLoginLogic) PhoneLogin(in *accountrpc.PhoneLoginReq) (*accountrpc.LoginResp, error) {
	// 校验参数
	if !valid.IsPhoneValid(in.Phone) {
		return nil, bizerr.NewBizError(bizerr.CodeInvalidParam, "手机号格式不正确")
	}

	// 验证用户是否存在
	account, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Phone)
	if err != nil {
		return nil, bizerr.NewBizError(bizerr.CodeUserNotExist, "用户不存在")
	}

	// 验证code是否正确
	key := rediskey.GetCaptchaKey(constant.CodeTypePhoneLogin, in.Phone)
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, in.VerifyCode) {
		return nil, bizerr.NewBizError(bizerr.CodeCaptchaVerify, "验证码错误")
	}

	return onLogin(l.ctx, l.svcCtx, account, constant.LoginTypePhone)
}
