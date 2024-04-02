package uploadrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVoiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVoiceLogic {
	return &UploadVoiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传语言
func (l *UploadVoiceLogic) UploadVoice(in *blog.UploadRecordReq) (*blog.UploadRecordResp, error) {
	// todo: add your logic here and delete this line

	return &blog.UploadRecordResp{}, nil
}
