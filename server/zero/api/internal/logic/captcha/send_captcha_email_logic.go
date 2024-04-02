package captcha

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCaptchaEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCaptchaEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCaptchaEmailLogic {
	return &SendCaptchaEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCaptchaEmailLogic) SendCaptchaEmail(req *types.CaptchaEmailReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
