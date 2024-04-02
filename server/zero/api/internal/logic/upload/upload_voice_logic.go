package upload

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVoiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadVoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVoiceLogic {
	return &UploadVoiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVoiceLogic) UploadVoice() (resp *types.UploadRecord, err error) {
	// todo: add your logic here and delete this line

	return
}
