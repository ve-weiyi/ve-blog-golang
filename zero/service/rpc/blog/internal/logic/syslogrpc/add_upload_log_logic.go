package syslogrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUploadLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUploadLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUploadLogLogic {
	return &AddUploadLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件
func (l *AddUploadLogLogic) AddUploadLog(in *syslogrpc.UploadLogReq) (*syslogrpc.UploadLogResp, error) {
	entity := convertUploadIn(in)
	_, err := l.svcCtx.TUploadRecordModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertUploadOut(entity), nil
}

func convertUploadIn(in *syslogrpc.UploadLogReq) (out *model.TUploadRecord) {
	out = &model.TUploadRecord{
		Id:       in.Id,
		UserId:   in.UserId,
		Label:    in.Label,
		FileName: in.FileName,
		FileSize: in.FileSize,
		FileMd5:  in.FileMd5,
		FileUrl:  in.FileUrl,
		//CreatedAt: time.Unix(in.CreatedAt, 0),
		//UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func convertUploadOut(in *model.TUploadRecord) (out *syslogrpc.UploadLogResp) {
	out = &syslogrpc.UploadLogResp{
		Id:        in.Id,
		UserId:    in.UserId,
		Label:     in.Label,
		FileName:  in.FileName,
		FileSize:  in.FileSize,
		FileMd5:   in.FileMd5,
		FileUrl:   in.FileUrl,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
