package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/infra/constants/enums"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/notificationservice"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/userauthservice"
)

type MobileLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 手机验证码登录（自动注册）
func NewMobileLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileLoginLogic {
	return &MobileLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MobileLoginLogic) MobileLogin(req *types.MobileLoginReq) (resp *types.LoginResp, err error) {
	_, err = l.svcCtx.NotificationService.VerifyPhoneCode(l.ctx, &notificationservice.VerifyPhoneCodeRequest{
		Phone: req.Mobile,
		Scene: enums.CodeSceneMobileLogin,
		Code:  req.Code,
		BizId: "",
	})
	if err != nil {
		return nil, err
	}

	out, err := l.svcCtx.UserAuthService.LoginByMobile(l.ctx, &userauthservice.LoginByMobileRequest{
		Mobile: req.Mobile,
	})
	if err != nil {
		return nil, err
	}

	return onLogin(l.ctx, l.svcCtx, out)
}
