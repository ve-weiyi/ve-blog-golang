package uploadrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUploadRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUploadRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUploadRecordLogic {
	return &AddUploadRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件
func (l *AddUploadRecordLogic) AddUploadRecord(in *blog.UploadRecordReq) (*blog.UploadRecordResp, error) {
	entity := convert.ConvertUploadPbToModel(in)
	_, err := l.svcCtx.UploadRecordModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertUploadModelToPb(entity), nil
}
