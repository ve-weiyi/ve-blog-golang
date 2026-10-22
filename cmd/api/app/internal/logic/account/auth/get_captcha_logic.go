package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/types"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/store/captchax"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取验证码
func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(req *types.GetCaptchaReq) (resp *types.GetCaptchaResp, err error) {
	if req.Width < 0 || req.Height < 0 ||
		(req.Width > 0 && (req.Width < captchax.MinRequestDimension || req.Width > captchax.MaxRequestDimension)) ||
		(req.Height > 0 && (req.Height < captchax.MinRequestDimension || req.Height > captchax.MaxRequestDimension)) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "验证码尺寸不合法")
	}

	id, base64, _, err := l.svcCtx.CaptchaStore.GetMathImageCaptcha(int(req.Height), int(req.Width))
	if err != nil {
		return nil, err
	}

	return &types.GetCaptchaResp{
		CaptchaKey:    id,
		CaptchaBase64: base64,
	}, nil
}
