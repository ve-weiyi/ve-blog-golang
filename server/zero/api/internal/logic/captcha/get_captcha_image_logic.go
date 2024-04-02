package captcha

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaImageLogic {
	return &GetCaptchaImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaImageLogic) GetCaptchaImage(req *types.CaptchaReq) (resp *types.CaptchaDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
