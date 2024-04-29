package uploadrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件
func (l *UploadFileLogic) UploadFile(in *blog.UploadRecordReq) (*blog.UploadRecordResp, error) {
	entity := convert.ConvertUploadPbToModel(in)
	insert, err := l.svcCtx.UploadRecordModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertUploadModelToPb(insert), nil
}
