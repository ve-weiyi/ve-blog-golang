package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/notificationservice"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

type SendEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送邮箱验证码
func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(req *types.SendEmailCodeReq) (resp *types.SendEmailCodeResp, err error) {
	if !enums.IsEmailCodeScene(req.Type) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "验证码场景不合法")
	}

	_, err = l.svcCtx.NotificationService.SendEmailCode(l.ctx, &notificationservice.SendEmailCodeRequest{
		Email: req.Email,
		Scene: req.Type,
		BizId: "",
	})
	if err != nil {
		return nil, err
	}

	return &types.SendEmailCodeResp{}, nil
}
