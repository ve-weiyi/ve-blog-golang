package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/infra/constants/enums"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/notificationservice"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/userauthservice"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 重置密码
func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) (resp *types.ResetPasswordResp, err error) {
	// 验证邮箱验证码（API 层负责）
	_, err = l.svcCtx.NotificationService.VerifyEmailCode(l.ctx, &notificationservice.VerifyEmailCodeRequest{
		Email: req.Email,
		Scene: enums.CodeSceneRestPassword,
		Code:  req.Code,
		BizId: "", // 空值，使用默认规则 scene:email
	})
	if err != nil {
		return nil, err
	}

	// 调用 RPC 重置密码
	in := userauthservice.ResetPasswordRequest{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	_, err = l.svcCtx.UserAuthService.ResetPassword(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.ResetPasswordResp{}, nil
}
