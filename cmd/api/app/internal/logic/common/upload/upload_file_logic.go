package upload

import (
	"context"
	"io"
	"path/filepath"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/syslogservice"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传文件（服务端上传）
func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileReq, file io.Reader, filename string, size int64) (resp *types.UploadFileResp, err error) {
	start := time.Now()
	// 使用 StorageProvider 上传
	result, err := l.svcCtx.StorageProvider.Upload(l.ctx, file, filename)
	if err != nil {
		return nil, err
	}
	fileURL := result.AccessURL

	// 记录上传日志
	uid, _ := metax.GetApiUserIdFromCtx(l.ctx)
	did, _ := metax.GetApiDeviceIdFromCtx(l.ctx)
	if _, logErr := l.svcCtx.SyslogService.CreateUploadLog(l.ctx, &syslogservice.CreateUploadLogRequest{
		UserId:   uid,
		DeviceId: did,
		FileBase: req.FileBase,
		FileName: filename,
		FileType: filepath.Ext(filename),
		FileSize: size,
		FileUrl:  fileURL,
	}); logErr != nil {
		l.Errorf("创建上传日志失败: %v", logErr)
	}

	resp = &types.UploadFileResp{
		FileInfo: types.FileInfoVO{
			FileBase:  req.FileBase,
			FileName:  filename,
			FileType:  filepath.Ext(filename),
			FileSize:  size,
			FileUrl:   fileURL,
			UpdatedAt: time.Now().UnixMilli(),
		},
	}

	l.Infof("[Upload] File completed, cost=%dms", time.Since(start).Milliseconds())
	return resp, nil
}
